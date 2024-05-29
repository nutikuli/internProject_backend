package bank

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/bank/entities"
)

type BankRepository interface {
	GetBanks(ctx context.Context) (*entities.Bank, error)
	GetBankById(ctx context.Context, id *int64) (*entities.Bank, error)
	//INSERT
	CreateBank(ctx context.Context, bankdata *entities.BankCreatedReq) (*int64, error)
}
