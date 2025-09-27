package domain

import (
	"context"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type ProductCommentRepository interface {
	GetByID(ctx context.Context, productCommentID string) (*model.ProductComment, error)
	Create(ctx context.Context, productID, currentUserID string, comment *entity.ProductCommentCreateRequest) error
	Update(ctx context.Context, productCommentID string, comment *entity.ProductCommentUpdateRequest) error
	Delete(ctx context.Context, productCommentID string) error
}

type ProductCommentService interface {
	GetProductCommentByID(ctx context.Context, productCommentID string) (*model.ProductComment, error)
	CreateProductComment(ctx context.Context, productID, currentUserID string, comment *entity.ProductCommentCreateRequest) error
	UpdateProductComment(ctx context.Context, productCommentID string, comment *entity.ProductCommentUpdateRequest) error
	DeleteProductComment(ctx context.Context, productCommentID string) error
}

type ProductCommentHandler interface {
	CreateProductCommentHandler(w http.ResponseWriter, r *http.Request)
	UpdateProductCommentHandler(w http.ResponseWriter, r *http.Request)
	DeleteProductCommentHandler(w http.ResponseWriter, r *http.Request)
}
