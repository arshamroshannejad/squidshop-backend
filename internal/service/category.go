package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type categoryServiceImpl struct {
	categoryRepository domain.CategoryRepository
	logger             *slog.Logger
}

func NewCategoryService(categoryRepository domain.CategoryRepository, logger *slog.Logger) domain.CategoryService {
	return &categoryServiceImpl{
		categoryRepository: categoryRepository,
		logger:             logger,
	}
}

func (s *categoryServiceImpl) GetAllCategories(ctx context.Context) (*[]model.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	categories, err := s.categoryRepository.GetAll(ctx)
	if err != nil {
		s.logger.Error("failed to get all categories", "error", err)
		return nil, err
	}
	return categories, nil
}

func (s *categoryServiceImpl) CreateCategory(ctx context.Context, category *entity.CategoryCreateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.categoryRepository.Create(ctx, category); err != nil {
		s.logger.Error("failed to create category", "error", err)
		return err
	}
	return nil
}

func (s *categoryServiceImpl) UpdateCategory(ctx context.Context, categoryID string, category *entity.CategoryUpdateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.categoryRepository.Update(ctx, categoryID, category); err != nil {
		s.logger.Error("failed to update category", "error", err)
		return err
	}
	return nil
}

func (s *categoryServiceImpl) DeleteCategory(ctx context.Context, categoryID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.categoryRepository.Delete(ctx, categoryID); err != nil {
		s.logger.Error("failed to delete category", "error", err)
		return err
	}
	return nil
}

func (s *categoryServiceImpl) ExistsCategory(ctx context.Context, category *entity.CategoryQueryParamRequest) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	exists, err := s.categoryRepository.Exists(ctx, category)
	if err != nil {
		s.logger.Error("failed to check category existence", "error", err)
		return false, err
	}
	return exists, nil
}
