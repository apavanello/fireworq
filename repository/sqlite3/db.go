//go:generate go-assets-builder -p sqlite3 -o assets.go ../../data/repository/sqlite3

package sqlite3

import (
	"database/sql"
	"github.com/fireworq/fireworq/config"
	"io"

	_ "github.com/mattn/go-sqlite3" // initialize the driver
	"github.com/rs/zerolog/log"
)

var schema []string

func init() {
	schema = []string{
		"/data/repository/sqlite3/schema/queue.sql",
		"/data/repository/sqlite3/schema/queue_throttle.sql",
		"/data/repository/sqlite3/schema/routing.sql",
		"/data/repository/sqlite3/schema/config_revision.sql",
	}
}

// Dsn returns the data source name of the storage specified in the
// configuration.
func Dsn() string {
	dsn := config.Get("repository_sqlite3_dsn")
	if dsn != "" {
		return dsn
	}
	return config.Get("sqlite3_dsn")
}

// NewDB creates an instance of DB handler.
func NewDB() (*sql.DB, error) {
	dsn := Dsn()

	log.Info().Msgf("Connecting database %s ...", dsn)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	for _, path := range schema {
		f, err := Assets.Open(path)
		if err != nil {
			log.Panic().Msg(err.Error())
		}

		query, err := io.ReadAll(f)

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		err = f.Close()

		if err != nil {
			log.Panic().Msg(err.Error())
		}

		_, err = db.Exec(string(query))
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
