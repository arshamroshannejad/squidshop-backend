package repository

import (
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
)

type repositoryImpl struct {
	userRepository domain.UserRepository
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repositoryImpl{
		userRepository: NewUserRepository(db),
	}
}

func (r *repositoryImpl) User() domain.UserRepository {
	return r.userRepository
}
