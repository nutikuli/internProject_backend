package account

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
)
type AccountUsecase interface {
	AccountCustomerfile(ctx context.Context) ([]*dtos.CustomerAccountFileRes, int, error)
	AccountStorefile(ctx context.Context) ([]*_storeDtos.StoreWithFileRes, int, error)
}