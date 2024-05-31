package order_product

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
)

type OrderProductRepository interface {
	CreateOrder(ctx context.Context, order *entities.OrderProductCreateReq) (*int64, error)
	GetOrderProductByOrderId(ctx context.Context, orderId *int64) ([]*entities.OrderProduct, error)
}
