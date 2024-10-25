package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type CreatePostRequest struct {
	UserID       string `json:"user_id" binding:"required,uuid"`
	SuggestionID string `json:"suggestion_id" binding:"required,uuid"`
	Text         string `json:"text" binding:"required"` // Changed from DraftText to Text
}

type PostResponse struct {
	ID           pgtype.UUID
	UserID       pgtype.UUID
	SuggestionID pgtype.UUID
	Text         string // Changed from DraftText to Text
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

func (server *Server) CreatePost(ctx *gin.Context) {
	var req CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	suggestionID, err := uuid.Parse(req.SuggestionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{ // Changed from CreateDraftParams to CreatePostParams
		UserID:       pgtype.UUID{Bytes: userID, Valid: true},
		SuggestionID: pgtype.UUID{Bytes: suggestionID, Valid: true},
		Text:         req.Text, // Changed from DraftText to Text
	}

	post, err := server.store.CreatePost(ctx, arg) // Changed from CreateDraft to CreatePost
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PostResponse{ // Changed from DraftResponse to PostResponse
		ID:           post.ID,
		UserID:       post.UserID,
		SuggestionID: post.SuggestionID,
		Text:         post.Text, // Changed from DraftText to Text
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	})
}

type GetPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) GetPost(ctx *gin.Context) { // Changed from GetDraft to GetPost
	var req GetPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := server.store.GetPostByID(ctx, pgtype.UUID{Bytes: postID, Valid: true}) // Changed from GetDraftByID to GetPostByID
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PostResponse{ // Changed from DraftResponse to PostResponse
		ID:           post.ID,
		UserID:       post.UserID,
		SuggestionID: post.SuggestionID,
		Text:         post.Text, // Changed from DraftText to Text
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	})
}

type ListPostsByUserIDRequest struct { // Changed from ListDraftsByUserIDRequest to ListPostsByUserIDRequest
	UserID   string `form:"user_id" binding:"required,uuid"`
	Page     int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) ListPostsByUserID(ctx *gin.Context) { // Changed from ListDraftsByUserID to ListPostsByUserID
	var req ListPostsByUserIDRequest
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

	posts, err := server.store.ListPostsByUserID(ctx, db.ListPostsByUserIDParams{ // Changed from ListDraftsByUserID to ListPostsByUserID
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

type UpdatePostRequest struct { // Changed from UpdateDraftRequest to UpdatePostRequest
	ID   string `json:"id" binding:"required,uuid"`
	Text string `json:"text" binding:"required"` // Changed from DraftText to Text
}

func (server *Server) UpdatePost(ctx *gin.Context) { // Changed from UpdateDraft to UpdatePost
	var req UpdatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePostParams{ // Changed from UpdateDraftParams to UpdatePostParams
		ID:   pgtype.UUID{Bytes: postID, Valid: true},
		Text: req.Text, // Changed from DraftText to Text
	}

	post, err := server.store.UpdatePost(ctx, arg) // Changed from UpdateDraft to UpdatePost
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PostResponse{ // Changed from DraftResponse to PostResponse
		ID:           post.ID,
		UserID:       post.UserID,
		SuggestionID: post.SuggestionID,
		Text:         post.Text, // Changed from DraftText to Text
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	})
}

type DeletePostRequest struct { // Changed from DeleteDraftRequest to DeletePostRequest
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeletePost(ctx *gin.Context) { // Changed from DeleteDraft to DeletePost
	var req DeletePostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeletePost(ctx, pgtype.UUID{Bytes: postID, Valid: true}) // Changed from DeleteDraft to DeletePost
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
