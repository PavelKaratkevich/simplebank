package main

import (
	"database/sql"
	"log"
	db "simplebank/db/sqlc"
	"simplebank/api"
	_"github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
const serverAddress = "0.0.0.0:8080"

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error while opening db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(serverAddress)
}