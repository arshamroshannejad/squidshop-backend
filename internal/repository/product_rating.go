package repository

import (
	"context"
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
)

type productRatingRepositoryImpl struct {
	db *sql.DB
}

func NewProductRatingRepository(db *sql.DB) domain.ProductRatingRepository {
	return &productRatingRepositoryImpl{
		db: db,
	}
}

func (r *productRatingRepositoryImpl) CreateOrUpdate(ctx context.Context, productID, userID string, rate *entity.ProductRatingRequest) error {
	const createOrUpdateProductRateQuery string = `
		INSERT INTO product_ratings (product_id, user_id, rating)
		VALUES ($1, $2, $3)
		ON CONFLICT (product_id, user_id)
		DO UPDATE SET rating = $3
	`
	args := []any{productID, userID, rate.Rate}
	_, err := r.db.ExecContext(ctx, createOrUpdateProductRateQuery, args...)
	return err
}

func (r *productRatingRepositoryImpl) Delete(ctx context.Context, productID, userID string) error {
	const deleteProductRateQuery string = "DELETE FROM product_ratings WHERE product_id = $1 AND user_id = $2"
	args := []any{productID, userID}
	_, err := r.db.ExecContext(ctx, deleteProductRateQuery, args...)
	return err
}
