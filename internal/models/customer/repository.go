package customer

import (
	"context"

	_accountEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerRepository interface {
	GetAccountCustomers(ctx context.Context) ([]*entities.Customer, error)
	GetCustomerById(ctx context.Context, customerId int64) (*entities.Customer, error)
	CreateCustomerAccount(ctx context.Context, user *entities.CustomerRegisterReq) (*int64, error)
	UpdateCustomerPasswordById(ctx context.Context, customerdata *_accountEntities.UpdatePass) error
	UpdateCustomerById(ctx context.Context, userId int64, user *entities.CustomerUpdateReq) error
	DeleteCustomerById(ctx context.Context, Id int64) error
}
