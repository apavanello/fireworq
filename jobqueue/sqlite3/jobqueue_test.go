package sqlite3

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/fireworq/fireworq/config"
	"github.com/fireworq/fireworq/model"
	"github.com/fireworq/fireworq/test"
	"github.com/fireworq/fireworq/test/jobqueue"
	"github.com/fireworq/fireworq/test/mysql"
)

func TestMain(m *testing.M) {
	config.Locally("driver", "sqlite3", func() {
		status, err := test.Run(m)
		if err != nil {
			panic(err)
		}
		os.Exit(status)
	})
}

// Common tests

func TestNew(t *testing.T) {
	_ = New(&model.Queue{Name: "test", MaxWorkers: 30}, ":memory:")
}

func TestSubtests(t *testing.T) {
	jqtest.TestSubtests(t, runSubtests)
}

func runSubtests(t *testing.T, db, q string, tests []jqtest.Subtest) {
	dsn := Dsn()

	jq := New(&model.Queue{Name: q, MaxWorkers: 30}, dsn)
	jq.Start()
	defer func() { <-jq.Stop() }()
	time.Sleep(500 * time.Millisecond) // wait for up

	for _, test := range tests {
		err := mysqltest.TruncateTables(dsn)
		if err != nil {
			t.Error(err)
		}
		log.Println("Using DSN ", dsn)
		test(t, jq)
	}
}
