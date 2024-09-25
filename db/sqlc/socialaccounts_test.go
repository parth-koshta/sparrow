package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSocialAccount(t *testing.T) {
	runTestInTransaction(t, func(testQueries *Queries) {
		// Create a user first
		user := createRandomUser(t, testQueries)

		arg := CreateSocialAccountParams{
			UserID:      user.ID,
			Platform:    "twitter",
			AccountName: "example_account",
			AccessToken: "sample_access_token",
		}
		socialAccount, err := testQueries.CreateSocialAccount(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, socialAccount)

		require.Equal(t, arg.UserID, socialAccount.UserID)
		require.Equal(t, arg.Platform, socialAccount.Platform)
		require.Equal(t, arg.AccountName, socialAccount.AccountName)
		require.Equal(t, arg.AccessToken, socialAccount.AccessToken)
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
