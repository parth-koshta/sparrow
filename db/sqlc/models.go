// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	ID           pgtype.UUID      `json:"id"`
	UserID       pgtype.UUID      `json:"user_id"`
	SuggestionID pgtype.UUID      `json:"suggestion_id"`
	Text         string           `json:"text"`
	Status       string           `json:"status"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}

type PostSchedule struct {
	ID              pgtype.UUID      `json:"id"`
	UserID          pgtype.UUID      `json:"user_id"`
	PostID          pgtype.UUID      `json:"post_id"`
	SocialAccountID pgtype.UUID      `json:"social_account_id"`
	ScheduledTime   pgtype.Timestamp `json:"scheduled_time"`
	ExecutedTime    pgtype.Timestamp `json:"executed_time"`
	Status          string           `json:"status"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
}

type PostSuggestion struct {
	ID        pgtype.UUID      `json:"id"`
	PromptID  pgtype.UUID      `json:"prompt_id"`
	Text      string           `json:"text"`
	Status    string           `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type Prompt struct {
	ID        pgtype.UUID      `json:"id"`
	UserID    pgtype.UUID      `json:"user_id"`
	Text      string           `json:"text"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type SocialAccount struct {
	ID             pgtype.UUID      `json:"id"`
	UserID         pgtype.UUID      `json:"user_id"`
	Platform       string           `json:"platform"`
	AccountName    string           `json:"account_name"`
	AccountEmail   string           `json:"account_email"`
	AccessToken    string           `json:"access_token"`
	IDToken        string           `json:"id_token"`
	TokenExpiresAt pgtype.Timestamp `json:"token_expires_at"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
	LinkedinSub    pgtype.Text      `json:"linkedin_sub"`
}

type User struct {
	ID              pgtype.UUID      `json:"id"`
	Username        pgtype.Text      `json:"username"`
	Email           string           `json:"email"`
	PasswordHash    pgtype.Text      `json:"password_hash"`
	IsEmailVerified bool             `json:"is_email_verified"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
}

type VerifyEmail struct {
	ID         int64            `json:"id"`
	Email      string           `json:"email"`
	SecretCode string           `json:"secret_code"`
	IsUsed     bool             `json:"is_used"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	ExpiredAt  pgtype.Timestamp `json:"expired_at"`
}
