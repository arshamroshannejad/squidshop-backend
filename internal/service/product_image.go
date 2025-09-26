package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
)

type productImageServiceImpl struct {
	productImageRepository domain.ProductImageRepository
	logger                 *slog.Logger
}

func NewProductImageService(productImageRepository domain.ProductImageRepository, logger *slog.Logger) domain.ProductImageService {
	return &productImageServiceImpl{
		productImageRepository: productImageRepository,
		logger:                 logger,
	}
}

func (s *productImageServiceImpl) CreateProductImage(ctx context.Context, productID string, imagesURL []string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productImageRepository.Create(ctx, productID, imagesURL); err != nil {
		s.logger.Error("failed to create product images", "error", err)
		return err
	}
	return nil
}
