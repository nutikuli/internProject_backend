package repository

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accountEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	"github.com/nutikuli/internProject_backend/internal/models/customer/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type CustomerRepo struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) customer.CustomerRepository {
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

func (c *CustomerRepo) GetCustomerById(ctx context.Context, id *int64) (*entities.Customer, error) {
	var accountbyid entities.Customer

	err := c.db.GetContext(ctx, &accountbyid, repository_query.SQL_get_account_customer_by_id, "customer", *id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &accountbyid, nil
}

func (c *CustomerRepo) CreateCustomerAccount(ctx context.Context, user *entities.CustomerRegisterReq) (*int64, error) {

	args := utils.Array{
		user.Id,
		user.Name,
		user.Password,
		user.Phone,
		user.Location,
		user.Email,
		user.Role,
		user.Status,
	}

	log.Info(args)

	res, err := c.db.ExecContext(ctx, repository_query.SQL_create_account_customer, args...)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	userId, _ := res.RowsAffected()

	return &userId, nil
}

func (c *CustomerRepo) UpdateCustomerPasswordById(ctx context.Context, admindata *_accountEntities.UpdatePass) error {
	args := utils.Array{
		admindata.Id,
		admindata.Password,
		admindata.Role,
	}

	log.Info(args)

	res, err := c.db.ExecContext(ctx, repository_query.SQL_update_password_account_customer, args...)
	if err != nil {
		log.Error(err)
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *CustomerRepo) UpdateCustomerById(ctx context.Context, userId int64, user *entities.CustomerUpdateReq) error {
	args := utils.Array{
		user.Name,
		user.Password,
		user.Phone,
		user.Location,
		user.Email,
		user.Role,
		user.Status,
	}

	log.Info(args)

	res, err := r.db.ExecContext(ctx, repository_query.SQL_update_account_customer, args...)
	if err != nil {
		log.Error(err)
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (a *CustomerRepo) DeleteCustomerById(ctx context.Context, Id int64) error {
	res, err := a.db.ExecContext(ctx, repository_query.SQL_delete_account_customer, Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
