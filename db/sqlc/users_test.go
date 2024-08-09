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

func TestGetUserByEmail(t *testing.T) {
	// First, create a new user
	arg := CreateUserParams{
		Email:        generateRandomEmail(),
		PasswordHash: generateRandomPasswordHash(),
	}
	createdUser, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	// Retrieve the user by email
	retrievedUser, err := testQueries.GetUserByEmail(context.Background(), arg.Email)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	// Check if the retrieved user's details match what was created
	require.Equal(t, createdUser.ID, retrievedUser.ID)
	require.Equal(t, createdUser.Email, retrievedUser.Email)
	require.Equal(t, createdUser.Username, retrievedUser.Username)
	require.WithinDuration(t, createdUser.CreatedAt.Time, retrievedUser.CreatedAt.Time, time.Second)
	require.WithinDuration(t, createdUser.UpdatedAt.Time, retrievedUser.UpdatedAt.Time, time.Second)
}

func TestGetUserByID(t *testing.T) {
	// First, create a new user
	arg := CreateUserParams{
		Email:        generateRandomEmail(),
		PasswordHash: generateRandomPasswordHash(),
	}
	createdUser, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	// Retrieve the user by ID
	retrievedUser, err := testQueries.GetUserByID(context.Background(), createdUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	// Check if the retrieved user's details match what was created
	require.Equal(t, createdUser.ID, retrievedUser.ID)
	require.Equal(t, createdUser.Email, retrievedUser.Email)
	require.Equal(t, createdUser.Username, retrievedUser.Username)
	require.WithinDuration(t, createdUser.CreatedAt.Time, retrievedUser.CreatedAt.Time, time.Second)
	require.WithinDuration(t, createdUser.UpdatedAt.Time, retrievedUser.UpdatedAt.Time, time.Second)
}

func TestListUsers(t *testing.T) {
	// Create a few users for testing
	for i := 0; i < 5; i++ {
		arg := CreateUserParams{
			Email:        generateRandomEmail(),
			PasswordHash: generateRandomPasswordHash(),
		}
		_, err := testQueries.CreateUser(context.Background(), arg)
		require.NoError(t, err)
	}

	// Define pagination parameters
	arg := ListUsersParams{
		Limit:  3,
		Offset: 0,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 3)

	for _, user := range users {
		require.NotEmpty(t, user)
		require.NotZero(t, user.ID)
		require.NotZero(t, user.CreatedAt)
	}
}
