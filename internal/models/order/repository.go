package order

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/order/entities"
)

type OrderRepository interface {
	GetOrder(ctx context.Context) (*entities.Order, error)
	GetOrderByStoreId(ctx context.Context, id *int64) (*entities.Order, error)
}
