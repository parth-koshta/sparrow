package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type createSocialAccountRequest struct {
	UserID      string `json:"user_id" binding:"required,uuid"`
	Platform    string `json:"platform" binding:"required"`
	AccountName string `json:"account_name" binding:"required"`
	AccessToken string `json:"access_token" binding:"required"`
}

// socialAccountResponse defines the response structure for a social account
type socialAccountResponse struct {
	ID          pgtype.UUID
	UserID      pgtype.UUID
	Platform    string
	AccountName string
	AccessToken string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

// CreateSocialAccount handles the creation of a social account
func (server *Server) createSocialAccount(ctx *gin.Context) {
	var req createSocialAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSocialAccountParams{
		UserID:      pgtype.UUID{Bytes: userID, Valid: true},
		Platform:    req.Platform,
		AccountName: req.AccountName,
		AccessToken: req.AccessToken,
	}

	socialAccount, err := server.store.CreateSocialAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, socialAccountResponse{
		ID:          socialAccount.ID,
		UserID:      socialAccount.UserID,
		Platform:    socialAccount.Platform,
		AccountName: socialAccount.AccountName,
		AccessToken: socialAccount.AccessToken,
		CreatedAt:   socialAccount.CreatedAt,
		UpdatedAt:   socialAccount.UpdatedAt,
	})
}

type getSocialAccountRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// GetSocialAccount handles retrieval of a social account by ID
func (server *Server) getSocialAccount(ctx *gin.Context) {
	var req getSocialAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accountID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	socialAccount, err := server.store.GetSocialAccountByID(ctx, pgtype.UUID{Bytes: accountID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, socialAccountResponse{
		ID:          socialAccount.ID,
		UserID:      socialAccount.UserID,
		Platform:    socialAccount.Platform,
		AccountName: socialAccount.AccountName,
		AccessToken: socialAccount.AccessToken,
		CreatedAt:   socialAccount.CreatedAt,
		UpdatedAt:   socialAccount.UpdatedAt,
	})
}

type listSocialAccountsByUserIDRequest struct {
	UserID   string `form:"user_id" binding:"required,uuid"`
	Page     int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
}

// ListSocialAccountsByUserID handles listing of social accounts for a specific user
func (server *Server) listSocialAccountsByUserID(ctx *gin.Context) {
	var req listSocialAccountsByUserIDRequest
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

	socialAccounts, err := server.store.ListSocialAccountsByUserID(ctx, db.ListSocialAccountsByUserIDParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var responses []socialAccountResponse
	for _, acc := range socialAccounts {
		responses = append(responses, socialAccountResponse{
			ID:          acc.ID,
			UserID:      acc.UserID,
			Platform:    acc.Platform,
			AccountName: acc.AccountName,
			AccessToken: acc.AccessToken,
			CreatedAt:   acc.CreatedAt,
			UpdatedAt:   acc.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, responses)
}

// updateSocialAccountRequest defines the request structure for updating a social account
type updateSocialAccountRequest struct {
	Platform    string `json:"platform"`
	AccountName string `json:"account_name"`
	AccessToken string `json:"access_token"`
}

// UpdateSocialAccount handles updating an existing social account
func (server *Server) updateSocialAccount(ctx *gin.Context) {
	var req updateSocialAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accountID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSocialAccountParams{
		ID:          pgtype.UUID{Bytes: accountID, Valid: true},
		Platform:    req.Platform,
		AccountName: req.AccountName,
		AccessToken: req.AccessToken,
	}

	socialAccount, err := server.store.UpdateSocialAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, socialAccountResponse{
		ID:          socialAccount.ID,
		UserID:      socialAccount.UserID,
		Platform:    socialAccount.Platform,
		AccountName: socialAccount.AccountName,
		AccessToken: socialAccount.AccessToken,
		CreatedAt:   socialAccount.CreatedAt,
		UpdatedAt:   socialAccount.UpdatedAt,
	})
}

// deleteSocialAccountRequest defines the request structure for deleting a social account
type deleteSocialAccountRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// DeleteSocialAccount handles the deletion of a social account
func (server *Server) deleteSocialAccount(ctx *gin.Context) {
	var req deleteSocialAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accountID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeleteSocialAccount(ctx, pgtype.UUID{Bytes: accountID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
