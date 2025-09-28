package domain

import (
	"context"
	"net/http"
)

type ProductImageRepository interface {
	Create(ctx context.Context, productID string, imagesURL []string) error
}

type ProductImageService interface {
	CreateProductImage(ctx context.Context, productID string, imagesURL []string) error
}

type ProductImageHandler interface {
	CreateProductImageHandler(w http.ResponseWriter, r *http.Request)
}
