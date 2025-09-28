package domain

import (
	"context"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
)

type ProductRatingRepository interface {
	CreateOrUpdate(ctx context.Context, productID, userID string, rate *entity.ProductRatingRequest) error
	Delete(ctx context.Context, productID, userID string) error
}

type ProductRatingService interface {
	CreateOrUpdateProductRating(ctx context.Context, productID, userID string, rate *entity.ProductRatingRequest) error
	DeleteProductRating(ctx context.Context, productID, userID string) error
}

type ProductRatingHandler interface {
	CreateOrUpdateProductRatingHandler(w http.ResponseWriter, r *http.Request)
	DeleteProductRatingHandler(w http.ResponseWriter, r *http.Request)
}
