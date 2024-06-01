package usecase

import (
	"context"
	"net/http"

	order_product "github.com/nutikuli/internProject_backend/internal/models/order-product"
	"github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
)

type order_productUsecase struct {
	order_productRepo order_product.OrderProductRepository
}

func NewOrderProductUsecase(order_productRepo order_product.OrderProductRepository) order_product.OrderProductUsecase {
	return &order_productUsecase{
		order_productRepo: order_productRepo,
	}
}

func (s *order_productUsecase) OnCreateOrderProducts(ctx context.Context, orderId int64, orders []*entities.OrderProductCreateReq) ([]*int64, int, error) {
	var createdOrderIDs = make([]*int64, 0)

	for _, order := range orders {
		newOrderID, err := s.order_productRepo.CreateOrderProduct(ctx, orderId, order)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		createdOrderIDs = append(createdOrderIDs, newOrderID)
	}

	return createdOrderIDs, http.StatusOK, nil
}
