// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: socialaccounts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSocialAccount = `-- name: CreateSocialAccount :one
INSERT INTO socialaccounts (
  user_id, platform, account_name, access_token
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, user_id, platform, account_name, access_token, created_at, updated_at
`

type CreateSocialAccountParams struct {
	UserID      pgtype.UUID
	Platform    string
	AccountName string
	AccessToken string
}

func (q *Queries) CreateSocialAccount(ctx context.Context, arg CreateSocialAccountParams) (Socialaccount, error) {
	row := q.db.QueryRow(ctx, createSocialAccount,
		arg.UserID,
		arg.Platform,
		arg.AccountName,
		arg.AccessToken,
	)
	var i Socialaccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.AccountName,
		&i.AccessToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSocialAccount = `-- name: DeleteSocialAccount :one
DELETE FROM socialaccounts
WHERE id = $1
RETURNING id, user_id, platform, account_name, access_token, created_at, updated_at
`

func (q *Queries) DeleteSocialAccount(ctx context.Context, id pgtype.UUID) (Socialaccount, error) {
	row := q.db.QueryRow(ctx, deleteSocialAccount, id)
	var i Socialaccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.AccountName,
		&i.AccessToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSocialAccountByID = `-- name: GetSocialAccountByID :one
SELECT id, user_id, platform, account_name, access_token, created_at, updated_at
FROM socialaccounts
WHERE id = $1
`

func (q *Queries) GetSocialAccountByID(ctx context.Context, id pgtype.UUID) (Socialaccount, error) {
	row := q.db.QueryRow(ctx, getSocialAccountByID, id)
	var i Socialaccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.AccountName,
		&i.AccessToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listSocialAccountsByUserID = `-- name: ListSocialAccountsByUserID :many
SELECT id, user_id, platform, account_name, access_token, created_at, updated_at
FROM socialaccounts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListSocialAccountsByUserIDParams struct {
	UserID pgtype.UUID
	Limit  int32
	Offset int32
}

func (q *Queries) ListSocialAccountsByUserID(ctx context.Context, arg ListSocialAccountsByUserIDParams) ([]Socialaccount, error) {
	rows, err := q.db.Query(ctx, listSocialAccountsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Socialaccount
	for rows.Next() {
		var i Socialaccount
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Platform,
			&i.AccountName,
			&i.AccessToken,
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

const updateSocialAccount = `-- name: UpdateSocialAccount :one
UPDATE socialaccounts
SET platform = $2,
    account_name = $3,
    access_token = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, platform, account_name, access_token, created_at, updated_at
`

type UpdateSocialAccountParams struct {
	ID          pgtype.UUID
	Platform    string
	AccountName string
	AccessToken string
}

func (q *Queries) UpdateSocialAccount(ctx context.Context, arg UpdateSocialAccountParams) (Socialaccount, error) {
	row := q.db.QueryRow(ctx, updateSocialAccount,
		arg.ID,
		arg.Platform,
		arg.AccountName,
		arg.AccessToken,
	)
	var i Socialaccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.AccountName,
		&i.AccessToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
