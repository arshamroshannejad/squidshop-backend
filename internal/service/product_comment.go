package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type productCommentServiceImpl struct {
	productCommentRepository domain.ProductCommentRepository
	logger                   *slog.Logger
}

func NewProductCommentService(productCommentRepository domain.ProductCommentRepository, logger *slog.Logger) domain.ProductCommentService {
	return &productCommentServiceImpl{
		productCommentRepository: productCommentRepository,
		logger:                   logger,
	}
}

func (s *productCommentServiceImpl) GetProductCommentByID(ctx context.Context, productCommentID string) (*model.ProductComment, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	productComment, err := s.productCommentRepository.GetByID(ctx, productCommentID)
	if err != nil {
		s.logger.Error("failed to get product comment", "error", err)
		return nil, err
	}
	return productComment, nil
}

func (s *productCommentServiceImpl) CreateProductComment(ctx context.Context, productID, currentUserID string, productComment *entity.ProductCommentCreateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productCommentRepository.Create(ctx, productID, currentUserID, productComment); err != nil {
		s.logger.Error("failed to create product comment", "error", err)
		return err
	}
	return nil
}

func (s *productCommentServiceImpl) UpdateProductComment(ctx context.Context, productCommentID string, productComment *entity.ProductCommentUpdateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productCommentRepository.Update(ctx, productCommentID, productComment); err != nil {
		s.logger.Error("failed to update product comment", "error", err)
		return err
	}
	return nil
}

func (s *productCommentServiceImpl) DeleteProductComment(ctx context.Context, productCommentID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productCommentRepository.Delete(ctx, productCommentID); err != nil {
		s.logger.Error("failed to delete product comment", "error", err)
		return err
	}
	return nil
}
