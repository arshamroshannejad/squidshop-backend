package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type userServiceImpl struct {
	userRepository domain.UserRepository
	redisDB        *redis.Client
	logger         *slog.Logger
	cfg            *config.Config
}

func NewUserService(userRepository domain.UserRepository, redisDB *redis.Client, logger *slog.Logger, cfg *config.Config) domain.UserService {
	return &userServiceImpl{
		userRepository: userRepository,
		redisDB:        redisDB,
		logger:         logger,
		cfg:            cfg,
	}
}

func (s *userServiceImpl) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	user, err := s.userRepository.GetByID(ctx, userID)
	if err != nil {
		s.logger.Error("failed to get user with id", "error", err)
		return nil, err
	}
	return user, nil
}

func (s *userServiceImpl) GetUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	user, err := s.userRepository.GetByPhone(ctx, phone)
	if err != nil {
		s.logger.Error("failed to get user with phone", "error", err)
		return nil, err
	}
	return user, nil
}

func (s *userServiceImpl) CreateUser(ctx context.Context, user *entity.UserAuthRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.userRepository.Create(ctx, user); err != nil {
		s.logger.Error("failed to create user", "error", err)
		return err
	}
	return nil
}

func (s *userServiceImpl) GenerateUserJwtToken(ctx context.Context, user *model.User) (string, error) {
	exp := time.Now().Add(s.cfg.App.AccessHourTTL).Unix()
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"phone":    user.Phone,
		"is_admin": user.IsAdmin,
		"exp":      exp,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.cfg.App.Secret))
	if err != nil {
		s.logger.Error("failed to create access token", "error:", err)
	}
	return token, nil
}
