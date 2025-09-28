package service

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) GetByID(ctx context.Context, userID string) (*model.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *mockUserRepository) GetByPhone(ctx context.Context, phone string) (*model.User, error) {
	args := m.Called(ctx, phone)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *mockUserRepository) Create(ctx context.Context, user *entity.UserAuthRequest) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func newTestService(repo *mockUserRepository) domain.UserService {
	cfg := &config.Config{
		App: &config.App{
			Secret:        "testsecret",
			AccessHourTTL: time.Hour,
		},
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &userServiceImpl{
		userRepository: repo,
		logger:         logger,
		cfg:            cfg,
	}
}

func TestUserServiceImpl_GetUserByID(t *testing.T) {
	repo := new(mockUserRepository)
	svc := newTestService(repo)
	expected := &model.User{ID: "123", Phone: "555", IsAdmin: false}
	repo.On("GetByID", mock.Anything, "123").Return(expected, nil).Once()
	user, err := svc.GetUserByID(context.Background(), "123")
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
	repo.AssertExpectations(t)
}

func TestUserServiceImpl_GetUserByPhone(t *testing.T) {
	repo := new(mockUserRepository)
	svc := newTestService(repo)
	expected := &model.User{ID: "u1", Phone: "777", IsAdmin: true}
	repo.On("GetByPhone", mock.Anything, "777").Return(expected, nil).Once()
	user, err := svc.GetUserByPhone(context.Background(), "777")
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
	repo.AssertExpectations(t)
}

func TestUserServiceImpl_Create(t *testing.T) {
	repo := new(mockUserRepository)
	svc := newTestService(repo)
	req := &entity.UserAuthRequest{Phone: "1234"}
	repo.On("Create", mock.Anything, req).Return(nil).Once()
	err := svc.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUserServiceImpl_CreateError(t *testing.T) {
	repo := new(mockUserRepository)
	svc := newTestService(repo)
	req := &entity.UserAuthRequest{Phone: "1234"}
	repo.On("Create", mock.Anything, req).Return(errors.New("db error")).Once()
	err := svc.CreateUser(context.Background(), req)
	assert.Error(t, err)
	repo.AssertExpectations(t)
}

func TestUserServiceImpl_GenerateUserJwtToken(t *testing.T) {
	repo := new(mockUserRepository)
	svc := newTestService(repo)
	user := &model.User{ID: "u1", Phone: "111", IsAdmin: true}
	tokenStr, err := svc.GenerateUserJwtToken(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("testsecret"), nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)
	claims := token.Claims.(jwt.MapClaims)
	assert.Equal(t, "u1", claims["user_id"])
	assert.Equal(t, "111", claims["phone"])
	assert.Equal(t, true, claims["is_admin"])
}
