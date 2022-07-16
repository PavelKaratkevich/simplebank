package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_"github.com/lib/pq"
)

var testQuery *Queries
const dbDriver = "postgres"
const dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

var testDB *sql.DB

func TestMain(m *testing.M) {

	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error while opening db:", err)
	}

	testQuery = New(testDB)

	os.Exit(m.Run())
}