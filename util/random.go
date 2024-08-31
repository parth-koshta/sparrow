package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/rand"
)

func GenerateRandomUUID() pgtype.UUID {
	id := uuid.New()
	return pgtype.UUID{Bytes: id, Valid: true}
}

func GenerateRandomEmail() string {
	randomInt := rand.Intn(1000)
	return fmt.Sprintf("user%d@gmail.com", randomInt)
}

func GenerateRandomPasswordHash() pgtype.Text {
	randomInt := rand.Intn(1000)
	randomPassword := fmt.Sprintf("password%d", randomInt)
	hash := sha256.New()
	hash.Write([]byte(randomPassword))
	passwordHash := hex.EncodeToString(hash.Sum(nil))
	return pgtype.Text{String: passwordHash, Valid: true}
}
