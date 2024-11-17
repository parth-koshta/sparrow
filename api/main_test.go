package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/mocks"
	"github.com/parth-koshta/sparrow/util"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.GenerateRandomString().String,
		AccessTokenDuration: time.Minute * 5,
	}
	mockTaskDistributor := mocks.NewTaskDistributor(t)
	mockLinkedinClient := mocks.NewLinkedinAPIClient(t)
	server, err := NewServer(store, config, mockTaskDistributor, mockLinkedinClient)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
