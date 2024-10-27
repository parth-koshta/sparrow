package api

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/parth-koshta/sparrow/client"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/token"
	"github.com/parth-koshta/sparrow/util"
	"github.com/parth-koshta/sparrow/worker"
)

type Server struct {
	store           db.Store
	tokenMaker      token.Maker
	config          util.Config
	router          *gin.Engine
	linkedinClient  *client.LinkedinClient
	openaiClient    *client.OpenAIClient
	taskDistributor worker.TaskDistributor
}

func NewServer(store db.Store, config util.Config, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	linkedinClient := client.NewLinkedInClient(config.LinkedInClientID, config.LinkedInClientSecret)
	openaiClient := client.NewOpenAIClient(config.OpenAIApiKey)

	server := &Server{store: store, tokenMaker: tokenMaker, config: config, linkedinClient: linkedinClient, openaiClient: openaiClient, taskDistributor: taskDistributor}

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
	router.Use(LoggerMiddleware)

	router.GET("/", server.HealthCheck)

	router.POST("/users/login", server.LoginUser)
	router.POST("/users", server.CreateUser)

	authenticatedRouter := router.Group("/").Use(AuthMiddleware(server.tokenMaker))
	authenticatedRouter.GET("/v1/users", server.ListUsers)
	authenticatedRouter.GET("/v1/users/:id", server.GetUser)

	// authenticatedRouter.POST("/drafts", server.CreateDraft)
	// authenticatedRouter.GET("/drafts/:id", server.GetDraft)
	// authenticatedRouter.PUT("/drafts/:id", server.UpdateDraft)
	// authenticatedRouter.DELETE("/drafts/:id", server.DeleteDraft)
	// authenticatedRouter.GET("/drafts/user/:id", server.ListDraftsByUserID)

	authenticatedRouter.POST("/v1/prompts", server.CreatePrompt)
	authenticatedRouter.GET("/v1/prompts/:id", server.GetPrompt)
	authenticatedRouter.PUT("/v1/prompts/:id", server.UpdatePrompt)
	authenticatedRouter.DELETE("/v1/prompts/:id", server.DeletePrompt)
	authenticatedRouter.GET("/v1/prompts/user/:id", server.ListPromptsByUserID)

	authenticatedRouter.POST("/v1/suggestions", server.CreatePostSuggestion)
	authenticatedRouter.GET("/v1/suggestions/:id", server.GetPostSuggestion)
	authenticatedRouter.PUT("/v1/suggestions/:id", server.UpdatePostSuggestion)
	authenticatedRouter.DELETE("/v1/suggestions/:id", server.DeletePostSuggestion)
	authenticatedRouter.GET("/v1/suggestions/prompt/:id", server.ListPostSuggestionsByPromptID)
	authenticatedRouter.POST("/v1/suggestions/ai", server.GetAISuggestionsByPrompt)

	authenticatedRouter.GET("/v1/socialaccounts/:id", server.GetSocialAccount)
	authenticatedRouter.DELETE("/v1/socialaccounts/:id", server.DeleteSocialAccount)
	authenticatedRouter.GET("/v1/socialaccounts/user/:id", server.ListSocialAccountsByUserID)
	authenticatedRouter.POST("/v1/socialaccounts/linkedin", server.AddLinkedInAccount)
	authenticatedRouter.PUT("/v1/socialaccounts/accesstoken/linkedin", server.UpdateLinkedInAccessToken)

	authenticatedRouter.POST("/v1/schedules", server.CreateScheduledPost)
	authenticatedRouter.GET("/v1/schedules/:id", server.GetScheduledPost)
	authenticatedRouter.PUT("/v1/schedules/:id", server.UpdateScheduledPost)
	authenticatedRouter.DELETE("/v1/schedules/:id", server.DeleteScheduledPost)
	authenticatedRouter.GET("/v1/schedules/user/:id", server.ListScheduledPostsByUserID)

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
