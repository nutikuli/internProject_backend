package order

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/order/entities"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entities.OrderCreate) (*int64, error)
	GetOrderByCustomerId(ctx context.Context, id *int64) (*entities.Order, error)
	GetOrderById(ctx context.Context, Id *int64) (*entities.Order, error)
	GetOrderByStoreId(ctx context.Context, Id *int64) (*entities.Order, error)
}
