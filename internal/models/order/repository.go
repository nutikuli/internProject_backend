package order

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/order/entities"
)

type OrderRepository interface {
	GetOrderByCustomer(ctx context.Context) (*entities.Order, error)
	GetOrderById(ctx context.Context, Id *int64) (*entities.Order, error)
	CreateOrder(ctx context.Context, order *entities.OrderCreate) (*int64, error)
}
