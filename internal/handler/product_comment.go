package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	"github.com/go-playground/validator/v10"
)

type productCommentHandlerImpl struct {
	service   domain.Service
	validator *validator.Validate
}

func NewProductCommentHandler(service domain.Service, validator *validator.Validate) domain.ProductCommentHandler {
	return &productCommentHandlerImpl{
		service:   service,
		validator: validator,
	}
}

// CreateProductCommentHandler godoc
//
//	@Summary		create product comment endpoint
//	@Description	create product comment
//	@Accept			json
//	@Produce		json
//	@Tags			Product Comment
//	@Param			id		path	string								true	"product id"
//	@Param			request	body	entity.ProductCommentCreateRequest	true	"product comment data for create"
//	@Security		Bearer
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/product/comment/{id} [post]
func (h *productCommentHandlerImpl) CreateProductCommentHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	currentUserID := r.Context().Value(helper.CtxUserID).(string)
	var reqBody entity.ProductCommentCreateRequest
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
	if err := h.service.ProductComment().CreateProductComment(r.Context(), productID, currentUserID, &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateProductCommentHandler godoc
//
//	@Summary		update product comment endpoint
//	@Description	update product comment
//	@Accept			json
//	@Produce		json
//	@Tags			Product Comment
//	@Param			id		path	string								true	"product comment id"
//	@Param			request	body	entity.ProductCommentUpdateRequest	true	"product comment data for update"
//	@Security		Bearer
//	@Success		200
//	@Failure		400
//	@Failure		403
//	@Failure		500
//	@Router			/product/comment/{id} [put]
func (h *productCommentHandlerImpl) UpdateProductCommentHandler(w http.ResponseWriter, r *http.Request) {
	productCommentID := r.PathValue("id")
	currentUserID := r.Context().Value(helper.CtxUserID).(string)
	var reqBody entity.ProductCommentUpdateRequest
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
	productComment, err := h.service.ProductComment().GetProductCommentByID(r.Context(), productCommentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if productComment.UserID != currentUserID {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if err := h.service.ProductComment().UpdateProductComment(r.Context(), productCommentID, &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProductCommentHandler godoc
//
//	@Summary		delete product comment endpoint
//	@Description	delete product comment
//	@Accept			json
//	@Produce		json
//	@Tags			Product Comment
//	@Param			id	path	string	true	"product comment id"
//	@Security		Bearer
//	@Success		204
//	@Failure		400
//	@Failure		403
//	@Failure		500
//	@Router			/product/comment/{id} [delete]
func (h *productCommentHandlerImpl) DeleteProductCommentHandler(w http.ResponseWriter, r *http.Request) {
	productCommentID := r.PathValue("id")
	currentUserID := r.Context().Value(helper.CtxUserID).(string)
	isAdmin := r.Context().Value(helper.CtxIsAdmin).(bool)
	productComment, err := h.service.ProductComment().GetProductCommentByID(r.Context(), productCommentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if productComment.UserID != currentUserID && !isAdmin {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if err := h.service.ProductComment().DeleteProductComment(r.Context(), productCommentID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
