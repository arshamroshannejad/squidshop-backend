package service

import (
	"log/slog"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/redis/go-redis/v9"
)

type serviceImpl struct {
	userRepository               domain.UserRepository
	categoryRepository           domain.CategoryRepository
	productRepository            domain.ProductRepository
	productRatingRepository      domain.ProductRatingRepository
	productImageRepository       domain.ProductImageRepository
	productCommentRepository     domain.ProductCommentRepository
	productCommentLikeRepository domain.ProductCommentLikeRepository
	redisDB                      *redis.Client
	logger                       *slog.Logger
	cfg                          *config.Config
}

func NewService(repositories domain.Repository, redisDB *redis.Client, logger *slog.Logger, cfg *config.Config) domain.Service {
	return &serviceImpl{
		userRepository:               repositories.User(),
		categoryRepository:           repositories.Category(),
		productRepository:            repositories.Product(),
		productRatingRepository:      repositories.ProductRating(),
		productImageRepository:       repositories.ProductImage(),
		productCommentRepository:     repositories.ProductComment(),
		productCommentLikeRepository: repositories.ProductCommentLike(),
		redisDB:                      redisDB,
		logger:                       logger,
		cfg:                          cfg,
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
	return NewProductService(s.productRepository, s.logger, s.cfg)
}

func (s *serviceImpl) ProductRating() domain.ProductRatingService {
	return NewProductRatingService(s.productRatingRepository, s.logger)
}

func (s *serviceImpl) ProductImage() domain.ProductImageService {
	return NewProductImageService(s.productImageRepository, s.logger)
}

func (s *serviceImpl) ProductComment() domain.ProductCommentService {
	return NewProductCommentService(s.productCommentRepository, s.logger)
}

func (s *serviceImpl) ProductCommentLike() domain.ProductCommentLikeService {
	return NewProductCommentLikeService(s.productCommentLikeRepository, s.logger)
}

func (s *serviceImpl) S3() domain.S3Service {
	return NewS3Service(s.cfg, s.logger)
}
