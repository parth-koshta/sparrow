package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/parth-koshta/sparrow/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	hashedPassword, err := util.HashPassword(util.GenerateRandomPassword().String)
	require.NoError(t, err)
	runTestInTransaction(t, func(testQueries *Queries) {
		arg := CreateUserParams{
			Email: util.GenerateRandomEmail(),
			PasswordHash: pgtype.Text{
				String: hashedPassword,
				Valid:  true,
			},
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
			Email:        util.GenerateRandomEmail(),
			PasswordHash: util.GenerateRandomPassword(),
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
			Email:        util.GenerateRandomEmail(),
			PasswordHash: util.GenerateRandomPassword(),
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
				Email:        util.GenerateRandomEmail(),
				PasswordHash: util.GenerateRandomPassword(),
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

// func runTestInTransaction(t *testing.T, testFunc func(*Queries)) {
// 	store := NewStore(testDBPool)
// 	sqlStore, ok := store.(*SQLStore)
// 	require.True(t, ok)

// 	// Begin a transaction
// 	tx, err := sqlStore.db.Begin(context.Background())
// 	require.NoError(t, err)

// 	// Create a Queries object tied to this transaction
// 	q := New(tx)

// 	// Run the test function
// 	testFunc(q)

// 	// Rollback the transaction after the test is complete
// 	err = tx.Rollback(context.Background())
// 	require.NoError(t, err)
// }
