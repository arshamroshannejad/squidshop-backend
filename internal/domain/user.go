package domain

import (
	"context"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, userID string) (*model.User, error)
	GetByPhone(ctx context.Context, phone string) (*model.User, error)
	Create(ctx context.Context, user *entity.UserAuthRequest) error
}

type UserService interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*model.User, error)
	CreateUser(ctx context.Context, user *entity.UserAuthRequest) error
	GenerateUserJwtToken(ctx context.Context, user *model.User) (string, error)
}

type UserHandler interface {
	UserProfileHandler(w http.ResponseWriter, r *http.Request)
}
