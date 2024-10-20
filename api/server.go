package api

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
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

	err = server.initializeSentry()
	if err != nil {
		return nil, err
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(sentrygin.New(sentrygin.Options{Repanic: true}))

	router.POST("/users/login", server.loginUser)
	router.POST("/users", server.createUser)

	authenticatedRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authenticatedRouter.GET("/users", server.listUsers)
	authenticatedRouter.GET("/users/:id", server.getUser)

	authenticatedRouter.POST("/drafts", server.createDraft)
	authenticatedRouter.GET("/drafts/:id", server.getDraft)
	authenticatedRouter.PUT("/drafts/:id", server.updateDraft)
	authenticatedRouter.DELETE("/drafts/:id", server.deleteDraft)
	authenticatedRouter.GET("/drafts/user/:id", server.listDraftsByUserID)

	authenticatedRouter.POST("/prompts", server.createPrompt)
	authenticatedRouter.GET("/prompts/:id", server.getPrompt)
	authenticatedRouter.PUT("/prompts/:id", server.updatePrompt)
	authenticatedRouter.DELETE("/prompts/:id", server.deletePrompt)
	authenticatedRouter.GET("/prompts/user/:id", server.listPromptsByUserID)

	authenticatedRouter.POST("/suggestions", server.createPostSuggestion)
	authenticatedRouter.GET("/suggestions/:id", server.getPostSuggestion)
	authenticatedRouter.PUT("/suggestions/:id", server.updatePostSuggestion)
	authenticatedRouter.DELETE("/suggestions/:id", server.deletePostSuggestion)
	authenticatedRouter.GET("/suggestions/prompt/:id", server.listPostSuggestionsByPromptID)

	authenticatedRouter.POST("/socialaccounts", server.createSocialAccount)
	authenticatedRouter.GET("/socialaccounts/:id", server.getSocialAccount)
	authenticatedRouter.PUT("/socialaccounts/:id", server.updateSocialAccount)
	authenticatedRouter.DELETE("/socialaccounts/:id", server.deleteSocialAccount)
	authenticatedRouter.GET("/socialaccounts/user/:id", server.listSocialAccountsByUserID)

	authenticatedRouter.POST("/posts", server.createScheduledPost)
	authenticatedRouter.GET("/posts/:id", server.getScheduledPost)
	authenticatedRouter.PUT("/posts/:id", server.updateScheduledPost)
	authenticatedRouter.DELETE("/posts/:id", server.deleteScheduledPost)
	authenticatedRouter.GET("/posts/user/:id", server.listScheduledPostsByUserID)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) initializeSentry() error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://c1b68ed689e46da2852f882751206e84@o4508154405847040.ingest.us.sentry.io/4508154407550976",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})

	return err
}
