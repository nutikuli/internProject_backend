package customer

import (
	"context"

	_accountEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerRepository interface {
	GetAccountCustomers(ctx context.Context) (*entities.Customer, error)
	GetCustomerById(ctx context.Context, id *int64) (*entities.Customer, error)
	CreateCustomerAccount(ctx context.Context, user *entities.CustomerRegisterReq) (*int64, error)
	UpdateCustomerPasswordById(ctx context.Context, admindata *_accountEntities.UpdatePass) error
}
