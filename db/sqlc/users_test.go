package db

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func generateRandomEmail() string {
	// rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(1000)
	return fmt.Sprintf("user%d@gmail.com", randomInt)
}

func generateRandomPasswordHash() pgtype.Text {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(1000)
	randomPassword := fmt.Sprintf("password%d", randomInt)
	hash := sha256.New()
	hash.Write([]byte(randomPassword))
	passwordHash := hex.EncodeToString(hash.Sum(nil))
	return pgtype.Text{String: passwordHash, Valid: true}
}

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Email:        generateRandomEmail(),
		PasswordHash: generateRandomPasswordHash(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
