package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type SchedulePostRequest struct {
	PostID          string           `json:"post_id" binding:"required,uuid"`
	ScheduledTime   pgtype.Timestamp `json:"scheduled_time" binding:"required"`
	SocialAccountID string           `json:"social_account_id" binding:"required,uuid"`
}

func (server *Server) SchedulePost(ctx *gin.Context) {
	var req SchedulePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	scheduledTime, err := req.ScheduledTime.Value()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid scheduled_time format: %w", err)))
		return
	}

	if t, ok := scheduledTime.(time.Time); ok {
		scheduledTime = t.UTC()
	}

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	socialAccountID, err := uuid.Parse(req.SocialAccountID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	socialAccount, err := server.store.GetSocialAccountByID(ctx, pgtype.UUID{Bytes: socialAccountID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if socialAccount.UserID.Bytes != userID {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.PostID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	createPostScheduleTxArg := db.SchedulePostTxParams{
		UserID:          pgtype.UUID{Bytes: userID, Valid: true},
		PostID:          pgtype.UUID{Bytes: postID, Valid: true},
		ScheduledTime:   pgtype.Timestamp{Time: scheduledTime.(time.Time), Valid: true},
		SocialAccountID: pgtype.UUID{Bytes: socialAccountID, Valid: true},
	}

	scheduledPost, err := server.store.SchedulePostTx(ctx, createPostScheduleTxArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scheduledPost)
}

type GetScheduleRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) GetSchedule(ctx *gin.Context) {
	var req GetScheduleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	scheduledPost, err := server.store.GetPostScheduleByID(ctx, pgtype.UUID{Bytes: postID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scheduledPost)
}

type DeleteScheduleRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeleteSchedule(ctx *gin.Context) {
	var req DeleteScheduleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeleteSchedule(ctx, pgtype.UUID{Bytes: postID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

type ListSchedulesByUserIDRequest struct {
	Limit  int32 `form:"limit" binding:"required"`
	Offset int32 `form:"offset" binding:"required"`
}

func (server *Server) ListPostSchedulesForUser(ctx *gin.Context) {
	var req ListSchedulesByUserIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPostSchedulesByUserIDParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	posts, err := server.store.ListPostSchedulesByUserID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
