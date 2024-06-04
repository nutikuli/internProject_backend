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

// UpdateOrderTransportDetail implements order.OrderRepository.

func NewOrderRepository(db *sqlx.DB) order.OrderRepository {
	return &OrderRepo{
		db: db,
	}
}

func (c *OrderRepo) GetOrdersByCustomerId(ctx context.Context, Id *int64) ([]*entities.Order, error) {
	var order = make([]*entities.Order, 0)

	err := c.db.SelectContext(ctx, &order, repository_query.SQL_get_order_by_customerId, *Id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return order, nil
}

func (c *OrderRepo) GetOrderById(ctx context.Context, Id *int64) (*entities.Order, error) {
	order := &entities.Order{}

	err := c.db.GetContext(ctx, order, repository_query.SQL_get_order_by_Id, *Id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return order, nil
}

func (c *OrderRepo) GetOrdersByStoreId(ctx context.Context, Id *int64) ([]*entities.Order, error) {
	orders := make([]*entities.Order, 0)

	err := c.db.SelectContext(ctx, &orders, repository_query.SQL_get_order_by_storeId, *Id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return orders, nil
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

	userId, _ := res.LastInsertId()

	return &userId, nil
}

func (c *OrderRepo) UpdateOrderTransportDetail(ctx context.Context, orderId int64, order *entities.OrderTransportDetailReq) error {

	args := utils.Array{
		order.DeliveryType,
		order.ParcelNumber,
		order.SentDate,
		orderId,
	}

	log.Info(args)

	res, err := c.db.ExecContext(ctx, repository_query.SQL_update_order_transport_detail, args...)
	if err != nil {
		log.Info(err)
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Info(err)
		return err
	}

	if affected == 0 {
		log.Info("No order was updated ")
		return nil
	}

	return nil
}

// UpdateOrderStatus implements order.OrderRepository.
func (c *OrderRepo) UpdateOrderStatus(ctx context.Context, orderId int64, order *entities.OrderStateReq) error {
	args := utils.Array{
		order.State,
		orderId,
	}

	log.Info(args)

	res, err := c.db.ExecContext(ctx, repository_query.SQL_update_order_state, args...)
	if err != nil {
		log.Info(err)
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Info(err)
		return err
	}

	if affected == 0 {
		log.Info("No order transport detail was updated ")
		return nil
	}

	return nil
}
