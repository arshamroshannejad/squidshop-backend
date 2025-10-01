package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	"github.com/go-playground/validator/v10"
)

type productCommentLikeHandler struct {
	service   domain.Service
	validator *validator.Validate
}

func NewProductCommentLikeHandler(service domain.Service, validator *validator.Validate) domain.ProductCommentLikeHandler {
	return &productCommentLikeHandler{
		service:   service,
		validator: validator,
	}
}

// CreateProductCommentLikeHandler godoc
//
//	@Summary		create product comment like endpoint
//	@Description	create product comment like
//	@Accept			json
//	@Produce		json
//	@Tags			Product Comment Like
//	@Param			id		path	string								true	"comment id"
//	@Param			request	body	entity.PostCommentLikeCreateUpdate	true	"product comment like data for create"
//	@Security		Bearer
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/product/comment/like/{id} [post]
func (h *productCommentLikeHandler) CreateProductCommentLikeHandler(w http.ResponseWriter, r *http.Request) {
	commentID := r.PathValue("id")
	currentUser := r.Context().Value(helper.CtxUserID).(string)
	var reqBody entity.PostCommentLikeCreateUpdate
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	if err := h.validator.Struct(reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	if err := h.service.ProductCommentLike().CreateProductCommentLike(r.Context(), commentID, currentUser, &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateProductCommentLikeHandler godoc
//
//	@Summary		update product comment like endpoint
//	@Description	update product comment like
//	@Accept			json
//	@Produce		json
//	@Tags			Product Comment Like
//	@Param			id		path	string								true	"comment id"
//	@Param			request	body	entity.PostCommentLikeCreateUpdate	true	"product comment like data for update"
//	@Security		Bearer
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/product/comment/like/{id} [put]
func (h *productCommentLikeHandler) UpdateProductCommentLikeHandler(w http.ResponseWriter, r *http.Request) {
	commentID := r.PathValue("id")
	currentUser := r.Context().Value(helper.CtxUserID).(string)
	var reqBody entity.PostCommentLikeCreateUpdate
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	if err := h.validator.Struct(reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	if err := h.service.ProductCommentLike().UpdateProductCommentLike(r.Context(), commentID, currentUser, &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProductCommentLikeHandler godoc
//
//	@Summary		delete product comment like endpoint
//	@Description	delete product comment like
//	@Accept			json
//	@Produce		json
//	@Tags			Product Comment Like
//	@Param			id	path	string	true	"comment id"
//	@Security		Bearer
//	@Success		204
//	@Failure		500
//	@Router			/product/comment/like/{id} [delete]
func (h *productCommentLikeHandler) DeleteProductCommentLikeHandler(w http.ResponseWriter, r *http.Request) {
	commentID := r.PathValue("id")
	currentUser := r.Context().Value(helper.CtxUserID).(string)
	if err := h.service.ProductCommentLike().DeleteProductCommentLike(r.Context(), commentID, currentUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
