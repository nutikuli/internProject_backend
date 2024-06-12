package order_product

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
)

type OrderProductRepository interface {
	CreateOrderProduct(ctx context.Context, orderId int64, order *entities.OrderProductCreateReq) (*int64, error)
	GetOrderProductByOrderId(ctx context.Context, orderId int64) ([]*entities.OrderProduct, error)
	GetOrderProductByProductId(ctx context.Context, productId int64) (*entities.OrderProduct, error)
}
