package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	_ "github.com/arshamroshannejad/squidshop-backend/internal/model"
	"github.com/go-playground/validator/v10"
)

type productHandlerImpl struct {
	service   domain.Service
	validator *validator.Validate
}

func NewProductHandler(service domain.Service, validator *validator.Validate) domain.ProductHandler {
	return &productHandlerImpl{
		service:   service,
		validator: validator,
	}
}

// GetAllProductsHandler godoc
//
//	@Summary		get all products endpoint
//	@Description	get all products in db
//	@Accept			json
//	@Produce		json
//	@Tags			Product
//	@Success		200	{array}	model.Product
//	@Failure		500
//	@Router			/product [get]
func (h *productHandlerImpl) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.Product().GetAllProducts(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// GetProductByIDHandler godoc
//
//	@Summary		get product by id endpoint
//	@Description	get product by id
//	@Accept			json
//	@Produce		json
//	@Tags			Product
//	@Param			id	path		string	true	"product id"
//	@Success		200	{object}	model.Product
//	@Failure		500
//	@Router			/product/id/{id} [get]
func (h *productHandlerImpl) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	product, err := h.service.Product().GetProductByID(r.Context(), productID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// GetProductBySlugHandler godoc
//
//	@Summary		get product by slug endpoint
//	@Description	get product by slug
//	@Accept			json
//	@Produce		json
//	@Tags			Product
//	@Param			slug	path		string	true	"product slug"
//	@Success		200		{object}	model.Product
//	@Failure		500
//	@Router			/product/slug/{slug} [get]
func (h *productHandlerImpl) GetProductBySlugHandler(w http.ResponseWriter, r *http.Request) {
	productSlug := r.PathValue("slug")
	product, err := h.service.Product().GetProductBySlug(r.Context(), productSlug)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// CreateProductHandler godoc
//
//	@Summary		create product endpoint
//	@Description	create product
//	@Accept			json
//	@Produce		json
//	@Tags			Product
//	@Param			request	body	entity.ProductCreateRequest	true	"product data for create"
//	@Security		Bearer
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/product [post]
func (h *productHandlerImpl) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody entity.ProductCreateRequest
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
	if err := h.service.Product().CreateProduct(r.Context(), &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateProductHandler godoc
//
//	@Summary		update product endpoint
//	@Description	update product by id
//	@Accept			json
//	@Produce		json
//	@Tags			Product
//	@Param			id		path	string						true	"product id"
//	@Param			request	body	entity.ProductUpdateRequest	true	"product data for update"
//	@Security		Bearer
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/product/{id} [put]
func (h *productHandlerImpl) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	var reqBody entity.ProductUpdateRequest
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
	if err := h.service.Product().UpdateProduct(r.Context(), productID, &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProductHandler godoc
//
//	@Summary		delete product endpoint
//	@Description	delete product by id
//	@Accept			json
//	@Produce		json
//	@Tags			Product
//	@Param			id	path	string	true	"product id"
//	@Security		Bearer
//	@Success		204
//	@Failure		500
//	@Router			/product/{id} [delete]
func (h *productHandlerImpl) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	if err := h.service.Product().DeleteProduct(r.Context(), productID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ExistsProductHandler godoc
//
//	@Summary		exists product endpoint
//	@Description	check if product exists by slug
//	@Accept			json
//	@Produce		json
//	@Tags			Product
//	@Param			slug	path	string	true	"product slug"
//	@Success		200
//	@Failure		404
//	@Failure		500
//	@Router			/product/exists/{slug} [get]
func (h *productHandlerImpl) ExistsProductHandler(w http.ResponseWriter, r *http.Request) {
	productSlug := r.PathValue("slug")
	exists, err := h.service.Product().ExistsProduct(r.Context(), productSlug)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
