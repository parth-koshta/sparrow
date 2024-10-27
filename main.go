package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/parth-koshta/sparrow/api"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/mail"
	"github.com/parth-koshta/sparrow/util"
	"github.com/parth-koshta/sparrow/worker"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Error().Err(err).Msg("cannot load config")
		return
	}

	if config.ENVIRONMENT == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	conn, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to db")
		return
	}
	defer conn.Close()

	if err := conn.Ping(ctx); err != nil {
		log.Error().Err(err).Msg("cannot ping db")
		return
	}

	store := db.NewStore(conn)
	redisOptions := asynq.RedisClientOpt{Addr: config.RedisAddress}
	taskDistributor := worker.NewRedisTaskDistributor(redisOptions)

	waitGroup, ctx := errgroup.WithContext(ctx)

	runTaskProcessor(ctx, waitGroup, redisOptions, store, config)
	runServer(ctx, waitGroup, store, config, taskDistributor)

	if err := waitGroup.Wait(); err != nil {
		log.Error().Err(err).Msg("error from server or task processor")
	}
}

func runTaskProcessor(ctx context.Context, waitGroup *errgroup.Group, redisOptions asynq.RedisClientOpt, store db.Store, config util.Config) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOptions, store, mailer)

	if err := taskProcessor.Start(); err != nil {
		log.Error().Err(err).Msg("cannot start task processor")
		return
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		taskProcessor.Shutdown()
		return nil
	})
}

func runServer(ctx context.Context, waitGroup *errgroup.Group, store db.Store, config util.Config, taskDistributor worker.TaskDistributor) {
	server, err := api.NewServer(store, config, taskDistributor)
	if err != nil {
		log.Error().Err(err).Msg("cannot create server")
		return
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("starting server at %s", config.ServerAddress)
		return server.Start(config.ServerAddress)
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		if err := server.Stop(ctx); err != nil {
			log.Error().Err(err).Msg("server shutdown error")
		}
		return nil
	})
}
