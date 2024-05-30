package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/product"
	"github.com/nutikuli/internProject_backend/internal/models/product/entities"
	"github.com/nutikuli/internProject_backend/internal/models/product/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type productRepo struct {
	db *sqlx.DB
}

func NewproductRepository(db *sqlx.DB) product.ProductRepository {
	return &productRepo{
		db,
	}
}

func (s *productRepo) CreateProduct(ctx context.Context, req entities.ProductCreateReq) (*int64, error) {
	args := utils.Array{
		req.Name,
		req.Detail,
		req.Price,
		req.Status,
		req.ProductAvatar,
		req.Stock,
		req.CategoryId,
		req.StoreId,
	}

	res, err := s.db.ExecContext(ctx, repository_query.InsertProduct, args...)
	if err != nil {
		return nil, err
	}

	newProductId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &newProductId, nil

}

func (s *productRepo) GetProductById(ctx context.Context, productId *int64) (*entities.Product, error) {
	var product *entities.Product
	err := s.db.GetContext(ctx, product, repository_query.GetProductById, productId)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productRepo) GetProductsByStoreId(ctx context.Context, storeId *int64) ([]*entities.Product, error) {
	var products []*entities.Product
	err := s.db.SelectContext(ctx, products, repository_query.GetProductsByStoreId, storeId)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// DeleteProductById implements product.ProductRepository.
func (s *productRepo) DeleteProductById(ctx context.Context, productId *int64) error {
	res, err := s.db.ExecContext(ctx, repository_query.DeleteProductById, productId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("Can't Delete, Product not found")
	}

	return nil
}
