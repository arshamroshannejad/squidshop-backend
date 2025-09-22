package service

import (
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/redis/go-redis/v9"
)

type serviceImpl struct {
	userRepository          domain.UserRepository
	categoryRepository      domain.CategoryRepository
	productRepository       domain.ProductRepository
	productRatingRepository domain.ProductRatingRepository
	redisDB                 *redis.Client
	logger                  *slog.Logger
	cfg                     *config.Config
}

func NewService(repositories domain.Repository, redisDB *redis.Client, logger *slog.Logger, cfg *config.Config) domain.Service {
	return &serviceImpl{
		userRepository:          repositories.User(),
		categoryRepository:      repositories.Category(),
		productRepository:       repositories.Product(),
		productRatingRepository: repositories.ProductRating(),
		redisDB:                 redisDB,
		logger:                  logger,
		cfg:                     cfg,
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

func (s *serviceImpl) Category() domain.CategoryService {
	return NewCategoryService(s.categoryRepository, s.logger)
}

func (s *serviceImpl) Product() domain.ProductService {
	return NewProductService(s.productRepository, s.logger)
}

func (s *serviceImpl) ProductRating() domain.ProductRatingService {
	return NewProductRatingService(s.productRatingRepository, s.logger)
}
