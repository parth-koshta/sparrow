// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: social_accounts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSocialAccount = `-- name: CreateSocialAccount :one
INSERT INTO social_accounts (
  user_id, platform, name, email, access_token, id_token, token_expires_at, linkedin_sub
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, user_id, platform, name, email, access_token, id_token, token_expires_at, created_at, updated_at, linkedin_sub
`

type CreateSocialAccountParams struct {
	UserID         pgtype.UUID      `json:"user_id"`
	Platform       string           `json:"platform"`
	Name           string           `json:"name"`
	Email          string           `json:"email"`
	AccessToken    string           `json:"access_token"`
	IDToken        string           `json:"id_token"`
	TokenExpiresAt pgtype.Timestamp `json:"token_expires_at"`
	LinkedinSub    pgtype.Text      `json:"linkedin_sub"`
}

func (q *Queries) CreateSocialAccount(ctx context.Context, arg CreateSocialAccountParams) (SocialAccount, error) {
	row := q.db.QueryRow(ctx, createSocialAccount,
		arg.UserID,
		arg.Platform,
		arg.Name,
		arg.Email,
		arg.AccessToken,
		arg.IDToken,
		arg.TokenExpiresAt,
		arg.LinkedinSub,
	)
	var i SocialAccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.Name,
		&i.Email,
		&i.AccessToken,
		&i.IDToken,
		&i.TokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LinkedinSub,
	)
	return i, err
}

const deleteSocialAccount = `-- name: DeleteSocialAccount :one
DELETE FROM social_accounts
WHERE id = $1
RETURNING id, user_id, platform, name, email, access_token, id_token, token_expires_at, created_at, updated_at, linkedin_sub
`

func (q *Queries) DeleteSocialAccount(ctx context.Context, id pgtype.UUID) (SocialAccount, error) {
	row := q.db.QueryRow(ctx, deleteSocialAccount, id)
	var i SocialAccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.Name,
		&i.Email,
		&i.AccessToken,
		&i.IDToken,
		&i.TokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LinkedinSub,
	)
	return i, err
}

const getSocialAccountByID = `-- name: GetSocialAccountByID :one
SELECT platform, user_id, name, access_token, linkedin_sub, token_expires_at, updated_at
FROM social_accounts
WHERE id = $1
`

type GetSocialAccountByIDRow struct {
	Platform       string           `json:"platform"`
	UserID         pgtype.UUID      `json:"user_id"`
	Name           string           `json:"name"`
	AccessToken    string           `json:"access_token"`
	LinkedinSub    pgtype.Text      `json:"linkedin_sub"`
	TokenExpiresAt pgtype.Timestamp `json:"token_expires_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) GetSocialAccountByID(ctx context.Context, id pgtype.UUID) (GetSocialAccountByIDRow, error) {
	row := q.db.QueryRow(ctx, getSocialAccountByID, id)
	var i GetSocialAccountByIDRow
	err := row.Scan(
		&i.Platform,
		&i.UserID,
		&i.Name,
		&i.AccessToken,
		&i.LinkedinSub,
		&i.TokenExpiresAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listSocialAccountsByUserID = `-- name: ListSocialAccountsByUserID :many
SELECT id, user_id, platform, name, email, token_expires_at, created_at, updated_at
FROM social_accounts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListSocialAccountsByUserIDParams struct {
	UserID pgtype.UUID `json:"user_id"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

type ListSocialAccountsByUserIDRow struct {
	ID             pgtype.UUID      `json:"id"`
	UserID         pgtype.UUID      `json:"user_id"`
	Platform       string           `json:"platform"`
	Name           string           `json:"name"`
	Email          string           `json:"email"`
	TokenExpiresAt pgtype.Timestamp `json:"token_expires_at"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) ListSocialAccountsByUserID(ctx context.Context, arg ListSocialAccountsByUserIDParams) ([]ListSocialAccountsByUserIDRow, error) {
	rows, err := q.db.Query(ctx, listSocialAccountsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSocialAccountsByUserIDRow
	for rows.Next() {
		var i ListSocialAccountsByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Platform,
			&i.Name,
			&i.Email,
			&i.TokenExpiresAt,
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
UPDATE social_accounts
SET platform = $2,
    name = $3,
    access_token = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, platform, name, email, access_token, id_token, token_expires_at, created_at, updated_at, linkedin_sub
`

type UpdateSocialAccountParams struct {
	ID          pgtype.UUID `json:"id"`
	Platform    string      `json:"platform"`
	Name        string      `json:"name"`
	AccessToken string      `json:"access_token"`
}

func (q *Queries) UpdateSocialAccount(ctx context.Context, arg UpdateSocialAccountParams) (SocialAccount, error) {
	row := q.db.QueryRow(ctx, updateSocialAccount,
		arg.ID,
		arg.Platform,
		arg.Name,
		arg.AccessToken,
	)
	var i SocialAccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.Name,
		&i.Email,
		&i.AccessToken,
		&i.IDToken,
		&i.TokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LinkedinSub,
	)
	return i, err
}

const updateSocialAccountToken = `-- name: UpdateSocialAccountToken :one
UPDATE social_accounts
SET access_token = $2,
    id_token = $3,
    token_expires_at = $4,
    updated_at = NOW()
WHERE user_id = $1
RETURNING id, user_id, platform, name, email, access_token, id_token, token_expires_at, created_at, updated_at, linkedin_sub
`

type UpdateSocialAccountTokenParams struct {
	UserID         pgtype.UUID      `json:"user_id"`
	AccessToken    string           `json:"access_token"`
	IDToken        string           `json:"id_token"`
	TokenExpiresAt pgtype.Timestamp `json:"token_expires_at"`
}

func (q *Queries) UpdateSocialAccountToken(ctx context.Context, arg UpdateSocialAccountTokenParams) (SocialAccount, error) {
	row := q.db.QueryRow(ctx, updateSocialAccountToken,
		arg.UserID,
		arg.AccessToken,
		arg.IDToken,
		arg.TokenExpiresAt,
	)
	var i SocialAccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Platform,
		&i.Name,
		&i.Email,
		&i.AccessToken,
		&i.IDToken,
		&i.TokenExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LinkedinSub,
	)
	return i, err
}
