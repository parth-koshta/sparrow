package token

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Maker interface {
	CreateToken(id pgtype.UUID, email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
