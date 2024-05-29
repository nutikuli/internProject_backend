package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/account"

	//"github.com/nutikuli/internProject_backend/internal/models/account/repository"
	"github.com/nutikuli/internProject_backend/internal/models/account/repository/repository_query"

	adminstruct "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	customerstruct "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	storestruct "github.com/nutikuli/internProject_backend/internal/models/store/entities"
)

type AccountRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) account.AccountRepository {
	return &AccountRepo{
		db: db,
	}
}

func (a *AccountRepo) GetAccountCustomers(ctx context.Context) ([]*customerstruct.Customer, error) {
	var customer []*customerstruct.Customer

	err := a.db.GetContext(ctx, &customer, repository_query.SQL_get_account_customer, "CUSTOMER")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return customer, nil
}

func (a *AccountRepo) GetAccountStores(ctx context.Context) ([]*storestruct.Store, error) {
	var store []*storestruct.Store

	err := a.db.GetContext(ctx, &store, repository_query.SQL_get_account_storeaccount, "STORE")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return store, nil
}

func (a *AccountRepo) GetAccountAdmins(ctx context.Context) ([]*adminstruct.Admin, error) {
	var admin []*adminstruct.Admin

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account_admin, "ADMIN")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return admin, nil
}
