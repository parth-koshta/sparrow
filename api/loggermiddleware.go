package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// LoggerMiddleware logs each HTTP request directly as a gin.HandlerFunc
func LoggerMiddleware(c *gin.Context) {
	start := time.Now()

	// Process the request
	c.Next()

	// Log request details after response is written
	log.Info().
		Str("method", c.Request.Method).
		Str("url", c.Request.URL.String()).
		Int("status", c.Writer.Status()).
		Dur("duration", time.Since(start)).
		Msg("HTTP request")
}
