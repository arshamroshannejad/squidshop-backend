package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
)

type productImageRepositoryImpl struct {
	db *sql.DB
}

func NewProductImageRepository(db *sql.DB) domain.ProductImageRepository {
	return &productImageRepositoryImpl{
		db: db,
	}
}

func (r *productImageRepositoryImpl) Create(ctx context.Context, productID string, imagesURL []string) error {
	values := make([]string, 0, len(imagesURL))
	args := make([]any, 0, len(imagesURL)+1)
	args = append(args, productID)
	for i, img := range imagesURL {
		values = append(values, fmt.Sprintf("($1, $%d)", i+2))
		args = append(args, img)
	}
	query := fmt.Sprintf(`INSERT INTO product_images (product_id, image_url) VALUES %s`, strings.Join(values, ", "))
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}
