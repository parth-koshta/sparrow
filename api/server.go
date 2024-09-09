package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/token"
	"github.com/parth-koshta/sparrow/util"
)

type Server struct {
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
	router     *gin.Engine
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users/login", server.loginUser)
	router.POST("/users", server.createUser)

	authenticatedRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authenticatedRouter.GET("/users", server.listUsers)
	authenticatedRouter.GET("/users/:id", server.getUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
