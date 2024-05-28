package account

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
)

type AccountRepository interface {
	GetAccountCustomer(ctx context.Context) (*entities.Account, error)
	GetAccountStore(ctx context.Context) (*entities.Account, error)
	GetAccountAdmin(ctx context.Context) (*entities.Account, error)
}