package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	_ "github.com/lib/pq"
)

func main() {

	config, err1 := util.LoadEnvVars(".")
	if err1 != nil {
		log.Fatal("Error while loading env variables:", err1.Error())
	}
	conn, err2 := sql.Open(config.DBDriver, config.DBSource)
	if err2 != nil {
		log.Fatal("Error while opening db: ", err2.Error())
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server")
	}

	server.Start(config.ServerAddress)
}
