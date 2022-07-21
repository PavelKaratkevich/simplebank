package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadEnvVars("../..")
	if err != nil {
		log.Fatal("Error while loading env vars: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error while opening db:", err)
	}

	testQuery = New(testDB)

	os.Exit(m.Run())
}