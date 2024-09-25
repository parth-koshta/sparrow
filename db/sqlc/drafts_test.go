package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateDraft(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user, prompt, and post suggestion
		user := createRandomUser(t, testQueries)
		prompt := createRandomPrompt(t, testQueries, user.ID)
		suggestion := createRandomPostSuggestion(t, testQueries, prompt.ID)

		arg := CreateDraftParams{
			UserID:       user.ID,
			SuggestionID: suggestion.ID,
			DraftText:    "This is a draft text for a blog post.",
		}
		draft, err := testQueries.CreateDraft(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, draft)

		require.Equal(t, arg.UserID, draft.UserID)
		require.Equal(t, arg.SuggestionID, draft.SuggestionID)
		require.Equal(t, arg.DraftText, draft.DraftText)
		require.NotZero(t, draft.CreatedAt)
	})
}

func TestListDraftsByUserID(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user and multiple drafts
		user := createRandomUser(t, testQueries)
		for i := 0; i < 5; i++ {
			suggestion := createRandomPostSuggestion(t, testQueries, createRandomPrompt(t, testQueries, user.ID).ID)
			createRandomDraft(t, testQueries, user.ID, suggestion.ID)
		}

		arg := ListDraftsByUserIDParams{
			UserID: user.ID,
			Limit:  3,
			Offset: 0,
		}
		drafts, err := testQueries.ListDraftsByUserID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, drafts, 3)

		arg.Offset = 3
		drafts, err = testQueries.ListDraftsByUserID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, drafts, 2)
	})
}
