// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	BulkCreatePostSuggestions(ctx context.Context, arg BulkCreatePostSuggestionsParams) error
	CreateDraft(ctx context.Context, arg CreateDraftParams) (Draft, error)
	CreatePostSuggestion(ctx context.Context, arg CreatePostSuggestionParams) (Postsuggestion, error)
	CreatePrompt(ctx context.Context, arg CreatePromptParams) (Prompt, error)
	CreateScheduledPost(ctx context.Context, arg CreateScheduledPostParams) (Scheduledpost, error)
	CreateSocialAccount(ctx context.Context, arg CreateSocialAccountParams) (Socialaccount, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteDraft(ctx context.Context, id pgtype.UUID) (Draft, error)
	DeletePostSuggestion(ctx context.Context, id pgtype.UUID) (Postsuggestion, error)
	DeletePrompt(ctx context.Context, id pgtype.UUID) (Prompt, error)
	DeleteScheduledPost(ctx context.Context, id pgtype.UUID) (Scheduledpost, error)
	DeleteSocialAccount(ctx context.Context, id pgtype.UUID) (Socialaccount, error)
	GetDraftByID(ctx context.Context, id pgtype.UUID) (Draft, error)
	GetPostSuggestionByID(ctx context.Context, id pgtype.UUID) (Postsuggestion, error)
	GetPromptByID(ctx context.Context, id pgtype.UUID) (Prompt, error)
	GetScheduledPostByID(ctx context.Context, id pgtype.UUID) (Scheduledpost, error)
	GetSocialAccountByID(ctx context.Context, id pgtype.UUID) (GetSocialAccountByIDRow, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (GetUserByIDRow, error)
	ListDraftsByUserID(ctx context.Context, arg ListDraftsByUserIDParams) ([]Draft, error)
	ListPostSuggestionsByPromptID(ctx context.Context, arg ListPostSuggestionsByPromptIDParams) ([]Postsuggestion, error)
	ListPromptsByUserID(ctx context.Context, arg ListPromptsByUserIDParams) ([]Prompt, error)
	ListScheduledPostsByUserID(ctx context.Context, arg ListScheduledPostsByUserIDParams) ([]Scheduledpost, error)
	ListSocialAccountsByUserID(ctx context.Context, arg ListSocialAccountsByUserIDParams) ([]ListSocialAccountsByUserIDRow, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateDraft(ctx context.Context, arg UpdateDraftParams) (Draft, error)
	UpdatePostSuggestion(ctx context.Context, arg UpdatePostSuggestionParams) (Postsuggestion, error)
	UpdatePrompt(ctx context.Context, arg UpdatePromptParams) (Prompt, error)
	UpdateScheduledPost(ctx context.Context, arg UpdateScheduledPostParams) (Scheduledpost, error)
	UpdateSocialAccount(ctx context.Context, arg UpdateSocialAccountParams) (Socialaccount, error)
	UpdateSocialAccountToken(ctx context.Context, arg UpdateSocialAccountTokenParams) (Socialaccount, error)
}

var _ Querier = (*Queries)(nil)
