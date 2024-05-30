package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	product_category "github.com/nutikuli/internProject_backend/internal/models/product-category"
	"github.com/nutikuli/internProject_backend/internal/models/product-category/entities"
	"github.com/nutikuli/internProject_backend/internal/models/product-category/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type productCategoryRepo struct {
	db *sqlx.DB
}

func NewProductCategoryRepository(db *sqlx.DB) product_category.ProductCategoryRepository {
	return &productCategoryRepo{
		db,
	}
}

func (s *productCategoryRepo) CreateProductCategory(ctx context.Context, req entities.ProductCategoryCreatedReq) (*int64, error) {
	args := utils.Array{
		req.Name,
		req.Code,
		req.Detail,
		req.Status,
	}

	res, err := s.db.ExecContext(ctx, repository_query.InsertProductCategory, args...)
	if err != nil {
		return nil, err
	}

	newCategoryId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &newCategoryId, nil
}

func (s *productCategoryRepo) GetProductCategoryById(ctx context.Context, categoryId *int64) (*entities.ProductCategory, error) {
	var category *entities.ProductCategory
	err := s.db.GetContext(ctx, category, repository_query.GetProductCategoryById, categoryId)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *productCategoryRepo) GetProductCategoriesByStoreId(ctx context.Context, storeId *int64) ([]*entities.ProductCategory, error) {
	var categories []*entities.ProductCategory
	err := s.db.SelectContext(ctx, categories, repository_query.GetProductCategoriesByStoreId, storeId)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// DeleteProductCategoryById implements product_category.ProductCategoryRepository.
func (s *productCategoryRepo) DeleteProductCategoryById(ctx context.Context, categoryId *int64) error {
	res, err := s.db.ExecContext(ctx, repository_query.DeleteProductCategoryById, categoryId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("Error on deleting product category, category not found")
	}

	return nil
}
