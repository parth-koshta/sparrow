package db

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/parth-koshta/sparrow/util"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

const (
	dbSource       = "postgresql://root:secret@localhost:5432/"
	testDBNameBase = "sparrow_test"
	migrationDir   = "../migration"
)

var testQueries *Queries
var testDBPool *pgxpool.Pool

func runMigrations(dbName string) error {
	cmd := exec.Command("migrate", "-path", migrationDir, "-database", fmt.Sprintf("postgresql://root:secret@localhost:5432/%s?sslmode=disable", dbName), "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Generate a unique database name
	testDBName := fmt.Sprintf("%s_%d", testDBNameBase, time.Now().UnixNano())

	// Connect to the main database to create a test database
	mainConn, err := pgxpool.New(ctx, dbSource+"postgres?sslmode=disable")
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to main db")
	}

	// Create the test database
	_, err = mainConn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", testDBName))
	if err != nil {
		log.Error().Err(err).Msg("cannot create test db")
	}

	// Close the main connection
	mainConn.Close()

	// Connect to the test database
	testDBPool, err = pgxpool.New(ctx, dbSource+testDBName+"?sslmode=disable")
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to test db")
	}

	testQueries = New(testDBPool)

	// Run migrations on the test database
	if err := runMigrations(testDBName); err != nil {
		log.Error().Err(err).Msg("failed to run migrations")
	}

	// Run tests
	code := m.Run()

	// Cleanup: drop the test database
	testDBPool.Close()
	mainConn, err = pgxpool.New(ctx, dbSource+"postgres?sslmode=disable")
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to main db for cleanup")
	}
	_, err = mainConn.Exec(ctx, fmt.Sprintf("DROP DATABASE %s", testDBName))
	if err != nil {
		log.Error().Err(err).Msg("cannot drop test db")
	}
	mainConn.Close()

	os.Exit(code)
}

func runTestInTransaction(t *testing.T, testFunc func(*Queries)) {
	store := NewStore(testDBPool)
	sqlStore, ok := store.(*SQLStore)
	require.True(t, ok)

	// Begin a transaction
	tx, err := sqlStore.db.Begin(context.Background())
	require.NoError(t, err)

	// Create a Queries object tied to this transaction
	q := New(tx)

	// Run the test function
	testFunc(q)

	// Rollback the transaction after the test is complete
	err = tx.Rollback(context.Background())
	require.NoError(t, err)
}

func createRandomUser(t *testing.T, testQueries *Queries) User {
	arg := CreateUserParams{
		Email:        util.GenerateRandomEmail(),
		PasswordHash: util.GenerateRandomPassword(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	return user
}

func createRandomPrompt(t *testing.T, testQueries *Queries, userID pgtype.UUID) Prompt {
	arg := CreatePromptParams{
		UserID: userID,
		Text:   "Example prompt text",
	}
	prompt, err := testQueries.CreatePrompt(context.Background(), arg)
	require.NoError(t, err)
	return prompt
}

func createRandomPostSuggestion(t *testing.T, testQueries *Queries, promptID pgtype.UUID) PostSuggestion {
	arg := CreatePostSuggestionParams{
		PromptID: promptID,
		Text:     fmt.Sprintf("Example suggestion text %s", util.GenerateRandomString().String),
	}
	suggestion, err := testQueries.CreatePostSuggestion(context.Background(), arg)
	require.NoError(t, err)
	return suggestion
}

func createRandomDraft(t *testing.T, testQueries *Queries, userID pgtype.UUID, suggestionID pgtype.UUID) Post {
	arg := CreatePostParams{
		UserID:       userID,
		SuggestionID: suggestionID,
		Text:         "Example draft text",
		Status:       "draft",
	}
	draft, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	return draft
}

func createRandomScheduledPost(t *testing.T, testQueries *Queries, userID pgtype.UUID, draftID pgtype.UUID) PostSchedule {
	arg := CreatePostScheduleParams{
		UserID:        userID,
		PostID:        draftID,
		ScheduledTime: pgtype.Timestamp{Time: time.Now().Add(24 * time.Hour), Valid: true},
		Status:        "scheduled",
	}
	post, err := testQueries.CreatePostSchedule(context.Background(), arg)
	require.NoError(t, err)
	return post
}

func createRandomSocialAccount(t *testing.T, testQueries *Queries, userID pgtype.UUID) SocialAccount {
	arg := CreateSocialAccountParams{
		UserID:         userID,
		Platform:       "ExamplePlatform",
		AccountName:    "ExampleAccountName",
		AccessToken:    "ExampleAccessToken",
		AccountEmail:   "example@gmail.com",
		IDToken:        "ExampleIDToken",
		TokenExpiresAt: pgtype.Timestamp{Time: time.Now().Add(24 * time.Hour), Valid: true},
	}
	socialAccount, err := testQueries.CreateSocialAccount(context.Background(), arg)
	require.NoError(t, err)
	return socialAccount
}
