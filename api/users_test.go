package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetUserApi(t *testing.T) {
	// Setup the test environment
	gin.SetMode(gin.TestMode)

	mockStore := mocks.NewStore(t)
	server := newTestServer(t, mockStore)

	// Generate a random UUID for testing
	testUUID := uuid.New()

	// Convert uuid.UUID to pgtype.UUID
	testPgxUUID := pgtype.UUID{
		Bytes: testUUID, // Direct conversion
		Valid: true,
	}

	// Create an expected response
	expectedUserRow := db.GetUserByIDRow{
		ID: testPgxUUID,
		Username: pgtype.Text{
			String: "testuser",
			Valid:  true,
		},
		Email: "testuser@example.com",
	}

	// Setup the mock expectation
	mockStore.On("GetUserByID", mock.Anything, testPgxUUID).Return(expectedUserRow, nil)

	// Create a test request
	req, err := http.NewRequest(http.MethodGet, "/users/"+testUUID.String(), nil)
	// set auth header
	addAuthorization(t, req, server.tokenMaker, authorizationTypeBearer, testPgxUUID, "user", server.config.AccessTokenDuration)
	require.NoError(t, err)

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Call the API endpoint
	server.router.ServeHTTP(recorder, req)

	// Check the response
	require.Equal(t, http.StatusOK, recorder.Code)
	mockStore.AssertNumberOfCalls(t, "GetUserByID", 1)
	mockStore.AssertExpectations(t)
}
