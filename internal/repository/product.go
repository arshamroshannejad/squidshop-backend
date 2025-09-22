package repository

import (
	"context"
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
	"github.com/gosimple/slug"
)

type productRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepositoryImpl{
		db: db,
	}
}

func (r *productRepositoryImpl) GetAll(ctx context.Context) (*[]model.Product, error) {
	const getAllProductsQuery string = `
		SELECT 
		    p.id,
		    p.name,
		    p.slug,
		    p.description,
		    p.short_description,
		    p.price,
		    p.quantity,
		    p.created_at,
		    p.updated_at,
		    p.category_id,
		    COALESCE(AVG(pr.rating), 0) AS average_rating,
	    	COUNT(pr.rating) AS rating_count
		FROM
		    products p
		LEFT JOIN
			product_ratings pr ON p.id = pr.product_id
		GROUP BY
		    p.id
	`
	rows, err := r.db.QueryContext(ctx, getAllProductsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return collectProductsRows(rows)
}

func (r *productRepositoryImpl) GetByID(ctx context.Context, productID string) (*model.Product, error) {
	const getProductByIDQuery string = `
		SELECT 
		    p.id,
		    p.name,
		    p.slug,
		    p.description,
		    p.short_description,
		    p.price,
		    p.quantity,
		    p.created_at,
		    p.updated_at,
		    p.category_id,
			COALESCE(AVG(pr.rating), 0) AS average_rating,
	    	COUNT(pr.rating) AS rating_count
		FROM
		    products p
		LEFT JOIN
			product_ratings pr ON p.id = pr.product_id
		WHERE
		    p.id = $1
		GROUP BY
		    p.id
	`
	args := []any{productID}
	row := r.db.QueryRowContext(ctx, getProductByIDQuery, args...)
	return collectProductRow(row)
}

func (r *productRepositoryImpl) GetBySlug(ctx context.Context, productSlug string) (*model.Product, error) {
	const getProductBySlugQuery string = `
		SELECT 
		    p.id,
		    p.name,
		    p.slug,
		    p.description,
		    p.short_description,
		    p.price,
		    p.quantity,
		    p.created_at,
		    p.updated_at,
		    p.category_id,
			COALESCE(AVG(pr.rating), 0) AS average_rating,
	    	COUNT(pr.rating) AS rating_count
		FROM
		    products p
		LEFT JOIN
			product_ratings pr ON p.id = pr.product_id
		WHERE
		    p.slug = $1
		GROUP BY
		    p.id
	`
	args := []any{productSlug}
	row := r.db.QueryRowContext(ctx, getProductBySlugQuery, args...)
	return collectProductRow(row)
}

func (r *productRepositoryImpl) Create(ctx context.Context, product *entity.ProductCreateRequest) error {
	const createProductQuery string = "INSERT INTO products (name, slug, description, short_description, price, quantity, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	args := []any{product.Name, slug.Make(product.Slug), product.Description, product.ShortDescription, product.Price, product.Quantity, product.CategoryID}
	_, err := r.db.ExecContext(ctx, createProductQuery, args...)
	return err
}

func (r *productRepositoryImpl) Update(ctx context.Context, productID string, product *entity.ProductUpdateRequest) error {
	const updateProductQuery string = "UPDATE products SET name = $1, slug = $2, description = $3, short_description = $4, price = $5, quantity = $6, category_id = $7 WHERE id = $8"
	args := []any{product.Name, slug.Make(product.Slug), product.Description, product.ShortDescription, product.Price, product.Quantity, product.CategoryID, productID}
	_, err := r.db.ExecContext(ctx, updateProductQuery, args...)
	return err
}

func (r *productRepositoryImpl) Delete(ctx context.Context, productID string) error {
	const deleteProductQuery string = "DELETE FROM products WHERE id = $1"
	args := []any{productID}
	_, err := r.db.ExecContext(ctx, deleteProductQuery, args...)
	return err
}

func (r *productRepositoryImpl) Exists(ctx context.Context, productSlug string) (bool, error) {
	const existsProductQuery string = "SELECT EXISTS (SELECT 1 FROM products WHERE slug = $1)"
	args := []any{productSlug}
	var exists bool
	if err := r.db.QueryRowContext(ctx, existsProductQuery, args...).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

func collectProductsRows(rows *sql.Rows) (*[]model.Product, error) {
	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Description,
			&product.ShortDescription,
			&product.Price,
			&product.Quantity,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.CategoryID,
			&product.AverageRating,
			&product.RatingCount,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return &products, nil
}

func collectProductRow(row *sql.Row) (*model.Product, error) {
	var product model.Product
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Slug,
		&product.Description,
		&product.ShortDescription,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.CategoryID,
		&product.AverageRating,
		&product.RatingCount,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
