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

func NewFileRepository(db *sqlx.DB) order.OrderRepository {
	return &OrderRepo{
		db: db,
	}
}

func (c *OrderRepo) GetOrderByCustomer(ctx context.Context) (*entities.Order, error) {
	var orderbycustomer entities.Order

	err := c.db.GetContext(ctx, &orderbycustomer, repository_query.SQL_get_order_by_customer, "order_bycustomer")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &orderbycustomer, nil
}

func (c *OrderRepo) GetOrderById(ctx context.Context, Id *int64) (*entities.Order, error) {
	var orderbyid entities.Order

	err := c.db.GetContext(ctx, &orderbyid, repository_query.SQL_get_order_by_Id, "order", *Id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &orderbyid, nil
}

func (c *OrderRepo) CreateOrder(ctx context.Context, order *entities.OrderCreate) (*int64, error) {

	args := utils.Array{
		order.Id,
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
		order.CreatedAt,
		order.UpdatedAt,
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
