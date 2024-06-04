package bank

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/bank/entities"
)

type BankRepository interface {
	CreateBank(ctx context.Context, bankdata *entities.BankCreatedReq) (*int64, error)
	GetBankById(ctx context.Context, id int64) (*entities.Bank, error)
	GetBanks(ctx context.Context) ([]*entities.Bank, error)
	GetBanksByStoreId(ctx context.Context, storeId int64) ([]*entities.Bank, error)
	DeleteBankById(ctx context.Context, bankId int64) error
	UpdateBankById(ctx context.Context, bankId int64, bankdata *entities.BankUpdateReq) error
}
