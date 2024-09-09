package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPas(t *testing.T) {
	password := GenerateRandomPassword()

	hashedPassword, err := HashPassword(password.String)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(hashedPassword, password.String)
	require.NoError(t, err)

	incorrectPassword := GenerateRandomPassword()
	err = CheckPassword(hashedPassword, incorrectPassword.String)
	require.NotEmpty(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
