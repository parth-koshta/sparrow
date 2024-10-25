package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePrompt(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user first
		user := createRandomUser(t, testQueries)

		arg := CreatePromptParams{
			UserID: user.ID,
			Text:   "Write a blog post about AI",
		}
		prompt, err := testQueries.CreatePrompt(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, prompt)

		require.Equal(t, arg.UserID, prompt.UserID)
		require.Equal(t, arg.Text, prompt.Text)
		require.NotZero(t, prompt.CreatedAt)
	})
}

func TestListPromptsByUserID(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user
		user := createRandomUser(t, testQueries)

		// Create multiple prompts for the user
		for i := 0; i < 5; i++ {
			createRandomPrompt(t, testQueries, user.ID)
		}

		// Paginate through the prompts
		arg := ListPromptsByUserIDParams{
			UserID: user.ID,
			Limit:  3,
			Offset: 0,
		}
		prompts, err := testQueries.ListPromptsByUserID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, prompts, 3)

		arg.Offset = 3
		prompts, err = testQueries.ListPromptsByUserID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, prompts, 2)
	})
}
