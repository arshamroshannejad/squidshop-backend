package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
)

type productCommentLikeServiceImpl struct {
	productCommentLikeRepository domain.ProductCommentLikeRepository
	logger                       *slog.Logger
}

func NewProductCommentLikeService(productCommentLikeRepository domain.ProductCommentLikeRepository, logger *slog.Logger) domain.ProductCommentLikeService {
	return &productCommentLikeServiceImpl{
		productCommentLikeRepository: productCommentLikeRepository,
		logger:                       logger,
	}
}

func (s *productCommentLikeServiceImpl) CreateProductCommentLike(ctx context.Context, commentID, currentUser string, vote *entity.PostCommentLikeCreateUpdate) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productCommentLikeRepository.Create(ctx, commentID, currentUser, vote); err != nil {
		s.logger.Error("failed to create product comment like", "error", err)
		return err
	}
	return nil
}

func (s *productCommentLikeServiceImpl) UpdateProductCommentLike(ctx context.Context, commentID, currentUser string, vote *entity.PostCommentLikeCreateUpdate) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productCommentLikeRepository.Update(ctx, commentID, currentUser, vote); err != nil {
		s.logger.Error("failed to update product comment like", "error", err)
		return err
	}
	return nil
}

func (s *productCommentLikeServiceImpl) DeleteProductCommentLike(ctx context.Context, commentID, currentUser string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productCommentLikeRepository.Delete(ctx, commentID, currentUser); err != nil {
		s.logger.Error("failed to delete product comment like", "error", err)
		return err
	}
	return nil
}
