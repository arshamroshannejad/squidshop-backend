package repository

import (
	"context"
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
)

type productCommentLikeRepositoryImpl struct {
	db *sql.DB
}

func NewProductCommentLikeRepository(db *sql.DB) domain.ProductCommentLikeRepository {
	return &productCommentLikeRepositoryImpl{
		db: db,
	}
}

func (r *productCommentLikeRepositoryImpl) Create(ctx context.Context, commentID, currentUser string, productCommentLike *entity.PostCommentLikeCreateUpdate) error {
	const createProductCommentLikeQuery = `INSERT INTO product_comment_likes (comment_id, user_id, vote) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`
	args := []any{commentID, currentUser, productCommentLike.Vote}
	_, err := r.db.ExecContext(ctx, createProductCommentLikeQuery, args...)
	return err
}

func (r *productCommentLikeRepositoryImpl) Update(ctx context.Context, commentID, currentUser string, productCommentLike *entity.PostCommentLikeCreateUpdate) error {
	const updateProductCommentLikeQuery = `UPDATE product_comment_likes SET vote = $1 WHERE comment_id = $2 AND user_id = $3`
	args := []any{productCommentLike.Vote, commentID, currentUser}
	_, err := r.db.ExecContext(ctx, updateProductCommentLikeQuery, args...)
	return err
}

func (r *productCommentLikeRepositoryImpl) Delete(ctx context.Context, commentID, currentUser string) error {
	const deleteProductCommentLikeQuery = `DELETE FROM product_comment_likes WHERE comment_id = $1 AND user_id = $2`
	args := []any{commentID, currentUser}
	_, err := r.db.ExecContext(ctx, deleteProductCommentLikeQuery, args...)
	return err
}
