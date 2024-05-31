package order

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/order/entities"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entities.OrderCreate) (*int64, error)
	GetOrdersByCustomerId(ctx context.Context, Id *int64) ([]*entities.Order, error)
	GetOrderById(ctx context.Context, Id *int64) (*entities.Order, error)
	GetOrdersByStoreId(ctx context.Context, Id *int64) ([]*entities.Order, error)
	UpdateOrderTransportDetail(ctx context.Context, order *entities.OrderTransportDetailReq) error
	UpdateOrderStatus(ctx context.Context, order *entities.OrderStateReq) error
}
