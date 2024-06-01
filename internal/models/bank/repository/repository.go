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

func NewBankRepository(db *sqlx.DB) bank.BankRepository {
	return &BankRepo{
		db: db,
	}
}

func (a *BankRepo) GetBanks(ctx context.Context) ([]*entities.Bank, error) {
	var banks = make([]*entities.Bank, 0)

	err := a.db.SelectContext(ctx, banks, repository_query.SQL_get_banks)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return banks, nil
}

func (a *BankRepo) GetBanksByStoreId(ctx context.Context, storeId int64) ([]*entities.Bank, error) {
	var banks = make([]*entities.Bank, 0)

	err := a.db.SelectContext(ctx, banks, repository_query.SQL_get_banks_by_store_id, storeId)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return banks, nil
}

func (a *BankRepo) GetBankById(ctx context.Context, id int64) (*entities.Bank, error) {
	var bank entities.Bank

	err := a.db.GetContext(ctx, &bank, repository_query.SQL_get_bank_by_id, id)
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

// DeleteBankById implements bank.BankRepository.
func (a *BankRepo) DeleteBankById(ctx context.Context, bankId int64) error {
	res, err := a.db.ExecContext(ctx, repository_query.SQL_delete_bank_by_id, bankId)
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
		log.Info("No bank was deleted")
		return nil
	}

	return nil
}

// UpdateBankById implements bank.BankRepository.
func (a *BankRepo) UpdateBankById(ctx context.Context, bankId int64, bankdata *entities.BankCreatedReq) error {
	res, err := a.db.ExecContext(ctx, repository_query.SQL_update_bank_by_id, bankdata.Name, bankdata.AccNumber, bankdata.AccName, bankdata.AvartarUrl, bankId)
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
		log.Info("No bank was updated")
		return nil
	}

	return nil
}
