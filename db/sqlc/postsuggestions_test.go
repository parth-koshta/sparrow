package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePostSuggestion(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user and a prompt
		user := createRandomUser(t, testQueries)
		prompt := createRandomPrompt(t, testQueries, user.ID)

		arg := CreatePostSuggestionParams{
			PromptID:       prompt.ID,
			SuggestionText: "Use AI to automate content creation",
		}
		postSuggestion, err := testQueries.CreatePostSuggestion(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, postSuggestion)

		require.Equal(t, arg.PromptID, postSuggestion.PromptID)
		require.Equal(t, arg.SuggestionText, postSuggestion.SuggestionText)
		require.NotZero(t, postSuggestion.CreatedAt)
	})
}

func TestListPostSuggestionsByPromptID(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user and a prompt
		user := createRandomUser(t, testQueries)
		prompt := createRandomPrompt(t, testQueries, user.ID)

		// Create multiple suggestions
		for i := 0; i < 5; i++ {
			createRandomPostSuggestion(t, testQueries, prompt.ID)
		}

		// Paginate through the suggestions
		arg := ListPostSuggestionsByPromptIDParams{
			PromptID: prompt.ID,
			Limit:    3,
			Offset:   0,
		}
		suggestions, err := testQueries.ListPostSuggestionsByPromptID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, suggestions, 3)

		arg.Offset = 3
		suggestions, err = testQueries.ListPostSuggestionsByPromptID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, suggestions, 2)
	})
}
