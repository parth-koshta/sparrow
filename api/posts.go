package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/parth-koshta/sparrow/client"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	dbtypes "github.com/parth-koshta/sparrow/db/types"
	"github.com/rs/zerolog/log"
)

type CreatePostRequest struct {
	SuggestionID string `json:"suggestion_id" binding:"required,uuid"`
	Text         string `json:"text" binding:"required"`
}

type PostResponse struct {
	ID           pgtype.UUID
	SuggestionID pgtype.UUID
	Text         string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

func (server *Server) CreatePost(ctx *gin.Context) {
	var req CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	suggestionID, err := uuid.Parse(req.SuggestionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		UserID:       pgtype.UUID{Bytes: userID, Valid: true},
		SuggestionID: pgtype.UUID{Bytes: suggestionID, Valid: true},
		Text:         req.Text,
	}

	post, err := server.store.CreatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PostResponse{
		ID:           post.ID,
		SuggestionID: post.SuggestionID,
		Text:         post.Text,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	})
}

type GetPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) GetPost(ctx *gin.Context) {
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

	post, err := server.store.GetPostByID(ctx, pgtype.UUID{Bytes: postID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PostResponse{
		ID:           post.ID,
		SuggestionID: post.SuggestionID,
		Text:         post.Text,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	})
}

type ListPostsByUserIDRequest struct {
	Page     int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) ListPostsByUserID(ctx *gin.Context) {
	var req ListPostsByUserIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize

	posts, err := server.store.ListPostsByUserID(ctx, db.ListPostsByUserIDParams{
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

type UpdatePostRequest struct {
	ID   string `json:"id" binding:"required,uuid"`
	Text string `json:"text" binding:"required"`
}

func (server *Server) UpdatePost(ctx *gin.Context) {
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

	arg := db.UpdatePostParams{
		ID:   pgtype.UUID{Bytes: postID, Valid: true},
		Text: req.Text,
	}

	post, err := server.store.UpdatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PostResponse{
		ID:           post.ID,
		SuggestionID: post.SuggestionID,
		Text:         post.Text,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	})
}

type DeletePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeletePost(ctx *gin.Context) {
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

	_, err = server.store.DeletePost(ctx, pgtype.UUID{Bytes: postID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

type PublishOnLinkedInRequest struct {
	PostID          string `json:"post_id" binding:"required,uuid"`
	SocialAccountID string `json:"social_account_id" binding:"required,uuid"`
}

type PublishOnLinkedInResponse struct {
	PostID          string `json:"post_id"`
	SocialAccountID string `json:"social_account_id"`
	Success         bool   `json:"success"`
}

func (server *Server) PublishOnLinkedIn(ctx *gin.Context) {
	log.Info().Msg("PublishOnLinkedIn called")
	var req PublishOnLinkedInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.PostID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	socialAccountID, err := uuid.Parse(req.SocialAccountID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Info().Msg("PublishOnLinkedIn called userID: " + userID.String())

	socialAccount, err := server.store.GetSocialAccountByID(ctx, pgtype.UUID{Bytes: socialAccountID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if socialAccount.UserID.Bytes != userID {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	post, err := server.store.GetPostByID(ctx, pgtype.UUID{Bytes: postID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.linkedinClient.PublishPost(socialAccount.AccessToken, socialAccount.LinkedinSub.String, client.PayloadPublishPost{
		PostID: post.ID,
		Text:   post.Text,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	updatePostStatusArgs := db.UpdatePostStatusParams{
		ID:     post.ID,
		Status: string(dbtypes.PostStatusPublished),
	}
	_, err = server.store.UpdatePostStatus(ctx, updatePostStatusArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, PublishOnLinkedInResponse{
		PostID:          req.PostID,
		SocialAccountID: req.SocialAccountID,
		Success:         true,
	})
}
