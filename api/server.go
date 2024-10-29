package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	sentrygin "github.com/getsentry/sentry-go/gin"
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
	httpServer      *http.Server
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

	server.httpServer = &http.Server{
		Addr:    config.ServerAddress,
		Handler: server.router,
	}

	return server, nil
}

func (server *Server) setupRouter() {
	gin.SetMode(server.config.GinMode)
	router := gin.Default()
	router.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	router.Use(LoggerMiddleware)

	router.GET("/", server.HealthCheck)

	router.POST("/v1/users/login", server.LoginUser)
	router.POST("/v1/users", server.CreateUser)
	router.GET("/v1/users/verify/email", server.VerifyUserEmail)

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
	authenticatedRouter.DELETE("/v1/suggestions/:id", server.DeletePostSuggestion)
	authenticatedRouter.GET("/v1/suggestions/prompt/:id", server.ListPostSuggestionsByPromptID)
	authenticatedRouter.POST("/v1/suggestions/ai", server.GetAISuggestionsByPrompt)
	authenticatedRouter.POST("/v1/suggestions/accept", server.AcceptPostSuggestion)

	authenticatedRouter.GET("/v1/socialaccounts/:id", server.GetSocialAccount)
	authenticatedRouter.DELETE("/v1/socialaccounts/:id", server.DeleteSocialAccount)
	authenticatedRouter.GET("/v1/socialaccounts/user/:id", server.ListSocialAccountsByUserID)
	authenticatedRouter.POST("/v1/socialaccounts/linkedin", server.AddLinkedInAccount)
	authenticatedRouter.PUT("/v1/socialaccounts/accesstoken/linkedin", server.UpdateLinkedInAccessToken)

	authenticatedRouter.POST("/v1/schedules", server.CreatePostSchedule)
	authenticatedRouter.GET("/v1/schedules/:id", server.GetPostSchedule)
	authenticatedRouter.DELETE("/v1/schedules/:id", server.DeletePostSchedule)

	server.router = router
}

func (server *Server) Start(address string) error {
	err := server.httpServer.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		log.Error().Err(err).Msg("HTTP server failed to serve")
		return err
	}
	return nil
}

func (server *Server) Stop(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := server.httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("server shutdown error")
		return err
	}
	return nil
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
