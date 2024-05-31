package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	order_product "github.com/nutikuli/internProject_backend/internal/models/order-product"
	"github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
	"github.com/nutikuli/internProject_backend/internal/models/order-product/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type OrderProductRepo struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) order_product.OrderProductRepository {
	return &OrderProductRepo{
		db: db,
	}
}

func (c *OrderProductRepo) CreateOrderProduct(ctx context.Context, orderId int64, orderProduct *entities.OrderProductCreateReq) (*int64, error) {

	args := utils.Array{
		orderId,
		orderProduct.ProductId,
		orderProduct.Quantity,
	}

	log.Info(args)

	res, err := c.db.ExecContext(ctx, repository_query.SQL_create_order_product, args...)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	userId, _ := res.RowsAffected()

	return &userId, nil
}

func (s *OrderProductRepo) GetOrderProductByOrderId(ctx context.Context, orderId int64) ([]*entities.OrderProduct, error) {
	var order_products []*entities.OrderProduct
	err := s.db.SelectContext(ctx, order_products, repository_query.SQL_get_order_products_by_order_id, orderId)
	if err != nil {
		return nil, err
	}

	return order_products, nil
}
