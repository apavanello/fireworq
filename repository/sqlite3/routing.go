package sqlite3

import (
	"database/sql"
	"github.com/fireworq/fireworq/model"
	"github.com/fireworq/fireworq/repository"
	"sync"
)

type routingRepository struct {
	sync.RWMutex
	db       *sql.DB
	routings map[string]string
}

// NewRoutingRepository creates a repository.RoutingRepository which uses
// SQLite3 as a data store.
func NewRoutingRepository(db *sql.DB) repository.RoutingRepository {
	r := &routingRepository{db: db}
	r.Reload()
	return r
}

func (r *routingRepository) Add(jobCategory string, queueName string) (bool, error) {
	updated := false

	selectSQL := `
		 SELECT 1 FROM queue
		 WHERE name = ?
	`
	var queueExists int
	err := r.db.QueryRow(selectSQL, queueName).Scan(
		&queueExists,
	)
	if err != nil {
		return updated, &repository.QueueNotFoundError{QueueName: queueName}
	}

	insertSQL := `
		INSERT INTO routing (job_category, queue_name) 
		VALUES (?, ?)
		ON CONFLICT (job_category) DO 
		    UPDATE SET queue_name = excluded.queue_name; 
	`
	res, err := r.db.Exec(insertSQL, jobCategory, queueName)
	if err != nil {
		return updated, err
	}
	i, err := res.RowsAffected()
	if err == nil {
		updated = updated || (i != 0)
	}

	if updated {
		r.Lock()
		defer r.Unlock()

		r.routings[jobCategory] = queueName
		return updated, r.updateRevision()
	}
	return updated, nil
}

func (r *routingRepository) FindQueueNameByJobCategory(category string) string {
	r.RLock()
	defer r.RUnlock()

	return r.routings[category]
}

func (r *routingRepository) FindAll() ([]model.Routing, error) {
	sqlStatement := `
		SELECT queue_name, job_category FROM routing
		ORDER BY queue_name
	`

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]model.Routing, 0)
	for rows.Next() {
		var row model.Routing
		if err := rows.Scan(&(row.QueueName), &(row.JobCategory)); err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	r.Lock()
	defer r.Unlock()

	r.routings = make(map[string]string, len(results))
	for _, routing := range results {
		r.routings[routing.JobCategory] = routing.QueueName
	}

	return results, nil
}

func (r *routingRepository) DeleteByJobCategory(category string) error {
	sqlStatement := `
		DELETE FROM routing
		WHERE job_category = ?
	`
	_, err := r.db.Exec(sqlStatement, category)
	if err != nil {
		return err
	}

	r.Lock()
	defer r.Unlock()

	delete(r.routings, category)
	return r.updateRevision()
}

func (r *routingRepository) Revision() (uint64, error) {
	var revision uint64

	sqlStatement := `
		SELECT revision FROM config_revision
		WHERE name = 'routing'
	`

	if err := r.db.QueryRow(sqlStatement).Scan(&revision); err != nil {
		return 0, err
	}
	return revision, nil
}

func (r *routingRepository) Reload() error {
	_, err := r.FindAll()
	return err
}

func (r *routingRepository) updateRevision() error {
	sqlStatement := `
		INSERT INTO config_revision (name, revision)
		VALUES ('routing', 1)
		ON CONFLICT (name) DO 
		    UPDATE SET revision = revision + 1;
	`

	_, err := r.db.Exec(sqlStatement)
	return err
}
