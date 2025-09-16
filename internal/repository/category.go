package repository

import (
	"context"
	"database/sql"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/model"
	"github.com/gosimple/slug"
)

type categoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepositoryImpl{
		db: db,
	}
}

func (r *categoryRepositoryImpl) GetAll(ctx context.Context) (*[]model.Category, error) {
	const query = `SELECT id, name, slug, parent_id FROM categories`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories, err := collectCategoryRows(rows)
	if err != nil {
		return nil, err
	}
	tree := buildCategoryTree(*categories, nil)
	return &tree, nil
}

func (r *categoryRepositoryImpl) Create(ctx context.Context, category *entity.CategoryCreateRequest) error {
	const createCategoryQuery string = "INSERT INTO categories (name, slug, parent_id) VALUES ($1, $2, $3)"
	args := []any{category.Name, slug.Make(category.Slug), category.ParentID}
	_, err := r.db.ExecContext(ctx, createCategoryQuery, args...)
	return err
}

func (r *categoryRepositoryImpl) Update(ctx context.Context, categoryID string, category *entity.CategoryUpdateRequest) error {
	const updateCategoryQuery string = "UPDATE categories SET name = $1, slug = $2, parent_id = $3 WHERE id = $4"
	args := []any{category.Name, slug.Make(category.Slug), category.ParentID, categoryID}
	_, err := r.db.ExecContext(ctx, updateCategoryQuery, args...)
	return err
}

func (r *categoryRepositoryImpl) Delete(ctx context.Context, categoryID string) error {
	const deleteCategoryQuery string = "DELETE FROM categories WHERE id = $1"
	args := []any{categoryID}
	_, err := r.db.ExecContext(ctx, deleteCategoryQuery, args...)
	return err
}

func (r *categoryRepositoryImpl) Exists(ctx context.Context, category *entity.CategoryQueryParamRequest) (bool, error) {
	const existsCategoryQuery string = "SELECT EXISTS (SELECT 1 FROM categories WHERE name = $1 OR slug = $2)"
	args := []any{category.Name, category.Slug}
	var exists bool
	if err := r.db.QueryRowContext(ctx, existsCategoryQuery, args...).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

func collectCategoryRows(rows *sql.Rows) (*[]model.Category, error) {
	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Slug,
			&category.ParentID,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return &categories, nil
}

func buildCategoryTree(categories []model.Category, parentID *string) []model.Category {
	var tree []model.Category
	for _, c := range categories {
		if (c.ParentID == nil && parentID == nil) ||
			(c.ParentID != nil && parentID != nil && *c.ParentID == *parentID) {
			c.SubCategories = buildCategoryTree(categories, &c.ID)
			tree = append(tree, c)
		}
	}
	return tree
}
