package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/parth-koshta/sparrow/db/sqlc"
)

type createDraftRequest struct {
	UserID       string `json:"user_id" binding:"required,uuid"`
	SuggestionID string `json:"suggestion_id" binding:"required,uuid"`
	DraftText    string `json:"draft_text" binding:"required"`
}

type draftResponse struct {
	ID           pgtype.UUID
	UserID       pgtype.UUID
	SuggestionID pgtype.UUID
	DraftText    string
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

func (server *Server) createDraft(ctx *gin.Context) {
	var req createDraftRequest
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

	arg := db.CreateDraftParams{
		UserID:       pgtype.UUID{Bytes: userID, Valid: true},
		SuggestionID: pgtype.UUID{Bytes: suggestionID, Valid: true},
		DraftText:    req.DraftText,
	}

	draft, err := server.store.CreateDraft(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, draftResponse{
		ID:           draft.ID,
		UserID:       draft.UserID,
		SuggestionID: draft.SuggestionID,
		DraftText:    draft.DraftText,
		CreatedAt:    draft.CreatedAt,
		UpdatedAt:    draft.UpdatedAt,
	})
}

type getDraftRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getDraft(ctx *gin.Context) {
	var req getDraftRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	draftID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	draft, err := server.store.GetDraftByID(ctx, pgtype.UUID{Bytes: draftID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, draftResponse{
		ID:           draft.ID,
		UserID:       draft.UserID,
		SuggestionID: draft.SuggestionID,
		DraftText:    draft.DraftText,
		CreatedAt:    draft.CreatedAt,
		UpdatedAt:    draft.UpdatedAt,
	})
}

type listDraftsByUserIDRequest struct {
	UserID   string `form:"user_id" binding:"required,uuid"`
	Page     int32  `form:"page" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) listDraftsByUserID(ctx *gin.Context) {
	var req listDraftsByUserIDRequest
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

	drafts, err := server.store.ListDraftsByUserID(ctx, db.ListDraftsByUserIDParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, drafts)
}

type updateDraftRequest struct {
	ID        string `json:"id" binding:"required,uuid"`
	DraftText string `json:"draft_text" binding:"required"`
}

func (server *Server) updateDraft(ctx *gin.Context) {
	var req updateDraftRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	draftID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateDraftParams{
		ID:        pgtype.UUID{Bytes: draftID, Valid: true},
		DraftText: req.DraftText,
	}

	draft, err := server.store.UpdateDraft(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, draftResponse{
		ID:           draft.ID,
		UserID:       draft.UserID,
		SuggestionID: draft.SuggestionID,
		DraftText:    draft.DraftText,
		CreatedAt:    draft.CreatedAt,
		UpdatedAt:    draft.UpdatedAt,
	})
}

type deleteDraftRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) deleteDraft(ctx *gin.Context) {
	var req deleteDraftRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	draftID, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeleteDraft(ctx, pgtype.UUID{Bytes: draftID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
