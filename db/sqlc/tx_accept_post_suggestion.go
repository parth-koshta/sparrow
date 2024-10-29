package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	dbtypes "github.com/parth-koshta/sparrow/db/types"
)

type AcceptPostSuggestionTxParams struct {
	UserID       pgtype.UUID
	SuggestionID pgtype.UUID
}

type AcceptPostSuggestionTxResult struct {
	PostID pgtype.UUID
}

func (store *SQLStore) AcceptPostSuggestionTx(ctx context.Context, arg AcceptPostSuggestionTxParams) (AcceptPostSuggestionTxResult, error) {
	var result AcceptPostSuggestionTxResult

	err := store.ExecTx(ctx, func(q *Queries) error {
		var err error

		// Create Post
		postArg := CreatePostParams{
			UserID:       arg.UserID,
			SuggestionID: arg.SuggestionID,
			Text:         "",
			Status:       string(dbtypes.PostStatusDraft),
		}

		// Retrieve the suggestion text (assuming you have a method to do this)
		suggestion, err := q.GetPostSuggestionByID(ctx, arg.SuggestionID)
		if err != nil {
			return err
		}

		// Assign suggestion text to post argument
		postArg.Text = suggestion.Text

		post, err := q.CreatePost(ctx, postArg)
		if err != nil {
			return err
		}

		// Update suggestion status
		updateSuggestionStatusArg := UpdatePostSuggestionStatusParams{
			ID:     arg.SuggestionID,
			Status: string(dbtypes.PostSuggestionStatusAccepted),
		}

		_, err = q.UpdatePostSuggestionStatus(ctx, updateSuggestionStatusArg)
		if err != nil {
			return err
		}

		result.PostID = post.ID
		return nil
	})

	return result, err
}
