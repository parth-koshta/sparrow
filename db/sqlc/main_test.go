package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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
		log.Fatal("cannot connect to main db: ", err)
	}

	// Create the test database
	_, err = mainConn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", testDBName))
	if err != nil {
		log.Fatal("cannot create test db: ", err)
	}

	// Close the main connection
	mainConn.Close()

	// Connect to the test database
	testDBPool, err = pgxpool.New(ctx, dbSource+testDBName+"?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to test db: ", err)
	}

	testQueries = New(testDBPool)

	// Run migrations on the test database
	if err := runMigrations(testDBName); err != nil {
		log.Fatal("failed to run migrations: ", err)
	}

	// Run tests
	code := m.Run()

	// Cleanup: drop the test database
	testDBPool.Close()
	mainConn, err = pgxpool.New(ctx, dbSource+"postgres?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to main db for cleanup: ", err)
	}
	_, err = mainConn.Exec(ctx, fmt.Sprintf("DROP DATABASE %s", testDBName))
	if err != nil {
		log.Fatal("cannot drop test db: ", err)
	}
	mainConn.Close()

	os.Exit(code)
}
