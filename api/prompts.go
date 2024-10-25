package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type CreatePromptRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
	Text   string `json:"text" binding:"required"`
}

type PromptResponse struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	Text      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (server *Server) CreatePrompt(ctx *gin.Context) {
	var req CreatePromptRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	parsedUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePromptParams{
		UserID: pgtype.UUID{Bytes: parsedUUID, Valid: true},
		Text:   req.Text,
	}

	prompt, err := server.store.CreatePrompt(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PromptResponse{
		ID:        prompt.ID,
		UserID:    prompt.UserID,
		Text:      prompt.Text,
		CreatedAt: prompt.CreatedAt,
		UpdatedAt: prompt.UpdatedAt,
	})
}

type GetPromptRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) GetPrompt(ctx *gin.Context) {
	var req GetPromptRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	promptID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	prompt, err := server.store.GetPromptByID(ctx, pgtype.UUID{Bytes: promptID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PromptResponse{
		ID:        prompt.ID,
		UserID:    prompt.UserID,
		Text:      prompt.Text,
		CreatedAt: prompt.CreatedAt,
		UpdatedAt: prompt.UpdatedAt,
	})
}

type ListPromptsByUserIDRequest struct {
	UserID   string `form:"user_id" binding:"required,uuid"`
	Page     int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) ListPromptsByUserID(ctx *gin.Context) {
	var req ListPromptsByUserIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize

	prompts, err := server.store.ListPromptsByUserID(ctx, db.ListPromptsByUserIDParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, prompts)
}

type UpdatePromptRequest struct {
	ID   string `json:"id" binding:"required,uuid"`
	Text string `json:"text" binding:"required"`
}

func (server *Server) UpdatePrompt(ctx *gin.Context) {
	var req UpdatePromptRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	promptID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePromptParams{
		ID:   pgtype.UUID{Bytes: promptID, Valid: true},
		Text: req.Text,
	}

	prompt, err := server.store.UpdatePrompt(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PromptResponse{
		ID:        prompt.ID,
		UserID:    prompt.UserID,
		Text:      prompt.Text,
		CreatedAt: prompt.CreatedAt,
		UpdatedAt: prompt.UpdatedAt,
	})
}

type DeletePromptRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeletePrompt(ctx *gin.Context) {
	var req DeletePromptRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	promptID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeletePrompt(ctx, pgtype.UUID{Bytes: promptID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
