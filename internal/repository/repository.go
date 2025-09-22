package repository

import (
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
)

type repositoryImpl struct {
	userRepository          domain.UserRepository
	categoryRepository      domain.CategoryRepository
	productRepository       domain.ProductRepository
	productRatingRepository domain.ProductRatingRepository
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repositoryImpl{
		userRepository:          NewUserRepository(db),
		categoryRepository:      NewCategoryRepository(db),
		productRepository:       NewProductRepository(db),
		productRatingRepository: NewProductRatingRepository(db),
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
