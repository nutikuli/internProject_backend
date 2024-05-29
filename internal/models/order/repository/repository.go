package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/order"
	"github.com/nutikuli/internProject_backend/internal/models/order/entities"
	"github.com/nutikuli/internProject_backend/internal/models/order/repository/repository_query"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) order.OrderRepository {
	return &OrderRepo{
		db: db,
	}
}

func (c *OrderRepo) GetOrder(ctx context.Context) (*entities.Order, error) {
	var order entities.Order

	err := c.db.GetContext(ctx, &order, repository_query.SQL_get_order, "order")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &order, nil
}

func (c *OrderRepo) GetOrderByStoreId(ctx context.Context, storeId *int64) (*entities.Order, error) {
	var orderbystoreid entities.Order

	err := c.db.GetContext(ctx, &orderbystoreid, repository_query.SQL_get_order_by_storeId, "order", *storeId)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &orderbystoreid, nil
}
