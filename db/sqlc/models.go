// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	ID           pgtype.UUID
	UserID       pgtype.UUID
	SuggestionID pgtype.UUID
	Text         string
	Status       string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type Postschedule struct {
	ID            pgtype.UUID
	UserID        pgtype.UUID
	PostID        pgtype.UUID
	ScheduledTime pgtype.Timestamp
	ExecutedTime  pgtype.Timestamp
	Status        string
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
}

type Postsuggestion struct {
	ID        pgtype.UUID
	PromptID  pgtype.UUID
	Text      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Prompt struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	Text      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Socialaccount struct {
	ID             pgtype.UUID
	UserID         pgtype.UUID
	Platform       string
	AccountName    string
	AccountEmail   string
	AccessToken    string
	IDToken        string
	TokenExpiresAt pgtype.Timestamp
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

type User struct {
	ID              pgtype.UUID
	Username        pgtype.Text
	Email           string
	PasswordHash    pgtype.Text
	IsEmailVerified bool
	CreatedAt       pgtype.Timestamp
	UpdatedAt       pgtype.Timestamp
}

type Verifyemail struct {
	ID         int64
	Email      string
	SecretCode string
	IsUsed     bool
	CreatedAt  pgtype.Timestamp
	ExpiredAt  pgtype.Timestamp
}
