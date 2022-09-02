package main

import (
	"database/sql"
	"log"

	"github.com/isbik/bank/api"
	"github.com/isbik/bank/util"

	_ "github.com/lib/pq"

	db "github.com/isbik/bank/db/sqlc"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cound't connect to db", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err == nil {
		log.Fatal("cannot start server", err)
	}

}
