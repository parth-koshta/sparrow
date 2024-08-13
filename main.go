package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/parth-koshta/sparrow/api"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5432/sparrow-dev?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal("cannot ping db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
