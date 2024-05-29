package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/bank"
	"github.com/nutikuli/internProject_backend/internal/models/bank/entities"
	"github.com/nutikuli/internProject_backend/internal/models/bank/repository/repository_query"
)

type BankRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) bank.BankRepository {
	return &BankRepo{
		db: db,
	}
}

func (a *BankRepo) GetBanks(ctx context.Context) (*entities.Bank, error) {
	var bank entities.Bank

	err := a.db.GetContext(ctx, &bank, repository_query.SQL_get_bank)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &bank, nil
}

func (a *BankRepo) GetBankById(ctx context.Context, id *int64) (*entities.Bank, error) {
	var bank entities.Bank

	err := a.db.GetContext(ctx, &bank, repository_query.SQL_get_bank_by_id, *id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &bank, nil
}

func (a *BankRepo) CreateBank(ctx context.Context, bankdata *entities.BankCreatedReq) (*int64, error) {

	bankstatus := 1
	res, err := a.db.ExecContext(ctx, repository_query.SQL_insert_bank, bankdata.Name, bankdata.AccNumber, bankdata.AccName, bankstatus, bankdata.AvartarUrl, bankdata.StoreId)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	createdId, err := res.LastInsertId()
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &createdId, nil
}
