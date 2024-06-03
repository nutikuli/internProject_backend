package product

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/product/entities"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req entities.ProductCreateReq) (*int64, error)
	GetProductById(ctx context.Context, productId *int64) (*entities.Product, error)
	GetProductsByOrderId(ctx context.Context, orderId *int64) ([]*entities.Product, error)
	GetProductsByStoreId(ctx context.Context, storeId *int64) ([]*entities.Product, error)
	DeleteProductById(ctx context.Context, productId *int64) error
	UpdateProductById(ctx context.Context, productId int64, req *entities.ProductUpdateReq) error
	GetAllProducts(ctx context.Context) ([]*entities.Product, error)
}
