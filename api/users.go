package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	token "github.com/parth-koshta/sparrow/token"
	"github.com/parth-koshta/sparrow/util"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,min=6"`
	Password string `json:"password" binding:"required"`
}

type userResponse struct {
	ID        pgtype.UUID
	Username  pgtype.Text
	Email     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email: req.Email,
		PasswordHash: pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		},
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

type getUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	requestUUID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authUUID, err := authPayload.ID.UUIDValue()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authUUIDConverted, err := uuid.FromBytes(authUUID.Bytes[:])
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if authUUIDConverted != requestUUID {
		err := fmt.Errorf("cannot access details of another user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	parsedUUID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userID := pgtype.UUID{
		Valid: true,
		Bytes: parsedUUID,
	}
	user, err := server.store.GetUserByID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUsersRequest struct {
	Page     int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize

	users, err := server.store.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,min=6"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	Token string       `json:"token"`
	User  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get the user from the database
	user, err := server.store.GetUserByEmail(ctx, req.Email)
	fmt.Println("USERRRR 1", user)
	if err != nil {
		fmt.Println("USERRRR  error", err)
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Check if the password is correct
	err = util.CheckPassword(user.PasswordHash.String, req.Password)
	if err != nil {
		fmt.Println("Chechpassword failed", err)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// Generate a new access token
	userUUID, err := user.ID.UUIDValue()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	accessToken, err := server.tokenMaker.CreateToken(userUUID, user.Email, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return the access token
	ctx.JSON(http.StatusOK, &loginUserResponse{
		Token: accessToken,
		User: userResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}
