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


func (a *AccountRepo) GetAccountCustomer(ctx context.Context) (*entities.Account, error) {
	var customer entities.Account

	err := a.db.GetContext(ctx, &customer, repository_query.SQL_get_account, "customer")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &customer, nil
}

func (a *AccountRepo) GetAccountStore(ctx context.Context) (*entities.Account, error) {
	var store entities.Account

	err := a.db.GetContext(ctx, &store, repository_query.SQL_get_account, "store")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &store, nil
}

func (a *AccountRepo) GetAccountAdmin(ctx context.Context) (*entities.Account, error) {
	var admin entities.Account

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account, "admin")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &admin, nil
}

func (a *AccountRepo) GetFileById(ctx context.Context, id *int64) (*entities.Account, error) {
	var accountbyid entities.Account

	err := a.db.GetContext(ctx, &accountbyid, repository_query.SQL_get_account, *id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &accountbyid, nil
}