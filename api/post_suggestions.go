package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type GetAISuggestionsRequest struct {
	Prompt string `form:"prompt" binding:"required"`
	Count  int    `form:"count" binding:"required,min=1,max=20"`
}

type GetAISuggestionsResponse struct {
	Suggestions []db.BulkCreatePostSuggestionsRow `json:"suggestions"`
}

func (s *Server) GetAISuggestionsByPrompt(ctx *gin.Context) {
	var req GetAISuggestionsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check if prompt already exists for this user
	getPromptArgs := db.GetPromptByUserIDAndTextParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Text:   req.Prompt,
	}
	existingPrompt, err := s.store.GetPromptByUserIDAndText(ctx, getPromptArgs)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	var promptID pgtype.UUID
	if existingPrompt.ID.Valid {
		promptID = existingPrompt.ID
	} else {
		promptArg := db.CreatePromptParams{
			UserID: pgtype.UUID{Bytes: userID, Valid: true},
			Text:   req.Prompt,
		}

		// Save the prompt to the database
		prompt, err := s.store.CreatePrompt(ctx, promptArg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		promptID = prompt.ID
	}

	// Step 2: Generate suggestions using the OpenAI client
	suggestions, err := s.openaiClient.GenerateLinkedInPosts(req.Prompt, req.Count)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Step 3: Save the suggestions in the database
	bulkArgs := db.BulkCreatePostSuggestionsParams{
		PromptID:    promptID,
		Suggestions: suggestions,
	}

	// Get the saved suggestions from the database
	savedSuggestions, err := s.store.BulkCreatePostSuggestions(ctx, bulkArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Step 4: Return the saved suggestions to the client
	ctx.JSON(http.StatusOK, GetAISuggestionsResponse{
		Suggestions: savedSuggestions,
	})
}

type CreatePostSuggestionRequest struct {
	PromptID string `json:"prompt_id" binding:"required,uuid"`
	Text     string `json:"text" binding:"required"`
}

type PostSuggestionResponse struct {
	ID        pgtype.UUID
	PromptID  pgtype.UUID
	Text      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (server *Server) CreatePostSuggestion(ctx *gin.Context) {
	var req CreatePostSuggestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	parsedUUID, err := uuid.Parse(req.PromptID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostSuggestionParams{
		PromptID: pgtype.UUID{Bytes: parsedUUID, Valid: true},
		Text:     req.Text,
	}

	suggestion, err := server.store.CreatePostSuggestion(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PostSuggestionResponse{
		ID:        suggestion.ID,
		PromptID:  suggestion.PromptID,
		Text:      suggestion.Text,
		CreatedAt: suggestion.CreatedAt,
		UpdatedAt: suggestion.UpdatedAt,
	})
}

type ListPostSuggestionsByPromptIDRequest struct {
	PromptID string `form:"prompt_id" binding:"required,uuid"`
	Page     int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) ListPostSuggestionsByPromptID(ctx *gin.Context) {
	var req ListPostSuggestionsByPromptIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	promptID, err := uuid.Parse(req.PromptID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize

	suggestions, err := server.store.ListPostSuggestionsByPromptID(ctx, db.ListPostSuggestionsByPromptIDParams{
		PromptID: pgtype.UUID{Bytes: promptID, Valid: true},
		Limit:    limit,
		Offset:   offset,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, suggestions)
}

type DeletePostSuggestionRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeletePostSuggestion(ctx *gin.Context) {
	var req DeletePostSuggestionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	suggestionID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeletePostSuggestion(ctx, pgtype.UUID{Bytes: suggestionID, Valid: true})
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNoContent, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

type AcceptPostSuggestionRequest struct {
	ID string `json:"id" binding:"required,uuid"`
}
type AcceptPostSuggestionResponse struct {
	PostID pgtype.UUID
}

func (s *Server) AcceptPostSuggestion(ctx *gin.Context) {
	var req AcceptPostSuggestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	suggestionID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var result db.AcceptPostSuggestionTxResult
	arg := db.AcceptPostSuggestionTxParams{
		UserID:       pgtype.UUID{Bytes: userID, Valid: true},
		SuggestionID: pgtype.UUID{Bytes: suggestionID, Valid: true},
	}

	result, err = s.store.AcceptPostSuggestionTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, AcceptPostSuggestionResponse{
		PostID: result.PostID,
	})
}
