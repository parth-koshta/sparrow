package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
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
	linkedinClient  client.LinkedinAPIClient
	openaiClient    *client.OpenAIClient
	taskDistributor worker.TaskDistributor
}

func NewServer(store db.Store, config util.Config, taskDistributor worker.TaskDistributor, linkedinClient client.LinkedinAPIClient) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"}, // Frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},        // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization"},                  // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                                 // Headers exposed to the client
		AllowCredentials: true,                                                       // Allow credentials like cookies
	}))

	router.GET("/", server.HealthCheck)

	router.POST("/v1/users/login", server.LoginUser)
	router.POST("/v1/users", server.CreateUser)
	router.GET("/v1/users/verify/email", server.VerifyUserEmail)
	router.POST("/v1/users/email/resend", server.ResendVerifyEmail)
	router.POST("v1/users/email/verified", server.IsEmailVerified)

	authenticatedRouter := router.Group("/").Use(AuthMiddleware(server.tokenMaker))
	authenticatedRouter.GET("/v1/users", server.ListUsers)
	authenticatedRouter.GET("/v1/users/:id", server.GetUser)

	authenticatedRouter.PUT("/v1/posts/:id", server.UpdatePost)
	authenticatedRouter.DELETE("/v1/posts/:id", server.DeletePost)
	authenticatedRouter.GET("/v1/posts", server.ListPostsByUserID)
	authenticatedRouter.POST("/v1/posts/publish/linkedin", server.PublishOnLinkedIn)

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
	authenticatedRouter.POST("/v1/suggestions/reject", server.RejectPostSuggestion)

	authenticatedRouter.GET("/v1/social_accounts/:id", server.GetSocialAccount)
	authenticatedRouter.DELETE("/v1/social_accounts/:id", server.DeleteSocialAccount)
	authenticatedRouter.GET("/v1/social_accounts/user/:id", server.ListSocialAccountsByUserID)
	authenticatedRouter.POST("/v1/social_accounts/linkedin", server.AddLinkedInAccount)
	authenticatedRouter.PUT("/v1/social_accounts/accesstoken/linkedin", server.UpdateLinkedInAccessToken)

	authenticatedRouter.POST("/v1/schedules", server.SchedulePost)
	authenticatedRouter.GET("/v1/schedules/:id", server.GetSchedule)
	authenticatedRouter.GET("/v1/schedules", server.ListPostSchedulesForUser)
	authenticatedRouter.DELETE("/v1/schedules/:id", server.DeleteSchedule)

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
	return gin.H{"message": err.Error()}
}

func customErrorResponse(err error, message string) gin.H {
	if err == nil {
		return gin.H{"message": message}
	}
	return gin.H{"message": message, "error": err.Error()}
}

func (server *Server) initializeSentry() error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://c1b68ed689e46da2852f882751206e84@o4508154405847040.ingest.us.sentry.io/4508154407550976",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})

	return err
}
