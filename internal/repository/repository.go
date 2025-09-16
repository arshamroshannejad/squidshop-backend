package repository

import (
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
)

type repositoryImpl struct {
	userRepository     domain.UserRepository
	categoryRepository domain.CategoryRepository
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repositoryImpl{
		userRepository:     NewUserRepository(db),
		categoryRepository: NewCategoryRepository(db),
	}
}

func (r *repositoryImpl) User() domain.UserRepository {
	return r.userRepository
}

func (r *repositoryImpl) Category() domain.CategoryRepository {
	return r.categoryRepository
}
