package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/hibiken/asynq"
)

const TaskSendVerifyEmail = "task:send-verify-email"

type PayloadSendVerifyEmail struct {
	Email string `json:"email"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("could not marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("could not enqueue task: %w", err)
	}
	log.Info().Msgf("task enqueued: id=%s queue=%s", info.ID, info.Queue)
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("could not unmarshal task payload: %w", asynq.SkipRetry)
	}

	// Send verification email to the user
	user, err := processor.store.GetUserByEmail(ctx, payload.Email)
	if err == sql.ErrNoRows {
		return fmt.Errorf("user does not exist: %w", asynq.SkipRetry)
	}
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	log.Info().Msgf("sending verification email to %s", user.Email)
	return nil
}
