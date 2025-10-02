package domain

import (
	"context"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
)

type ProductCommentLikeRepository interface {
	Create(ctx context.Context, commentID, currentUser string, vote *entity.PostCommentLikeCreateUpdate) error
	Update(ctx context.Context, commentID, currentUser string, vote *entity.PostCommentLikeCreateUpdate) error
	Delete(ctx context.Context, commentID, currentUser string) error
}

type ProductCommentLikeService interface {
	CreateProductCommentLike(ctx context.Context, commentID, currentUser string, vote *entity.PostCommentLikeCreateUpdate) error
	UpdateProductCommentLike(ctx context.Context, commentID, currentUser string, vote *entity.PostCommentLikeCreateUpdate) error
	DeleteProductCommentLike(ctx context.Context, commentID, currentUser string) error
}

type ProductCommentLikeHandler interface {
	CreateProductCommentLikeHandler(w http.ResponseWriter, r *http.Request)
	UpdateProductCommentLikeHandler(w http.ResponseWriter, r *http.Request)
	DeleteProductCommentLikeHandler(w http.ResponseWriter, r *http.Request)
}
