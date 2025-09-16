package domain

import (
	"context"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type CategoryRepository interface {
	GetAll(ctx context.Context) (*[]model.Category, error)
	Create(ctx context.Context, category *entity.CategoryCreateRequest) error
	Update(ctx context.Context, categoryID string, category *entity.CategoryUpdateRequest) error
	Delete(ctx context.Context, categoryID string) error
	Exists(ctx context.Context, category *entity.CategoryQueryParamRequest) (bool, error)
}

type CategoryService interface {
	GetAllCategories(ctx context.Context) (*[]model.Category, error)
	CreateCategory(ctx context.Context, category *entity.CategoryCreateRequest) error
	UpdateCategory(ctx context.Context, categoryID string, category *entity.CategoryUpdateRequest) error
	DeleteCategory(ctx context.Context, categoryID string) error
	ExistsCategory(ctx context.Context, category *entity.CategoryQueryParamRequest) (bool, error)
}

type CategoryHandler interface {
	GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request)
	CreateCategoryHandler(w http.ResponseWriter, r *http.Request)
	UpdateCategoryHandler(w http.ResponseWriter, r *http.Request)
	DeleteCategoryHandler(w http.ResponseWriter, r *http.Request)
	ExistsCategoryHandler(w http.ResponseWriter, r *http.Request)
}
