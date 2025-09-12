package domain

import (
	"context"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type UserRepository interface {
	GetByPhone(ctx context.Context, phone string) (*model.User, error)
	Create(ctx context.Context, user *entity.UserAuthRequest) error
}

type UserService interface {
	GetUserByPhone(ctx context.Context, phone string) (*model.User, error)
	CreateUser(ctx context.Context, user *entity.UserAuthRequest) error
	GenerateUserJwtToken(ctx context.Context, userID, phone string) (string, error)
}

type UserHandler interface {
	AuthUserHandler(w http.ResponseWriter, r *http.Request)
	VerifyAuthUserHandler(w http.ResponseWriter, r *http.Request)
}
