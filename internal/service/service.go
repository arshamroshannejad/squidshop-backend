package service

import (
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/redis/go-redis/v9"
)

type serviceImpl struct {
	userRepository domain.UserRepository
	redisDB        *redis.Client
	logger         *slog.Logger
	cfg            *config.Config
}

func NewService(repositories domain.Repository, redisDB *redis.Client, logger *slog.Logger, cfg *config.Config) domain.Service {
	return &serviceImpl{
		userRepository: repositories.User(),
		redisDB:        redisDB,
		logger:         logger,
		cfg:            cfg,
	}
}

func (s *serviceImpl) User() domain.UserService {
	return NewUserService(s.userRepository, s.redisDB, s.logger, s.cfg)
}

func (s *serviceImpl) OTP() domain.OTPService {
	return NewUserOTPService(s.redisDB, time.Minute*2)
}

func (s *serviceImpl) Sms() domain.SmsService {
	return NewSmsService(s.logger, s.cfg)
}
