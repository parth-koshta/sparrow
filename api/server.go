package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/users/:id", server.getUser)
	router.POST("/users", server.createUser)
	router.GET("/users", server.listUsers)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
