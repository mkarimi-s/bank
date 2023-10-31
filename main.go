package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mkarimi-s/bank/api"
	db "github.com/mkarimi-s/bank/db/sqlc"
)


const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/bank?sslmode=disable"
	serverAddree = "0.0.0.0:8080"
)


func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddree)
	if (err != nil){
		log.Fatal("can not start server", err)
	}
}