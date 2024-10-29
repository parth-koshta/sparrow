package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type TaskScheduler interface {
	Start() error
	Stop()
	ScheduleTaskEnqueueScheduledPosts(ctx context.Context, opts ...asynq.Option) error
}

type RedisTaskScheduler struct {
	scheduler *asynq.Scheduler
}

func NewRedisTaskScheduler(redisOptions asynq.RedisClientOpt) TaskScheduler {
	scheduler := asynq.NewScheduler(redisOptions, &asynq.SchedulerOpts{
		Logger: NewLogger(),
	})
	return &RedisTaskScheduler{
		scheduler: scheduler,
	}
}

func (s *RedisTaskScheduler) Start() error {
	if err := s.scheduler.Start(); err != nil {
		return err
	}
	log.Info().Msg("Scheduler started")
	return nil
}

func (s *RedisTaskScheduler) Stop() {
	s.scheduler.Shutdown()
}
