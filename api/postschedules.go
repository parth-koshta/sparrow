package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type CreatePostScheduleRequest struct {
	UserID        string           `json:"user_id" binding:"required,uuid"`
	PostID        string           `json:"post_id" binding:"required,uuid"`
	ScheduledTime pgtype.Timestamp `json:"scheduled_time" binding:"required"`
	Status        string           `json:"status" binding:"required"`
}

func (server *Server) CreatePostSchedule(ctx *gin.Context) {
	var req CreatePostScheduleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	PostID, err := uuid.Parse(req.PostID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostScheduleParams{
		UserID:        pgtype.UUID{Bytes: userID, Valid: true},
		PostID:        pgtype.UUID{Bytes: PostID, Valid: true},
		ScheduledTime: req.ScheduledTime,
		Status:        req.Status,
	}

	scheduledPost, err := server.store.CreatePostSchedule(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scheduledPost)
}

type GetPostScheduleRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) GetPostSchedule(ctx *gin.Context) {
	var req GetPostScheduleRequest
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

type DeletePostScheduleRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeletePostSchedule(ctx *gin.Context) {
	var req DeletePostScheduleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeletePostSchedule(ctx, pgtype.UUID{Bytes: postID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
