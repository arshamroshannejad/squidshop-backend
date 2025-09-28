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

type categoryHandlerImpl struct {
	service   domain.Service
	validator *validator.Validate
}

func NewCategoryHandler(service domain.Service, validator *validator.Validate) domain.CategoryHandler {
	return &categoryHandlerImpl{
		service:   service,
		validator: validator,
	}
}

// GetAllCategoriesHandler godoc
//
//	@Summary		get categories endpoint
//	@Description	get all categories in db
//	@Accept			json
//	@Produce		json
//	@Tags			Category
//	@Success		200 {array} model.Category
//	@Failure		500
//	@Router			/category [get]
func (h *categoryHandlerImpl) GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.Category().GetAllCategories(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(categories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// CreateCategoryHandler godoc
//
//	@Summary		create category endpoint
//	@Description	create new category
//	@Accept			json
//	@Produce		json
//	@Tags			Category
//	@Param			request	body	entity.CategoryCreateRequest	true	"category data for create"
//	@Security		Bearer
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/category [post]
func (h *categoryHandlerImpl) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody entity.CategoryCreateRequest
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
	if err := h.service.Category().CreateCategory(r.Context(), &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateCategoryHandler godoc
//
//	@Summary		update category endpoint
//	@Description	update category by id
//	@Accept			json
//	@Produce		json
//	@Tags			Category
//	@Param			id		path	string							true	"category id"
//	@Param			request	body	entity.CategoryUpdateRequest	true	"category data for update"
//	@Security		Bearer
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/category/{id} [put]
func (h *categoryHandlerImpl) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID := r.PathValue("id")
	var reqBody entity.CategoryUpdateRequest
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
	if err := h.service.Category().UpdateCategory(r.Context(), categoryID, &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteCategoryHandler godoc
//
//	@Summary		delete category endpoint
//	@Description	delete category by id
//	@Accept			json
//	@Produce		json
//	@Tags			Category
//	@Param			id	path	string	true	"category id"
//	@Security		Bearer
//	@Success		204
//	@Failure		500
//	@Router			/category/{id} [delete]
func (h *categoryHandlerImpl) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID := r.PathValue("id")
	if err := h.service.Category().DeleteCategory(r.Context(), categoryID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ExistsCategoryHandler godoc
//
//	@Summary		exists category endpoint
//	@Description	check category exists with name and slug. at least one of them must be set
//	@Accept			json
//	@Produce		json
//	@Tags			Category
//	@Param			name	query	string	false	"category name"
//	@Param			slug	query	string	false	"category slug"
//	@Security		Bearer
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/category/exists [get]
func (h *categoryHandlerImpl) ExistsCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var reqParam entity.CategoryQueryParamRequest
	reqParam.Name = r.URL.Query().Get("name")
	reqParam.Slug = r.URL.Query().Get("slug")
	if err := h.validator.Struct(reqParam); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	exists, err := h.service.Category().ExistsCategory(r.Context(), &reqParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
