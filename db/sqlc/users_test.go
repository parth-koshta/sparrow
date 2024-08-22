package db

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
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
	})
}

func TestGetUserByEmail(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		arg := CreateUserParams{
			Email:        generateRandomEmail(),
			PasswordHash: generateRandomPasswordHash(),
		}
		createdUser, err := testQueries.CreateUser(context.Background(), arg)
		require.NoError(t, err)

		retrievedUser, err := testQueries.GetUserByEmail(context.Background(), arg.Email)
		require.NoError(t, err)
		require.NotEmpty(t, retrievedUser)

		require.Equal(t, createdUser.ID, retrievedUser.ID)
		require.Equal(t, createdUser.Email, retrievedUser.Email)
		require.WithinDuration(t, createdUser.CreatedAt.Time, retrievedUser.CreatedAt.Time, time.Second)
		require.WithinDuration(t, createdUser.UpdatedAt.Time, retrievedUser.UpdatedAt.Time, time.Second)
	})
}

func TestGetUserByID(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		arg := CreateUserParams{
			Email:        generateRandomEmail(),
			PasswordHash: generateRandomPasswordHash(),
		}
		createdUser, err := testQueries.CreateUser(context.Background(), arg)
		require.NoError(t, err)

		retrievedUser, err := testQueries.GetUserByID(context.Background(), createdUser.ID)
		require.NoError(t, err)
		require.NotEmpty(t, retrievedUser)

		require.Equal(t, createdUser.ID, retrievedUser.ID)
		require.Equal(t, createdUser.Email, retrievedUser.Email)
		require.WithinDuration(t, createdUser.CreatedAt.Time, retrievedUser.CreatedAt.Time, time.Second)
		require.WithinDuration(t, createdUser.UpdatedAt.Time, retrievedUser.UpdatedAt.Time, time.Second)
	})
}
func TestListUsers(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {

		// Create users
		for i := 0; i < 5; i++ {
			arg := CreateUserParams{
				Email:        generateRandomEmail(),
				PasswordHash: generateRandomPasswordHash(),
			}
			_, err := testQueries.CreateUser(context.Background(), arg)
			require.NoError(t, err)
		}

		// Paginate through users
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

		arg.Offset = 3
		users, err = testQueries.ListUsers(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, users, 2)

		arg.Offset = 5
		arg.Limit = 3
		users, err = testQueries.ListUsers(context.Background(), arg)
		require.NoError(t, err)
		require.Empty(t, users)
	})
}

func generateRandomUUID() pgtype.UUID {
	id := uuid.New()
	return pgtype.UUID{Bytes: id, Valid: true}
}

func runTestInTransaction(t *testing.T, testFunc func(*Queries)) {
	tx, err := testQueries.db.Begin(context.Background())
	require.NoError(t, err)
	defer tx.Rollback(context.Background())

	// Override testQueries to use the transaction
	testQueries := New(tx)
	testFunc(testQueries)
}

func generateRandomEmail() string {
	randomInt := rand.Intn(1000)
	return fmt.Sprintf("user%d@gmail.com", randomInt)
}

func generateRandomPasswordHash() pgtype.Text {
	randomInt := rand.Intn(1000)
	randomPassword := fmt.Sprintf("password%d", randomInt)
	hash := sha256.New()
	hash.Write([]byte(randomPassword))
	passwordHash := hex.EncodeToString(hash.Sum(nil))
	return pgtype.Text{String: passwordHash, Valid: true}
}
