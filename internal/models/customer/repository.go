package customer

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerRepository interface {
	GetAccountCustomers(ctx context.Context) (*entities.Customer, error)
	GetAccountCustomerById(ctx context.Context, id *int64) (*entities.Customer, error)
	CreateAccountCustomer(ctx context.Context, user *entities.CustomerRegister) (*int64, error)
}
