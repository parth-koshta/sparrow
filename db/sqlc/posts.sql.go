// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  user_id, suggestion_id, text, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, user_id, suggestion_id, text, status, created_at, updated_at
`

type CreatePostParams struct {
	UserID       pgtype.UUID `json:"user_id"`
	SuggestionID pgtype.UUID `json:"suggestion_id"`
	Text         string      `json:"text"`
	Status       string      `json:"status"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.UserID,
		arg.SuggestionID,
		arg.Text,
		arg.Status,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SuggestionID,
		&i.Text,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING id, user_id, suggestion_id, text, status, created_at, updated_at
`

func (q *Queries) DeletePost(ctx context.Context, id pgtype.UUID) (Post, error) {
	row := q.db.QueryRow(ctx, deletePost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SuggestionID,
		&i.Text,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostByID = `-- name: GetPostByID :one
SELECT id, user_id, suggestion_id, text, created_at, updated_at
FROM posts
WHERE id = $1
`

type GetPostByIDRow struct {
	ID           pgtype.UUID      `json:"id"`
	UserID       pgtype.UUID      `json:"user_id"`
	SuggestionID pgtype.UUID      `json:"suggestion_id"`
	Text         string           `json:"text"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) GetPostByID(ctx context.Context, id pgtype.UUID) (GetPostByIDRow, error) {
	row := q.db.QueryRow(ctx, getPostByID, id)
	var i GetPostByIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SuggestionID,
		&i.Text,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPostsByUserID = `-- name: ListPostsByUserID :many
SELECT id, user_id, suggestion_id, text, created_at, updated_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListPostsByUserIDParams struct {
	UserID pgtype.UUID `json:"user_id"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

type ListPostsByUserIDRow struct {
	ID           pgtype.UUID      `json:"id"`
	UserID       pgtype.UUID      `json:"user_id"`
	SuggestionID pgtype.UUID      `json:"suggestion_id"`
	Text         string           `json:"text"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) ListPostsByUserID(ctx context.Context, arg ListPostsByUserIDParams) ([]ListPostsByUserIDRow, error) {
	rows, err := q.db.Query(ctx, listPostsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPostsByUserIDRow
	for rows.Next() {
		var i ListPostsByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SuggestionID,
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

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET text = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, suggestion_id, text, status, created_at, updated_at
`

type UpdatePostParams struct {
	ID   pgtype.UUID `json:"id"`
	Text string      `json:"text"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, updatePost, arg.ID, arg.Text)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SuggestionID,
		&i.Text,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
