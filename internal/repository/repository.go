package repository

import (
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
)

type repositoryImpl struct {
	userRepository           domain.UserRepository
	categoryRepository       domain.CategoryRepository
	productRepository        domain.ProductRepository
	productRatingRepository  domain.ProductRatingRepository
	productImageRepository   domain.ProductImageRepository
	productCommentRepository domain.ProductCommentRepository
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repositoryImpl{
		userRepository:           NewUserRepository(db),
		categoryRepository:       NewCategoryRepository(db),
		productRepository:        NewProductRepository(db),
		productRatingRepository:  NewProductRatingRepository(db),
		productImageRepository:   NewProductImageRepository(db),
		productCommentRepository: NewProductCommentRepository(db),
	}
}

func (r *repositoryImpl) User() domain.UserRepository {
	return r.userRepository
}

func (r *repositoryImpl) Category() domain.CategoryRepository {
	return r.categoryRepository
}

func (r *repositoryImpl) Product() domain.ProductRepository {
	return r.productRepository
}

func (r *repositoryImpl) ProductRating() domain.ProductRatingRepository {
	return r.productRatingRepository
}

func (r *repositoryImpl) ProductImage() domain.ProductImageRepository {
	return r.productImageRepository
}

func (r *repositoryImpl) ProductComment() domain.ProductCommentRepository {
	return r.productCommentRepository
}
