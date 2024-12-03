package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
	"github.com/parth-koshta/sparrow/util"
	"github.com/parth-koshta/sparrow/worker"
)

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,min=6"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID              pgtype.UUID      `json:"id"`
	Username        pgtype.Text      `json:"username"`
	Email           string           `json:"email"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	IsEmailVerified pgtype.Bool      `json:"is_email_verified"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Email: req.Email,
			PasswordHash: pgtype.Text{
				String: hashedPassword,
				Valid:  true,
			},
		},
		AfterCreate: func(user db.User) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				Email: req.Email,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}
			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			ctx.JSON(http.StatusConflict, customErrorResponse(err, "Email already exists. Please use a different email."))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userResponse := &UserResponse{
		ID:              txResult.User.ID,
		Username:        txResult.User.Username,
		Email:           txResult.User.Email,
		CreatedAt:       txResult.User.CreatedAt,
		UpdatedAt:       txResult.User.UpdatedAt,
		IsEmailVerified: pgtype.Bool{Bool: false, Valid: true},
	}
	ctx.JSON(http.StatusOK, userResponse)
}

type GetUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) GetUser(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	requestUUID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	parsedUUID, err := GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if parsedUUID != requestUUID {
		err := fmt.Errorf("cannot access details of another user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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

type ListUsersRequest struct {
	Page     int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) ListUsers(ctx *gin.Context) {
	var req ListUsersRequest
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

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,min=6"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, customErrorResponse(err, "Please re-check your email and password"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(user.PasswordHash.String, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, customErrorResponse(err, "Please re-check your email and password"))
		return
	}

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

	ctx.JSON(http.StatusOK, &LoginUserResponse{
		Token: accessToken,
		User: UserResponse{
			ID:              user.ID,
			Username:        user.Username,
			Email:           user.Email,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
			IsEmailVerified: pgtype.Bool{Bool: user.IsEmailVerified, Valid: true},
		},
	})
}

type VerifyUserEmailRequest struct {
	Email      string `form:"email" binding:"required,min=6"`
	SecretCode string `form:"secret_code" binding:"required"`
}

type VerifyUserEmailResponse struct {
	IsVerified bool `json:"is_verified"`
}

func (server *Server) VerifyUserEmail(ctx *gin.Context) {
	var req VerifyUserEmailRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	txResult, err := server.store.VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailId:    req.Email,
		SecretCode: req.SecretCode,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := &VerifyUserEmailResponse{
		IsVerified: txResult.User.IsEmailVerified,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type ResendVerifyEmailRequest struct {
	Email string `json:"email" binding:"required,min=6"`
}
type ResendVerifyEmailResponse struct {
	Message string `json:"message"`
}

func (server *Server) ResendVerifyEmail(ctx *gin.Context) {
	var req ResendVerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, customErrorResponse(err, "User does not exist"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.IsEmailVerified {
		ctx.JSON(http.StatusBadRequest, customErrorResponse(nil, "Email is already verified"))
		return
	}

	verifyEmail, err := server.store.GetVerifyEmail(ctx, user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, customErrorResponse(err, "Verify email record not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if time.Since(verifyEmail.CreatedAt.Time) < 30*time.Second {
		ctx.JSON(http.StatusBadRequest, customErrorResponse(err, "Please wait for 30 seconds before resending the verification email"))
		return
	}

	_, err = server.store.InvalidateVerifyEmail(ctx, verifyEmail.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	taskPayload := &worker.PayloadSendVerifyEmail{
		Email: req.Email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &ResendVerifyEmailResponse{
		Message: "Verification email sent successfully",
	})
}

// api to check if email is verified
type IsEmailVerifiedRequest struct {
	Email string `json:"email" binding:"required,min=6"`
}
type IsEmailVerifiedResponse struct {
	IsVerified bool `json:"is_email_verified"`
}

func (server *Server) IsEmailVerified(ctx *gin.Context) {
	var req IsEmailVerifiedRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, customErrorResponse(err, "User does not exist"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &IsEmailVerifiedResponse{
		IsVerified: user.IsEmailVerified,
	})
}
