package order_product

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
)

type OrderProductUsecase interface {
	OnCreateOrderProduct(ctx context.Context, orderId int64, orderProducts []*entities.OrderProductCreateReq) ([]*int64, int, error)
}
