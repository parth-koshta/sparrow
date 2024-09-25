package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateScheduledPost(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user, prompt, suggestion, draft
		user := createRandomUser(t, testQueries)
		prompt := createRandomPrompt(t, testQueries, user.ID)
		suggestion := createRandomPostSuggestion(t, testQueries, prompt.ID)
		draft := createRandomDraft(t, testQueries, user.ID, suggestion.ID)

		// Prepare the scheduled time as a pgtype.Timestamp
		scheduledTime := pgtype.Timestamp{
			Time:  time.Now().UTC().Add(24 * time.Hour), // Ensure time is in UTC
			Valid: true,
		}

		// Prepare the argument for creating a scheduled post
		arg := CreateScheduledPostParams{
			UserID:        user.ID,
			DraftID:       draft.ID,
			ScheduledTime: scheduledTime,
			Status:        "scheduled",
		}

		// Create the scheduled post
		scheduledPost, err := testQueries.CreateScheduledPost(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, scheduledPost)

		// Validate the properties of the created scheduled post
		require.Equal(t, arg.UserID, scheduledPost.UserID)
		require.Equal(t, arg.DraftID, scheduledPost.DraftID)
		require.True(t, scheduledPost.ScheduledTime.Valid)

		// Use the scheduledTime for comparison
		require.WithinDuration(t, scheduledPost.ScheduledTime.Time, scheduledTime.Time, time.Second)
		require.Equal(t, arg.Status, scheduledPost.Status)
	})
}

func TestListScheduledPostsByUserID(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user and multiple scheduled posts
		user := createRandomUser(t, testQueries)
		for i := 0; i < 5; i++ {
			draft := createRandomDraft(t, testQueries, user.ID, createRandomPostSuggestion(t, testQueries, createRandomPrompt(t, testQueries, user.ID).ID).ID)
			createRandomScheduledPost(t, testQueries, user.ID, draft.ID)
		}

		arg := ListScheduledPostsByUserIDParams{
			UserID: user.ID,
			Limit:  3,
			Offset: 0,
		}
		posts, err := testQueries.ListScheduledPostsByUserID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, posts, 3)

		arg.Offset = 3
		posts, err = testQueries.ListScheduledPostsByUserID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, posts, 2)
	})
}
