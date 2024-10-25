package main

import (
	"context"
	"os"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/parth-koshta/sparrow/api"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/util"
	"github.com/parth-koshta/sparrow/worker"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Error().Err(err).Msg("cannot load config")
	}

	if config.ENVIRONMENT == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx := context.Background()
	conn, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to db")
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Error().Err(err).Msg("cannot ping db")
	}

	store := db.NewStore(conn)

	redisOptions := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOptions)
	go runTaskProcessor(redisOptions, store)

	server, err := api.NewServer(store, config, taskDistributor)
	if err != nil {
		log.Error().Err(err).Msg("cannot create server")
	}
	log.Info().Msgf("starting server at %s", config.ServerAddress)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("cannot start server")
	}
}

func runTaskProcessor(redisOptions asynq.RedisClientOpt, store db.Store) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOptions, store)
	err := taskProcessor.Start()
	if err != nil {
		log.Error().Err(err).Msg("cannot start task processor")
	}
}
