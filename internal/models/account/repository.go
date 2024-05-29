//
package account

import (
	"context"
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	storestruct "github.com/nutikuli/internProject_backend/internal/models/store/entities"
	adminstruct "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	customerstruct "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type AccountRepository interface {
	GetAccountCustomers(ctx context.Context) ([]*customerstruct.Customer, error)
	GetAccountStores(ctx context.Context) ([]*storestruct.Store, error)
	GetAccountAdmins(ctx context.Context) ([]*adminstruct.Admin, error)
	FindUserAsPassport(ctx context.Context, email string) (*entities.UsersPassport, error)
	SignUsersAccessToken(req *entities.UserSignToken) (*dtos.UserToken, error)
	CreateUser(ctx context.Context, req *entities.UserCreatedReq) (*int64, error)
	// GetAccountAdminById(ctx context.Context, id *int64) (*entities.Admin, error)
	// GetAccountStoreById(ctx context.Context, id *int64) (*entities.StoreAccount, error)
	// GetAccountCustomerById(ctx context.Context, id *int64) (*entities.Account, error)
}