package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/parth-koshta/sparrow/token"
)

const (
	AUTHORIZATION_HEADER_KEY  = "authorization"
	AUTHORIZATION_TYPE_BEARER = "bearer"
	AUTHORIZATION_PAYLOAD_KEY = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AUTHORIZATION_HEADER_KEY)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AUTHORIZATION_TYPE_BEARER {
			err := fmt.Errorf("unsuppored authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			err := errors.New("access token is invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(AUTHORIZATION_PAYLOAD_KEY, payload)
		ctx.Next()
	}
}

func GetUserIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	authPayload, exists := ctx.MustGet(AUTHORIZATION_PAYLOAD_KEY).(*token.Payload)
	if !exists || !authPayload.ID.Valid {
		return uuid.UUID{}, fmt.Errorf("invalid user ID in auth payload")
	}

	userID, err := uuid.FromBytes(authPayload.ID.Bytes[:])
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse user ID: %v", err)
	}

	return userID, nil
}
