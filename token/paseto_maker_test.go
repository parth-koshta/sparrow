package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/parth-koshta/sparrow/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.GenerateRandomString().String)
	require.NoError(t, err)

	testUUID := uuid.New()
	testPgxUUID := pgtype.UUID{
		Bytes: testUUID,
		Valid: true,
	}
	email := util.GenerateRandomEmail()
	duration := time.Minute
	issueAt := time.Now()
	expiredAt := issueAt.Add(duration)

	token, err := maker.CreateToken(testPgxUUID, email, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, email, payload.Email)
	require.WithinDuration(t, issueAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.GenerateRandomString().String)
	require.NoError(t, err)

	email := util.GenerateRandomEmail()
	duration := -time.Minute
	testUUID := uuid.New()
	testPgxUUID := pgtype.UUID{
		Bytes: testUUID,
		Valid: true,
	}
	token, err := maker.CreateToken(testPgxUUID, email, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Empty(t, payload)
}
