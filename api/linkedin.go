package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetLinkedInAccessTokenRequest struct {
	Code string `form:"code" binding:"required"`
}

func (server *Server) GetLinkedInAccessToken(ctx *gin.Context) {
	var req GetLinkedInAccessTokenRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code is required"})
		return
	}

	accessToken, err := server.linkedInClient.GetAccessToken(req.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get access token", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
