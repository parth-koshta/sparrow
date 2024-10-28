package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskEnqueueScheduledPosts = "task:enqueue-scheduled-posts"

func (s *RedisTaskScheduler) ScheduleTaskEnqueueScheduledPosts(ctx context.Context, opts ...asynq.Option) error {
	task := asynq.NewTask(TaskEnqueueScheduledPosts, nil, opts...)
	_, err := s.scheduler.Register("* * * * *", task)
	if err != nil {
		return fmt.Errorf("could not schedule task: %w", err)
	}
	log.Info().Msgf("scheduler %s started", TaskEnqueueScheduledPosts)
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskEnqueueScheduledPosts(ctx context.Context, task *asynq.Task) error {
	log.Info().Msg("processing task: enqueue scheduled posts")
	return nil
}
