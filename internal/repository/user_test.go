package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepositoryImpl_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "failed to create mock database")
	defer db.Close()
	repo := NewUserRepository(db)
	fixedTime := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name         string
		userID       string
		setupMock    func()
		expectedUser *model.User
		expectedErr  error
	}{
		{
			name:   "Success - user found",
			userID: "user-123",
			setupMock: func() {
				rows := sqlmock.NewRows([]string{"id", "phone", "created_at", "is_admin"}).
					AddRow("user-123", "1234567890", fixedTime, false)
				mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").
					WithArgs("user-123").
					WillReturnRows(rows)
			},
			expectedUser: &model.User{
				ID:        "user-123",
				Phone:     "1234567890",
				CreatedAt: fixedTime,
				IsAdmin:   false,
			},
			expectedErr: nil,
		},
		{
			name:   "Error - user not found",
			userID: "user-999",
			setupMock: func() {
				mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").
					WithArgs("user-999").
					WillReturnError(sql.ErrNoRows)
			},
			expectedUser: nil,
			expectedErr:  sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			user, err := repo.GetByID(context.Background(), tt.userID)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedUser, user)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUserRepositoryImpl_GetByPhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "failed to create mock database")
	defer db.Close()
	repo := NewUserRepository(db)
	fixedTime := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name         string
		phone        string
		setupMock    func()
		expectedUser *model.User
		expectedErr  error
	}{
		{
			name:  "Success - user found",
			phone: "1234567890",
			setupMock: func() {
				rows := sqlmock.NewRows([]string{"id", "phone", "created_at", "is_admin"}).
					AddRow("user-1234", "1234567890", fixedTime, false)
				mock.ExpectQuery("SELECT \\* FROM users WHERE phone = \\$1").
					WithArgs("1234567890").
					WillReturnRows(rows)
			},
			expectedUser: &model.User{
				ID:        "user-1234",
				Phone:     "1234567890",
				CreatedAt: fixedTime,
				IsAdmin:   false,
			},
			expectedErr: nil,
		},
		{
			name:  "Error - user not found",
			phone: "1234567891",
			setupMock: func() {
				mock.ExpectQuery("SELECT \\* FROM users WHERE phone = \\$1").
					WithArgs("1234567891").
					WillReturnError(sql.ErrNoRows)
			},
			expectedUser: nil,
			expectedErr:  sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			user, err := repo.GetByPhone(context.Background(), tt.phone)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedUser, user)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUserRepositoryImpl_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "failed to create mock database")
	defer db.Close()
	repo := NewUserRepository(db)
	tests := []struct {
		name        string
		user        *entity.UserAuthRequest
		setupMock   func()
		expectedErr error
	}{
		{
			name: "Success - user created",
			user: &entity.UserAuthRequest{
				Phone: "+1234567890",
			},
			setupMock: func() {
				mock.ExpectExec("INSERT INTO users \\(phone\\) VALUES \\(\\$1\\) ON CONFLICT DO NOTHING").
					WithArgs("+1234567890").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name: "Success - phone already exists (conflict handled)",
			user: &entity.UserAuthRequest{
				Phone: "+1234567890",
			},
			setupMock: func() {
				mock.ExpectExec("INSERT INTO users \\(phone\\) VALUES \\(\\$1\\) ON CONFLICT DO NOTHING").
					WithArgs("+1234567890").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupMock != nil {
				tt.setupMock()
			}
			err := repo.Create(context.Background(), tt.user)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				if !errors.Is(tt.expectedErr, assert.AnError) {
					assert.ErrorIs(t, err, tt.expectedErr)
				}
			} else {
				assert.NoError(t, err)
			}
			if tt.setupMock != nil {
				assert.NoError(t, mock.ExpectationsWereMet())
			}
		})
	}
}

func TestCollectUserRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "failed to create mock database")
	defer db.Close()
	tests := []struct {
		name          string
		setupMock     func()
		expectedUser  *model.User
		expectedError error
	}{
		{
			name: "Success - collect user row",
			setupMock: func() {
				rows := sqlmock.NewRows([]string{"id", "phone", "created_at", "is_admin"}).
					AddRow("user-123", "+1234567890", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), true)
				mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").
					WithArgs("user-123").
					WillReturnRows(rows)
			},
			expectedUser: &model.User{
				ID:        "user-123",
				Phone:     "+1234567890",
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				IsAdmin:   true,
			},
			expectedError: nil,
		},
		{
			name: "Error - scan error (wrong number of columns)",
			setupMock: func() {
				rows := sqlmock.NewRows([]string{"id", "phone", "created_at"}).
					AddRow("user-123", "+1234567890", time.Now())
				mock.ExpectQuery("SELECT \\* FROM users WHERE id = \\$1").
					WithArgs("user-123").
					WillReturnRows(rows)
			},
			expectedUser:  nil,
			expectedError: sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			row := db.QueryRow("SELECT * FROM users WHERE id = $1", "user-123")
			user, err := collectUserRow(row)
			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser, user)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
