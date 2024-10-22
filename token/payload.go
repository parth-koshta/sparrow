package token

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrInvalidToken = errors.New("token is invalid")  //fmt.Errorf("token is invalid")
	ErrExpiredToken = errors.New("token has expired") //fmt.Errorf("token has expored")
)

type Payload struct {
	ID        pgtype.UUID `json:"id"`
	Email     string      `json:"email"`
	IssuedAt  time.Time   `json:"issuedAt"`
	ExpiredAt time.Time   `json:"expiredAt"`
}

func NewPayload(id pgtype.UUID, email string, duration time.Duration) (*Payload, error) {
	return &Payload{
		ID:        id,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
