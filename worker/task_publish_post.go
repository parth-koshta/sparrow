package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/parth-koshta/sparrow/client"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	dbtypes "github.com/parth-koshta/sparrow/db/types"
	"github.com/rs/zerolog/log"
)

const TaskPublishPosts = "task:publish-posts"

type PayloadPublishPost struct {
	PostID          pgtype.UUID      `json:"post_id"`
	ScheduleID      pgtype.UUID      `json:"schedule_id"`
	Text            string           `json:"text"`
	ScheduledTime   pgtype.Timestamp `json:"scheduled_time"`
	SocialAccountID pgtype.UUID      `json:"social_account_id"`
}

func (d *RedisTaskDistributor) DistributeTaskPublishPost(ctx context.Context, payload *PayloadPublishPost, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("could not marshal task payload: %w", err)
	}
	opts = append(opts, asynq.MaxRetry(10), asynq.Queue(QueueCritical))

	if payload.ScheduledTime.Time.After(time.Now()) {
		opts = append(opts, asynq.ProcessAt(payload.ScheduledTime.Time))
	}

	task := asynq.NewTask(TaskPublishPosts, jsonPayload, opts...)
	info, err := d.client.EnqueueContext(ctx, task, asynq.TaskID(uuid.UUID(payload.ScheduleID.Bytes).String()))
	if err != nil {
		log.Error().Err(err).Msg("could not enqueue task")
		return err
	}
	log.Info().Msgf("task enqueued: id=%s queue=%s", info.ID, info.Queue)

	return nil
}

func (p *RedisTaskProcessor) ProcessTaskPublishPost(ctx context.Context, task *asynq.Task) error {
	log.Info().Msg("ProcessTaskPublishPost: processing task: publish post")

	var payload PayloadPublishPost
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("ProcessTaskPublishPost: could not unmarshal task payload: %w", err)
	}

	log.Info().Msgf("ProcessTaskPublishPost: publishing post %s -> %s", uuid.UUID(payload.PostID.Bytes).String(), payload.Text)

	post, err := p.store.GetPostByID(ctx, payload.PostID)
	if err != nil {
		return fmt.Errorf("ProcessTaskPublishPost: could not get post by id: %w", err)
	}

	if post.Status != string(dbtypes.PostStatusScheduled) {
		return fmt.Errorf("ProcessTaskPublishPost: post status is not scheduled")
	}

	socialAccount, err := p.store.GetSocialAccountByID(ctx, payload.SocialAccountID)
	if err != nil {
		return fmt.Errorf("ProcessTaskPublishPost: could not get social account by id: %w", err)
	}

	if socialAccount.UserID.Bytes != post.UserID.Bytes {
		return fmt.Errorf("ProcessTaskPublishPost: social account does not belong to user")
	}

	err = p.linkedinClient.PublishPost(socialAccount.AccessToken, socialAccount.LinkedinSub.String, client.PayloadPublishPost{PostID: post.ID, Text: post.Text})
	if err != nil {
		return fmt.Errorf("ProcessTaskPublishPost: could not publish post: %w", err)
	}

	updatePostStatusArgs := db.UpdatePostStatusParams{
		ID:     payload.PostID,
		Status: string(dbtypes.PostStatusPublished),
	}
	_, err = p.store.UpdatePostStatus(ctx, updatePostStatusArgs)
	if err != nil {
		return fmt.Errorf("ProcessTaskPublishPost: could not update post status: %w", err)
	}

	_, err = p.store.UpdatePostScheduleExectued(ctx, payload.ScheduleID)
	if err != nil {
		return fmt.Errorf("ProcessTaskPublishPost: could not update post schedule executed time: %w", err)
	}

	return nil
}
