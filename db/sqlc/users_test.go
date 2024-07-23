package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	passwordHash := "this is some text"
	arg := CreateUserParams{
		Email:        "parth@gmail.com",
		PasswordHash: pgtype.Text{String: passwordHash, Valid: true},
	}
	fmt.Println(arg, "argggggg")
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
