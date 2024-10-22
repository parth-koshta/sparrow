package api

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/parth-koshta/sparrow/client"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/token"
	"github.com/parth-koshta/sparrow/util"
)

type Server struct {
	store          db.Store
	tokenMaker     token.Maker
	config         util.Config
	router         *gin.Engine
	linkedInClient *client.LinkedInClient
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	linkedInClient := client.NewLinkedInClient(config.LinkedInClientID, config.LinkedInClientSecret, config.LinkedInRedirectURL)

	server := &Server{store: store, tokenMaker: tokenMaker, config: config, linkedInClient: linkedInClient}

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

	router.POST("/users/login", server.LoginUser)
	router.POST("/users", server.CreateUser)

	authenticatedRouter := router.Group("/").Use(AuthMiddleware(server.tokenMaker))
	authenticatedRouter.GET("/users", server.ListUsers)
	authenticatedRouter.GET("/users/:id", server.GetUser)

	authenticatedRouter.POST("/drafts", server.CreateDraft)
	authenticatedRouter.GET("/drafts/:id", server.getDraft)
	authenticatedRouter.PUT("/drafts/:id", server.updateDraft)
	authenticatedRouter.DELETE("/drafts/:id", server.deleteDraft)
	authenticatedRouter.GET("/drafts/user/:id", server.listDraftsByUserID)

	authenticatedRouter.POST("/prompts", server.CreatePrompt)
	authenticatedRouter.GET("/prompts/:id", server.getPrompt)
	authenticatedRouter.PUT("/prompts/:id", server.UpdatePrompt)
	authenticatedRouter.DELETE("/prompts/:id", server.DeletePrompt)
	authenticatedRouter.GET("/prompts/user/:id", server.ListPromptsByUserID)

	authenticatedRouter.POST("/suggestions", server.CreatePostSuggestion)
	authenticatedRouter.GET("/suggestions/:id", server.GetPostSuggestion)
	authenticatedRouter.PUT("/suggestions/:id", server.UpdatePostSuggestion)
	authenticatedRouter.DELETE("/suggestions/:id", server.DeletePostSuggestion)
	authenticatedRouter.GET("/suggestions/prompt/:id", server.ListPostSuggestionsByPromptID)

	authenticatedRouter.POST("/socialaccounts", server.CreateSocialAccount)
	authenticatedRouter.GET("/socialaccounts/:id", server.GetSocialAccount)
	authenticatedRouter.PUT("/socialaccounts/:id", server.UpdateSocialAccount)
	authenticatedRouter.DELETE("/socialaccounts/:id", server.DeleteSocialAccount)
	authenticatedRouter.GET("/socialaccounts/user/:id", server.listSocialAccountsByUserID)

	authenticatedRouter.POST("/posts", server.CreateScheduledPost)
	authenticatedRouter.GET("/posts/:id", server.GetScheduledPost)
	authenticatedRouter.PUT("/posts/:id", server.UpdateScheduledPost)
	authenticatedRouter.DELETE("/posts/:id", server.DeleteScheduledPost)
	authenticatedRouter.GET("/posts/user/:id", server.ListScheduledPostsByUserID)

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
