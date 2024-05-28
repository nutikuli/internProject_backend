package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/account/repository/repository_query"
)

type AccountRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) account.AccountRepository {
	return &AccountRepo{
		db: db,
	}
}


func (a *AccountRepo) GetAccountCustomers(ctx context.Context) (*entities.Account, error) {
	var customer entities.Account

	err := a.db.GetContext(ctx, &customer, repository_query.SQL_get_account_customer, "customer")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &customer, nil
}

func (a *AccountRepo) GetAccountStores(ctx context.Context) (*entities.StoreAccount, error) {
	var store entities.StoreAccount

	err := a.db.GetContext(ctx, &store, repository_query.SQL_get_account_storeaccount, "store")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &store, nil
}

func (a *AccountRepo) GetAccountAdmins(ctx context.Context) (*entities.Admin, error) {
	var admin entities.Admin

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account_admin, "admin")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &admin, nil
}

func (a *AccountRepo) GetAccountCustomerById(ctx context.Context, id *int64) (*entities.Account, error) {
	var accountbyid entities.Account

	err := a.db.GetContext(ctx, &accountbyid, repository_query.SQL_get_account_customer_by_id, "customer",*id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &accountbyid, nil
}

func (a *AccountRepo) GetAccountStoreById(ctx context.Context, id *int64) (*entities.StoreAccount, error) {
	var accountbyid entities.StoreAccount

	err := a.db.GetContext(ctx, &accountbyid, repository_query.SQL_get_account_storeaccount_by_id,"store", *id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &accountbyid, nil
}

func (a *AccountRepo) GetAccountAdminById(ctx context.Context, id *int64) (*entities.Admin, error) {
	var admin entities.Admin

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account_admin_by_id, "admin", *id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &admin, nil
}