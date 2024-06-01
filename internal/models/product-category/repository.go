package product_category

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/product-category/entities"
)

type ProductCategoryRepository interface {
	CreateProductCategory(ctx context.Context, req *entities.ProductCategoryCreatedReq) (*int64, error)
	GetProductCategoryById(ctx context.Context, categoryId int64) (*entities.ProductCategory, error)
	GetProductCategoriesByStoreId(ctx context.Context, storeId int64) ([]*entities.ProductCategory, error)
	DeleteProductCategoryById(ctx context.Context, categoryId int64) error
	UpdateProductCategoryById(ctx context.Context, categoryId int64, req *entities.ProductCategoryCreatedReq) error
}
