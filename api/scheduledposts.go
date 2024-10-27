package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type CreateScheduledPostRequest struct {
	UserID        string           `json:"user_id" binding:"required,uuid"`
	PostID        string           `json:"post_id" binding:"required,uuid"`
	ScheduledTime pgtype.Timestamp `json:"scheduled_time" binding:"required"`
	Status        string           `json:"status" binding:"required"`
}

func (server *Server) CreateScheduledPost(ctx *gin.Context) {
	var req CreateScheduledPostRequest
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

type GetScheduledPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) GetScheduledPost(ctx *gin.Context) {
	var req GetScheduledPostRequest
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

type ListScheduledPostsByUserIDRequest struct {
	UserID   string `form:"user_id" binding:"required,uuid"`
	Page     int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) ListScheduledPostsByUserID(ctx *gin.Context) {
	var req ListScheduledPostsByUserIDRequest
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

	scheduledPosts, err := server.store.ListPostSchedulesByUserID(ctx, db.ListPostSchedulesByUserIDParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scheduledPosts)
}

type UpdateScheduledPostRequest struct {
	ID            string           `json:"id" binding:"required,uuid"`
	ScheduledTime pgtype.Timestamp `json:"scheduled_time"`
	Status        string           `json:"status"`
}

func (server *Server) UpdateScheduledPost(ctx *gin.Context) {
	var req UpdateScheduledPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePostScheduleParams{
		ID:            pgtype.UUID{Bytes: postID, Valid: true},
		ScheduledTime: req.ScheduledTime,
		Status:        req.Status,
	}

	scheduledPost, err := server.store.UpdatePostSchedule(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scheduledPost)
}

type DeleteScheduledPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeleteScheduledPost(ctx *gin.Context) {
	var req DeleteScheduledPostRequest
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
