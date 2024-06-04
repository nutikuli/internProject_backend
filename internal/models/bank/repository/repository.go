package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/bank"
	"github.com/nutikuli/internProject_backend/internal/models/bank/entities"
	"github.com/nutikuli/internProject_backend/internal/models/bank/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
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
	bank := entities.Bank{}

	err := a.db.GetContext(ctx, &bank, repository_query.SQL_get_bank_by_id, id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &bank, nil
}

func (a *BankRepo) CreateBank(ctx context.Context, bankdata entities.BankCreatedReq) (*int64, error) {
	args := utils.Array{
		bankdata.AccName,
		bankdata.AccNumber,
		bankdata.AvatarUrl,
		bankdata.Status,
		bankdata.Name,
		bankdata.StoreId,
	}
	res, err := a.db.ExecContext(ctx, repository_query.SQL_insert_bank, args...)
	if err != nil {
		return nil, err
	}

	createdId, err := res.LastInsertId()
	if err != nil {
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
func (a *BankRepo) UpdateBankById(ctx context.Context, bankId int64, bankdata *entities.BankUpdateReq) error {
	args := utils.Array{
		bankId,
		bankdata.Name,
		bankdata.AccNumber,
		bankdata.AccName,
		bankdata.AvatarUrl,
		bankdata.Status,
		bankdata.StoreId,
	}
	res, err := a.db.ExecContext(ctx, repository_query.SQL_update_bank_by_id, args...)
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
