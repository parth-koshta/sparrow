package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	dbtypes "github.com/parth-koshta/sparrow/db/types"
)

type SchedulePostTxParams struct {
	UserID          pgtype.UUID
	PostID          pgtype.UUID
	ScheduledTime   pgtype.Timestamp
	SocialAccountID pgtype.UUID
}

type SchedulePostTxResult struct {
	ID pgtype.UUID
}

func (store *SQLStore) SchedulePostTx(ctx context.Context, arg SchedulePostTxParams) (SchedulePostTxResult, error) {
	var result SchedulePostTxResult
	err := store.ExecTx(ctx, func(q *Queries) error {
		scheduleArg := SchedulePostParams{
			UserID:          arg.UserID,
			PostID:          arg.PostID,
			ScheduledTime:   arg.ScheduledTime,
			SocialAccountID: arg.SocialAccountID,
			Status:          string(dbtypes.PostStatusScheduled),
		}

		scheduledPost, err := q.SchedulePost(ctx, scheduleArg)
		if err != nil {
			return err
		}

		updatePostStatusArg := UpdatePostStatusParams{
			ID:     arg.PostID,
			Status: string(dbtypes.PostStatusScheduled),
		}

		_, err = q.UpdatePostStatus(ctx, updatePostStatusArg)
		if err != nil {
			return err
		}

		result.ID = scheduledPost.ID

		return nil
	})

	return result, err
}
