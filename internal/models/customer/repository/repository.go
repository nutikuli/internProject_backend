package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	"github.com/nutikuli/internProject_backend/internal/models/customer/repository/repository_query"
)

type CustomerRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) customer.CustomerRepository {
	return &CustomerRepo{
		db: db,
	}
}

func (c *CustomerRepo) GetAccountCustomers(ctx context.Context) (*entities.Customer, error) {
	var customer entities.Customer

	err := c.db.GetContext(ctx, &customer, repository_query.SQL_get_account_customer, "customer")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &customer, nil
}

func (c *CustomerRepo) GetAccountCustomerById(ctx context.Context, id *int64) (*entities.Customer, error) {
	var accountbyid entities.Customer

	err := c.db.GetContext(ctx, &accountbyid, repository_query.SQL_get_account_customer_by_id, "customer", *id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &accountbyid, nil
}
