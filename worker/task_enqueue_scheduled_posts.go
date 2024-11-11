package worker

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	dbtypes "github.com/parth-koshta/sparrow/db/types"
	"github.com/rs/zerolog/log"
)

const TaskEnqueueScheduledPosts = "task:enqueue-scheduled-posts"

func (s *RedisTaskScheduler) ScheduleTaskEnqueueScheduledPosts(ctx context.Context, opts ...asynq.Option) error {
	task := asynq.NewTask(TaskEnqueueScheduledPosts, nil, opts...)
	// _, err := s.scheduler.Register("* * * * *", task)
	_, err := s.scheduler.Register("0 * * * *", task)
	if err != nil {
		return fmt.Errorf("could not schedule task: %w", err)
	}
	log.Info().Msgf("scheduler %s started", TaskEnqueueScheduledPosts)
	return nil
}

func (p *RedisTaskProcessor) ProcessTaskEnqueueScheduledPosts(ctx context.Context, task *asynq.Task) error {
	log.Info().Msg("processing task: enqueue scheduled posts")
	getScheduledPostsArgs := db.GetScheduledPostsWithinTimeframeParams{
		HoursFrom: 1,
		HoursTill: 1,
	}

	posts, err := p.store.GetScheduledPostsWithinTimeframe(ctx, getScheduledPostsArgs)
	if err != nil {
		return fmt.Errorf("could not get scheduled posts: %w", err)
	}

	log.Info().Msgf("found %d scheduled posts", len(posts))

	for _, post := range posts {
		log.Info().Msgf("enqueuing task: publish post %s (%s) -> %s", uuid.UUID(post.ID.Bytes).String(), post.Status, post.Text)
		publishPostArgs := &PayloadPublishPost{
			PostID:        post.ID,
			Text:          post.Text,
			ScheduledTime: post.ScheduledTime,
			ScheduleID:    post.ScheduleID,
		}
		err := p.distributor.DistributeTaskPublishPost(ctx, publishPostArgs)
		if err != nil && err != asynq.ErrTaskIDConflict {
			return fmt.Errorf("could not distribute task: %w", err)
		}

		updatePostStatusArgs := db.UpdatePostStatusParams{
			ID:     post.ID,
			Status: string(dbtypes.PostStatusEnqueued),
		}
		_, err = p.store.UpdatePostStatus(ctx, updatePostStatusArgs)
		if err != nil {
			return fmt.Errorf("could not update post status: %w", err)
		}
	}
	return nil
}
