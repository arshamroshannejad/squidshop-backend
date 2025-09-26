package handler

import (
	"fmt"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/go-playground/validator/v10"
)

type productImageHandlerImpl struct {
	service   domain.Service
	validator *validator.Validate
}

func NewProductImageHandler(service domain.Service, validator *validator.Validate) domain.ProductImageHandler {
	return &productImageHandlerImpl{
		service:   service,
		validator: validator,
	}
}

// CreateProductImageHandler godoc
//
//	@Summary		create product image endpoint
//	@Description	create product image by product id
//	@Accept			multipart/form-data
//	@Produce		json
//	@Tags			Product Image
//	@Param			id		path		string	true	"product id"
//	@Param			images	formData	file	true	"product images"
//	@Security		Bearer
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/product/image/{id} [post]
func (h *productImageHandlerImpl) CreateProductImageHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	if err := r.ParseMultipartForm(1 << 20); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "no file uploaded"}`)))
		return
	}
	images, err := h.service.S3().UploadFiles(r.Context(), files, "products")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := h.service.ProductImage().CreateProductImage(r.Context(), productID, images); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
