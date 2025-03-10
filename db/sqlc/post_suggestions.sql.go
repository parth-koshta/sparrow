// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: post_suggestions.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const bulkCreatePostSuggestions = `-- name: BulkCreatePostSuggestions :many
INSERT INTO post_suggestions (prompt_id, text)
SELECT $1, unnest($2::text[])
RETURNING id, prompt_id, text, created_at
`

type BulkCreatePostSuggestionsParams struct {
	PromptID    pgtype.UUID `json:"prompt_id"`
	Suggestions []string    `json:"suggestions"`
}

type BulkCreatePostSuggestionsRow struct {
	ID        pgtype.UUID      `json:"id"`
	PromptID  pgtype.UUID      `json:"prompt_id"`
	Text      string           `json:"text"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) BulkCreatePostSuggestions(ctx context.Context, arg BulkCreatePostSuggestionsParams) ([]BulkCreatePostSuggestionsRow, error) {
	rows, err := q.db.Query(ctx, bulkCreatePostSuggestions, arg.PromptID, arg.Suggestions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BulkCreatePostSuggestionsRow
	for rows.Next() {
		var i BulkCreatePostSuggestionsRow
		if err := rows.Scan(
			&i.ID,
			&i.PromptID,
			&i.Text,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createPostSuggestion = `-- name: CreatePostSuggestion :one
INSERT INTO post_suggestions (
  prompt_id, text
) VALUES (
  $1, $2
)
RETURNING id, prompt_id, text, status, created_at, updated_at
`

type CreatePostSuggestionParams struct {
	PromptID pgtype.UUID `json:"prompt_id"`
	Text     string      `json:"text"`
}

func (q *Queries) CreatePostSuggestion(ctx context.Context, arg CreatePostSuggestionParams) (PostSuggestion, error) {
	row := q.db.QueryRow(ctx, createPostSuggestion, arg.PromptID, arg.Text)
	var i PostSuggestion
	err := row.Scan(
		&i.ID,
		&i.PromptID,
		&i.Text,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePostSuggestion = `-- name: DeletePostSuggestion :one
DELETE FROM post_suggestions
WHERE id = $1
RETURNING id, prompt_id, text, status, created_at, updated_at
`

func (q *Queries) DeletePostSuggestion(ctx context.Context, id pgtype.UUID) (PostSuggestion, error) {
	row := q.db.QueryRow(ctx, deletePostSuggestion, id)
	var i PostSuggestion
	err := row.Scan(
		&i.ID,
		&i.PromptID,
		&i.Text,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostSuggestionByID = `-- name: GetPostSuggestionByID :one
SELECT id, prompt_id, text, created_at, updated_at
FROM post_suggestions
WHERE id = $1
`

type GetPostSuggestionByIDRow struct {
	ID        pgtype.UUID      `json:"id"`
	PromptID  pgtype.UUID      `json:"prompt_id"`
	Text      string           `json:"text"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) GetPostSuggestionByID(ctx context.Context, id pgtype.UUID) (GetPostSuggestionByIDRow, error) {
	row := q.db.QueryRow(ctx, getPostSuggestionByID, id)
	var i GetPostSuggestionByIDRow
	err := row.Scan(
		&i.ID,
		&i.PromptID,
		&i.Text,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPostSuggestionsByPromptID = `-- name: ListPostSuggestionsByPromptID :many
SELECT id, prompt_id, text, created_at, updated_at
FROM post_suggestions
WHERE prompt_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListPostSuggestionsByPromptIDParams struct {
	PromptID pgtype.UUID `json:"prompt_id"`
	Limit    int32       `json:"limit"`
	Offset   int32       `json:"offset"`
}

type ListPostSuggestionsByPromptIDRow struct {
	ID        pgtype.UUID      `json:"id"`
	PromptID  pgtype.UUID      `json:"prompt_id"`
	Text      string           `json:"text"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) ListPostSuggestionsByPromptID(ctx context.Context, arg ListPostSuggestionsByPromptIDParams) ([]ListPostSuggestionsByPromptIDRow, error) {
	rows, err := q.db.Query(ctx, listPostSuggestionsByPromptID, arg.PromptID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPostSuggestionsByPromptIDRow
	for rows.Next() {
		var i ListPostSuggestionsByPromptIDRow
		if err := rows.Scan(
			&i.ID,
			&i.PromptID,
			&i.Text,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePostSuggestionStatus = `-- name: UpdatePostSuggestionStatus :one
UPDATE post_suggestions
SET status = $2
WHERE id = $1
RETURNING id, prompt_id, text, status, created_at, updated_at
`

type UpdatePostSuggestionStatusParams struct {
	ID     pgtype.UUID `json:"id"`
	Status string      `json:"status"`
}

func (q *Queries) UpdatePostSuggestionStatus(ctx context.Context, arg UpdatePostSuggestionStatusParams) (PostSuggestion, error) {
	row := q.db.QueryRow(ctx, updatePostSuggestionStatus, arg.ID, arg.Status)
	var i PostSuggestion
	err := row.Scan(
		&i.ID,
		&i.PromptID,
		&i.Text,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
