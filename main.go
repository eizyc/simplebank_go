package main

import (
	"context"
	"log"

	"github.com/eizyc/simplebank/api"
	db "github.com/eizyc/simplebank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := *db.NewStore(connPool)
	server, err := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
