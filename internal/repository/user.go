package repository

import (
	"context"
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) GetByPhone(ctx context.Context, phone string) (*model.User, error) {
	const getUserByPhoneQuery string = "SELECT id, phone, created_at FROM users WHERE phone = $1"
	args := []any{phone}
	row := r.db.QueryRowContext(ctx, getUserByPhoneQuery, args...)
	if err := row.Err(); err != nil {
		return nil, err
	}
	return collectUserRow(row)
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.UserAuthRequest) error {
	const createUserQuery string = "INSERT INTO users (phone) VALUES ($1) ON CONFLICT DO NOTHING "
	args := []any{user.Phone}
	_, err := r.db.ExecContext(ctx, createUserQuery, args...)
	return err
}

func collectUserRow(row *sql.Row) (*model.User, error) {
	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Phone,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
