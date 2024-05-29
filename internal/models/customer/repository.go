package customer

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerRepository interface {
	GetAccountCustomers(ctx context.Context) (*entities.Customer, error)
	GetCustomerById(ctx context.Context, id *int64) (*entities.Customer, error)
	CreateCustomerAccount(ctx context.Context, user *entities.CustomerRegister) (*int64, error)
}
