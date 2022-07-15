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


func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error while opening db:", err)
	}

	testQuery = New(conn)

	os.Exit(m.Run())
}