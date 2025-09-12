package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/redis/go-redis/v9"
)

type otpService struct {
	redisDB *redis.Client
	ttl     time.Duration
}

func NewUserOTPService(redisClient *redis.Client, ttl time.Duration) domain.OTPService {
	return &otpService{
		redisDB: redisClient,
		ttl:     ttl,
	}
}

func (s *otpService) Generate(ctx context.Context, phone string) (string, error) {
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	if err := s.redisDB.Set(ctx, "otp:"+phone, otp, s.ttl).Err(); err != nil {
		return "", err
	}
	return otp, nil
}

func (s *otpService) Verify(ctx context.Context, phone, code string) (bool, error) {
	key := "otp:" + phone
	storedOtp, err := s.redisDB.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, err
	}
	if storedOtp != code {
		return false, nil
	}
	_ = s.redisDB.Del(ctx, key).Err()
	return true, nil
}
