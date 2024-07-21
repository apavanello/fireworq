package sqlite3test

import (
	"database/sql"
	"fmt"
	"os"
)

// With runs a block with locking the DB and truncating all tables in
// the DB.
func With(dsn string, block func()) error {

	// Create a temporary database file
	if dsn == "" {
		file, err := os.CreateTemp("", "testdb-*.sqlite3")
		if err != nil {
			return fmt.Errorf("failed to create temporary db file: %w", err)
		}
		defer os.Remove(file.Name())
		dsn = file.Name()
	}
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	block()
	return nil
}
