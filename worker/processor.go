package worker

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/hibiken/asynq"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/mail"
	"github.com/parth-koshta/sparrow/util"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	Shutdown()
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
	ProcessTaskEnqueueScheduledPosts(ctx context.Context, task *asynq.Task) error
	// ProcessTaskPublishPost(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server      *asynq.Server
	store       db.Store
	mailer      mail.EmailSender
	config      util.Config
	distributor TaskDistributor
}

func NewRedisTaskProcessor(redisOptions asynq.RedisClientOpt, store db.Store, mailer mail.EmailSender, config util.Config, distributor TaskDistributor) TaskProcessor {
	server := asynq.NewServer(redisOptions, asynq.Config{
		Queues: map[string]int{
			QueueCritical: 10,
			QueueDefault:  5,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().Err(err).Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("task processing failed")
		}),
		Logger: NewLogger(),
	})

	return &RedisTaskProcessor{
		server:      server,
		store:       store,
		mailer:      mailer,
		config:      config,
		distributor: distributor,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)
	mux.HandleFunc(TaskEnqueueScheduledPosts, processor.ProcessTaskEnqueueScheduledPosts)
	mux.HandleFunc(TaskPublishPosts, processor.ProcessTaskPublishPost)
	log.Info().Msg("starting task processor")
	return processor.server.Start(mux)
}

func (processor *RedisTaskProcessor) Shutdown() {
	processor.server.Shutdown()
}
