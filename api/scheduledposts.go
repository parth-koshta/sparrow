package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type createScheduledPostRequest struct {
	UserID        string           `json:"user_id" binding:"required,uuid"`
	DraftID       string           `json:"draft_id" binding:"required,uuid"`
	ScheduledTime pgtype.Timestamp `json:"scheduled_time" binding:"required"`
	Status        string           `json:"status" binding:"required"`
}

func (server *Server) createScheduledPost(ctx *gin.Context) {
	var req createScheduledPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	draftID, err := uuid.Parse(req.DraftID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateScheduledPostParams{
		UserID:        pgtype.UUID{Bytes: userID, Valid: true},
		DraftID:       pgtype.UUID{Bytes: draftID, Valid: true},
		ScheduledTime: req.ScheduledTime,
		Status:        req.Status,
	}

	scheduledPost, err := server.store.CreateScheduledPost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scheduledPost)
}

type getScheduledPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getScheduledPost(ctx *gin.Context) {
	var req getScheduledPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	scheduledPost, err := server.store.GetScheduledPostByID(ctx, pgtype.UUID{Bytes: postID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scheduledPost)
}

type listScheduledPostsByUserIDRequest struct {
	UserID   string `form:"user_id" binding:"required,uuid"`
	Page     int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) listScheduledPostsByUserID(ctx *gin.Context) {
	var req listScheduledPostsByUserIDRequest
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

	scheduledPosts, err := server.store.ListScheduledPostsByUserID(ctx, db.ListScheduledPostsByUserIDParams{
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

type updateScheduledPostRequest struct {
	ID            string           `json:"id" binding:"required,uuid"`
	ScheduledTime pgtype.Timestamp `json:"scheduled_time"`
	Status        string           `json:"status"`
}

func (server *Server) updateScheduledPost(ctx *gin.Context) {
	var req updateScheduledPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateScheduledPostParams{
		ID:            pgtype.UUID{Bytes: postID, Valid: true},
		ScheduledTime: req.ScheduledTime,
		Status:        req.Status,
	}

	scheduledPost, err := server.store.UpdateScheduledPost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, scheduledPost)
}

type deleteScheduledPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) deleteScheduledPost(ctx *gin.Context) {
	var req deleteScheduledPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	postID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeleteScheduledPost(ctx, pgtype.UUID{Bytes: postID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
