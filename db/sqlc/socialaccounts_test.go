package db

import (
	"context"
	"testing"
	"time"

	pgtype "github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateSocialAccount(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user first
		user := createRandomUser(t, testQueries)

		arg := CreateSocialAccountParams{
			UserID:         user.ID,
			Platform:       "twitter",
			Name:           "example_account",
			AccessToken:    "sample_access_token",
			Email:          "test@gmail.com",
			IDToken:        "sample_id_token",
			TokenExpiresAt: pgtype.Timestamp{Time: time.Now().Add(24 * time.Hour).UTC(), Valid: true},
		}
		socialAccount, err := testQueries.CreateSocialAccount(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, socialAccount)

		require.Equal(t, arg.UserID, socialAccount.UserID)
		require.Equal(t, arg.Platform, socialAccount.Platform)
		require.Equal(t, arg.Name, socialAccount.Name)
		require.Equal(t, arg.AccessToken, socialAccount.AccessToken)
		require.Equal(t, arg.Email, socialAccount.Email)
		require.Equal(t, arg.IDToken, socialAccount.IDToken)
		require.Equal(t, arg.TokenExpiresAt.Time.Truncate(time.Second), socialAccount.TokenExpiresAt.Time.Truncate(time.Second))
		require.NotZero(t, socialAccount.CreatedAt)
	})
}

func TestListSocialAccountsByUserID(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user and add multiple social accounts
		user := createRandomUser(t, testQueries)
		for i := 0; i < 5; i++ {
			createRandomSocialAccount(t, testQueries, user.ID)
		}

		arg := ListSocialAccountsByUserIDParams{
			UserID: user.ID,
			Limit:  3,
			Offset: 0,
		}
		socialAccounts, err := testQueries.ListSocialAccountsByUserID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, socialAccounts, 3)

		arg.Offset = 3
		socialAccounts, err = testQueries.ListSocialAccountsByUserID(context.Background(), arg)
		require.NoError(t, err)
		require.Len(t, socialAccounts, 2)
	})
}
