package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mkarimi-s/bank/api"
	db "github.com/mkarimi-s/bank/db/sqlc"
	"github.com/mkarimi-s/bank/util"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file", err)
	}


	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if (err != nil){
		log.Fatal("can not start server", err)
	}
}