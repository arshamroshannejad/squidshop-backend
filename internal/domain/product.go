package domain

import (
	"context"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type ProductRepository interface {
	GetAll(ctx context.Context) (*[]model.Product, error)
	GetByID(ctx context.Context, productID string) (*model.Product, error)
	GetBySlug(ctx context.Context, productSlug string) (*model.Product, error)
	Create(ctx context.Context, product *entity.ProductCreateRequest) error
	Update(ctx context.Context, productID string, product *entity.ProductUpdateRequest) error
	Delete(ctx context.Context, productID string) error
	Exists(ctx context.Context, productSlug string) (bool, error)
}

type ProductService interface {
	GetAllProducts(ctx context.Context) (*[]model.Product, error)
	GetProductByID(ctx context.Context, productID string) (*model.Product, error)
	GetProductBySlug(ctx context.Context, productSlug string) (*model.Product, error)
	CreateProduct(ctx context.Context, product *entity.ProductCreateRequest) error
	UpdateProduct(ctx context.Context, productID string, product *entity.ProductUpdateRequest) error
	DeleteProduct(ctx context.Context, productID string) error
	ExistsProduct(ctx context.Context, productSlug string) (bool, error)
}

type ProductHandler interface {
	GetAllProductsHandler(w http.ResponseWriter, r *http.Request)
	GetProductByIDHandler(w http.ResponseWriter, r *http.Request)
	GetProductBySlugHandler(w http.ResponseWriter, r *http.Request)
	CreateProductHandler(w http.ResponseWriter, r *http.Request)
	UpdateProductHandler(w http.ResponseWriter, r *http.Request)
	DeleteProductHandler(w http.ResponseWriter, r *http.Request)
	ExistsProductHandler(w http.ResponseWriter, r *http.Request)
}
