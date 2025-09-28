package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
)

type productRatingServiceImpl struct {
	productRatingRepository domain.ProductRatingRepository
	logger                  *slog.Logger
}

func NewProductRatingService(productRatingRepository domain.ProductRatingRepository, logger *slog.Logger) domain.ProductRatingService {
	return &productRatingServiceImpl{
		productRatingRepository: productRatingRepository,
		logger:                  logger,
	}
}

func (s *productRatingServiceImpl) CreateOrUpdateProductRating(ctx context.Context, productID, userID string, rate *entity.ProductRatingRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productRatingRepository.CreateOrUpdate(ctx, productID, userID, rate); err != nil {
		s.logger.Error("failed to create or update product rating", "error", err)
		return err
	}
	return nil
}

func (s *productRatingServiceImpl) DeleteProductRating(ctx context.Context, productID, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productRatingRepository.Delete(ctx, productID, userID); err != nil {
		s.logger.Error("failed to delete product rating", "error", err)
		return err
	}
	return nil
}
