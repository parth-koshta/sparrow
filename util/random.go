package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/rand"
)

const alphabet = "abcdefgijklmnopqrstuvwxyz"

func GenerateRandomString() pgtype.Text {
	n := chacha20poly1305.KeySize
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return pgtype.Text{String: sb.String(), Valid: true}
}

func GenerateRandomUUID() pgtype.UUID {
	id := uuid.New()
	return pgtype.UUID{Bytes: id, Valid: true}
}

func GenerateRandomEmail() string {
	randomInt := rand.Intn(1000)
	return fmt.Sprintf("user%d@gmail.com", randomInt)
}

func GenerateRandomPassword() pgtype.Text {
	randomInt := rand.Intn(1000)
	randomPassword := fmt.Sprintf("password%d", randomInt)
	hash := sha256.New()
	hash.Write([]byte(randomPassword))
	passwordHash := hex.EncodeToString(hash.Sum(nil))
	return pgtype.Text{String: passwordHash, Valid: true}
}
