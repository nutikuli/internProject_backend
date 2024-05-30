package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/order"
	"github.com/nutikuli/internProject_backend/internal/models/order/entities"
	"github.com/nutikuli/internProject_backend/internal/models/order/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) order.OrderRepository {
	return &OrderRepo{
		db: db,
	}
}

func (c *OrderRepo) GetOrderByCustomerId(ctx context.Context, id *int64) (*entities.Order, error) {
	var order = &entities.Order{}

	err := c.db.GetContext(ctx, order, repository_query.SQL_get_order_by_customerId, id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return order, nil
}

func (c *OrderRepo) GetOrderById(ctx context.Context, Id *int64) (*entities.Order, error) {
	order := &entities.Order{}

	err := c.db.SelectContext(ctx, order, repository_query.SQL_get_order_by_Id, *Id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return order, nil
}

func (c *OrderRepo) GetOrderByStoreId(ctx context.Context, Id *int64) (*entities.Order, error) {
	order := &entities.Order{}

	err := c.db.SelectContext(ctx, order, repository_query.SQL_get_order_by_storeId, *Id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return order, nil
}

func (c *OrderRepo) CreateOrder(ctx context.Context, order *entities.OrderCreate) (*int64, error) {

	args := utils.Array{
		order.TotalAmount,
		order.Topic,
		order.SumPrice,
		order.State,
		order.DeliveryType,
		order.ParcelNumber,
		order.SentDate,
		order.CustomerId,
		order.StoreId,
		order.BankId,
	}

	log.Info(args)

	res, err := c.db.ExecContext(ctx, repository_query.SQL_create_order, args...)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	userId, _ := res.RowsAffected()

	return &userId, nil
}
