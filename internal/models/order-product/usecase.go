package order_product

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
)

type OrderProductUsecase interface {
	OnCreateOrderProduct(ctx context.Context, orders []*entities.OrderProductCreateReq) ([]*int64, int, error)
}
