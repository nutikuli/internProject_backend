package account

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
)

type AccountRepository interface {
	GetAccountCustomers(ctx context.Context) (*entities.Account, error)
	GetAccountStores(ctx context.Context) (*entities.StoreAccount, error)
	GetAccountAdmins(ctx context.Context) (*entities.Admin, error)
	GetAccountAdminById(ctx context.Context, id *int64) (*entities.Admin, error)
	GetAccountStoreById(ctx context.Context, id *int64) (*entities.StoreAccount, error)
	GetAccountCustomerById(ctx context.Context, id *int64) (*entities.Account, error)
}