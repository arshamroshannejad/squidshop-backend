package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type productServiceImpl struct {
	productRepository domain.ProductRepository
	logger            *slog.Logger
	config            *config.Config
}

func NewProductService(productRepository domain.ProductRepository, logger *slog.Logger, config *config.Config) domain.ProductService {
	return &productServiceImpl{
		productRepository: productRepository,
		logger:            logger,
		config:            config,
	}
}

func (s *productServiceImpl) GetAllProducts(ctx context.Context) ([]model.Products, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	products, err := s.productRepository.GetAll(ctx)
	if err != nil {
		s.logger.Error("failed to get all products", "error", err)
		return nil, err
	}
	for i := range products {
		if products[i].MainImage != nil {
			products[i].MainImage = helper.BuildMediaURL(s.config, products[i].MainImage)
		}
	}
	return products, nil
}

func (s *productServiceImpl) GetProductByID(ctx context.Context, productID string) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	product, err := s.productRepository.GetByID(ctx, productID)
	if err != nil {
		s.logger.Error("failed to get product by id", "error", err)
		return nil, err
	}
	if product.Images != nil {
		for i := range product.Images {
			product.Images[i].ImageURL = *helper.BuildMediaURL(s.config, &product.Images[i].ImageURL)
		}
	}
	return product, nil
}

func (s *productServiceImpl) GetProductBySlug(ctx context.Context, productSlug string) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	product, err := s.productRepository.GetBySlug(ctx, productSlug)
	if err != nil {
		s.logger.Error("failed to get product by slug", "error", err)
		return nil, err
	}
	if product.Images != nil {
		for i := range product.Images {
			product.Images[i].ImageURL = *helper.BuildMediaURL(s.config, &product.Images[i].ImageURL)
		}
	}
	return product, nil
}

func (s *productServiceImpl) CreateProduct(ctx context.Context, product *entity.ProductCreateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productRepository.Create(ctx, product); err != nil {
		s.logger.Error("failed to create product", "error", err)
		return err
	}
	return nil
}

func (s *productServiceImpl) UpdateProduct(ctx context.Context, productID string, product *entity.ProductUpdateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productRepository.Update(ctx, productID, product); err != nil {
		s.logger.Error("failed to update product", "error", err)
		return err
	}
	return nil
}

func (s *productServiceImpl) DeleteProduct(ctx context.Context, productID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.productRepository.Delete(ctx, productID); err != nil {
		s.logger.Error("failed to delete product", "error", err)
		return err
	}
	return nil
}

func (s *productServiceImpl) ExistsProduct(ctx context.Context, productSlug string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	exists, err := s.productRepository.Exists(ctx, productSlug)
	if err != nil {
		s.logger.Error("failed to check product existence", "error", err)
		return false, err
	}
	return exists, nil
}
