package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	"github.com/go-playground/validator/v10"
)

type productRatingHandlerImpl struct {
	service   domain.Service
	validator *validator.Validate
}

func NewProductRatingHandler(service domain.Service, validator *validator.Validate) domain.ProductRatingHandler {
	return &productRatingHandlerImpl{
		service:   service,
		validator: validator,
	}
}

// CreateOrUpdateProductRatingHandler godoc
//
//	@Summary		create or update product rating endpoint
//	@Description	create or update product rating
//	@Accept			json
//	@Produce		json
//	@Tags			Product Rating
//	@Param			id		path	string						true	"product id"
//	@Param			request	body	entity.ProductRatingRequest	true	"product rating data for create or update"
//	@Security		Bearer
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/product/rating/{id} [post]
func (h *productRatingHandlerImpl) CreateOrUpdateProductRatingHandler(w http.ResponseWriter, r *http.Request) {
	currentUserID := r.Context().Value(helper.CtxUserID).(string)
	productID := r.PathValue("id")
	var reqBody entity.ProductRatingRequest
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
	if err := h.service.ProductRating().CreateOrUpdateProductRating(r.Context(), productID, currentUserID, &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProductRatingHandler godoc
//
//	@Summary		delete product rating endpoint
//	@Description	delete product rating
//	@Accept			json
//	@Produce		json
//	@Tags			Product Rating
//	@Param			id	path	string	true	"product id"
//	@Security		Bearer
//	@Success		204
//	@Failure		400
//	@Failure		500
//	@Router			/product/rating/{id} [delete]
func (h *productRatingHandlerImpl) DeleteProductRatingHandler(w http.ResponseWriter, r *http.Request) {
	currentUserID := r.Context().Value(helper.CtxUserID).(string)
	productID := r.PathValue("id")
	if err := h.service.ProductRating().DeleteProductRating(r.Context(), productID, currentUserID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
