package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	dbtypes "github.com/parth-koshta/sparrow/db/types"
	"github.com/parth-koshta/sparrow/token"
)

type SocialAccountResponse struct {
	Platform       string
	AccountName    string
	TokenExpiresAt pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

type AddLinkedInAccountRequest struct {
	Code string `json:"code" binding:"required"`
}

func (server *Server) AddLinkedInAccount(ctx *gin.Context) {
	var req AddLinkedInAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AUTHORIZATION_PAYLOAD_KEY).(*token.Payload)
	if !authPayload.ID.Valid {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("invalid user ID in auth payload")))
		return
	}

	userIDBytes := authPayload.ID.Bytes[:]
	userID, err := uuid.FromBytes(userIDBytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("failed to parse user ID: %v", err)))
		return
	}

	accessTokenResp, err := server.linkedInClient.GetAccessToken(req.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("failed to get access token: %v", err)))
		return
	}

	accessToken := accessTokenResp.AccessToken
	idToken := accessTokenResp.IDToken
	expiresIn := accessTokenResp.ExpiresIn

	tokenExpiresAt := time.Now().Add(time.Second * time.Duration(expiresIn))
	tokenExpiresAtPg := pgtype.Timestamp{
		Time:  tokenExpiresAt,
		Valid: true,
	}

	userInfo, err := server.linkedInClient.GetUserInfo(accessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("failed to get user info: %v", err)))
		return
	}

	arg := db.CreateSocialAccountParams{
		UserID:         pgtype.UUID{Bytes: userID, Valid: true},
		Platform:       dbtypes.SocialPlatformLinkedIn.String(),
		AccountName:    userInfo.Name,
		AccountEmail:   userInfo.Email,
		AccessToken:    accessToken,
		IDToken:        idToken,
		TokenExpiresAt: tokenExpiresAtPg,
	}

	socialAccount, err := server.store.CreateSocialAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SocialAccountResponse{
		Platform:       socialAccount.Platform,
		AccountName:    socialAccount.AccountName,
		TokenExpiresAt: socialAccount.TokenExpiresAt,
		UpdatedAt:      socialAccount.UpdatedAt,
	})
}

type UpdateAccessTokenRequest struct {
	Code string `json:"code" binding:"required"`
}

func (server *Server) UpdateLinkedInAccessToken(ctx *gin.Context) {
	var req UpdateAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get user ID from the auth middleware
	authPayload := ctx.MustGet(AUTHORIZATION_PAYLOAD_KEY).(*token.Payload)
	if !authPayload.ID.Valid {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("invalid user ID in auth payload")))
		return
	}

	userIDBytes := authPayload.ID.Bytes[:]
	userID, err := uuid.FromBytes(userIDBytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("failed to parse user ID: %v", err)))
		return
	}

	// Call LinkedIn client to get new access token
	// linkedinClient := NewLinkedInClient() // Assuming you have a function to create a LinkedIn client
	tokenInfo, err := server.linkedInClient.GetAccessToken(req.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Calculate the token expiration time
	tokenExpiresAt := time.Now().Add(time.Duration(tokenInfo.ExpiresIn) * time.Second)

	// Update the token in the database
	arg := db.UpdateSocialAccountTokenParams{
		UserID:         pgtype.UUID{Bytes: userID, Valid: true},
		AccessToken:    tokenInfo.AccessToken,
		IDToken:        tokenInfo.IDToken,
		TokenExpiresAt: pgtype.Timestamp{Time: tokenExpiresAt, Valid: true},
	}

	socialAccount, err := server.store.UpdateSocialAccountToken(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SocialAccountResponse{
		Platform:       socialAccount.Platform,
		AccountName:    socialAccount.AccountName,
		TokenExpiresAt: socialAccount.TokenExpiresAt,
		UpdatedAt:      socialAccount.UpdatedAt,
	})
}

type GetSocialAccountRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) GetSocialAccount(ctx *gin.Context) {
	var req GetSocialAccountRequest
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

	ctx.JSON(http.StatusOK, SocialAccountResponse{
		Platform:       socialAccount.Platform,
		AccountName:    socialAccount.AccountName,
		TokenExpiresAt: socialAccount.TokenExpiresAt,
		UpdatedAt:      socialAccount.UpdatedAt,
	})
}

type ListSocialAccountsByUserIDRequest struct {
	UserID   string `form:"user_id" binding:"required,uuid"`
	Page     int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) ListSocialAccountsByUserID(ctx *gin.Context) {
	var req ListSocialAccountsByUserIDRequest
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

	var responses []SocialAccountResponse
	for _, acc := range socialAccounts {
		responses = append(responses, SocialAccountResponse{
			Platform:       acc.Platform,
			AccountName:    acc.AccountName,
			TokenExpiresAt: acc.TokenExpiresAt,
			UpdatedAt:      acc.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, responses)
}

type DeleteSocialAccountRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeleteSocialAccount(ctx *gin.Context) {
	var req DeleteSocialAccountRequest
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
