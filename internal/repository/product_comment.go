package repository

import (
	"context"
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
)

type productCommentRepositoryImpl struct {
	db *sql.DB
}

func NewProductCommentRepository(db *sql.DB) domain.ProductCommentRepository {
	return &productCommentRepositoryImpl{
		db: db,
	}
}

func (r *productCommentRepositoryImpl) GetByID(ctx context.Context, productCommentID string) (*model.ProductComment, error) {
	const getProductCommentQuery = `SELECT * FROM product_comments WHERE id = $1`
	args := []any{productCommentID}
	row := r.db.QueryRowContext(ctx, getProductCommentQuery, args...)
	return collectProductCommentRow(row)
}

func (r *productCommentRepositoryImpl) Create(ctx context.Context, productID, currentUserID string, comment *entity.ProductCommentCreateRequest) error {
	const createProductCommentQuery = `INSERT INTO product_comments (product_id, user_id, parent_id, comment) VALUES ($1, $2, $3, $4)`
	args := []any{productID, currentUserID, comment.ParentID, comment.Comment}
	_, err := r.db.ExecContext(ctx, createProductCommentQuery, args...)
	return err
}

func (r *productCommentRepositoryImpl) Update(ctx context.Context, productCommentID string, comment *entity.ProductCommentUpdateRequest) error {
	const updateProductCommentQuery = `UPDATE product_comments SET comment = $1 WHERE id = $2`
	args := []any{comment.Comment, productCommentID}
	_, err := r.db.ExecContext(ctx, updateProductCommentQuery, args...)
	return err
}

func (r *productCommentRepositoryImpl) Delete(ctx context.Context, productCommentID string) error {
	const deleteProductCommentQuery = `DELETE FROM product_comments WHERE id = $1`
	args := []any{productCommentID}
	_, err := r.db.ExecContext(ctx, deleteProductCommentQuery, args...)
	return err
}

func collectProductCommentRow(row *sql.Row) (*model.ProductComment, error) {
	var productComment model.ProductComment
	err := row.Scan(
		&productComment.ID,
		&productComment.ProductID,
		&productComment.UserID,
		&productComment.ParentID,
		&productComment.Comment,
		&productComment.CreatedAt,
		&productComment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &productComment, nil
}
