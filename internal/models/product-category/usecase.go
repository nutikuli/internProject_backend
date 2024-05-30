package product_category

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/product-category/entities"
)

type ProductCategoryUsecase interface {
	OnCreateProductCategory(ctx context.Context, req entities.ProductCategoryCreatedReq) (*int64, int, error)
	OnGetProductCategoryById(ctx context.Context, categoryId *int64) (*entities.ProductCategory, int, error)
	OnGetProductCategoriesByStoreId(ctx context.Context, storeId *int64) ([]*entities.ProductCategory, int, error)
	OnDeleteProductCategoryById(ctx context.Context, categoryId *int64) (int, error)
}
