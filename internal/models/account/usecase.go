package account

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
	_accDtos"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
)
type AccountUsecase interface {
	AccountCustomerfile(ctx context.Context) ([]*dtos.CustomerAccountFileRes, int, error)
	AccountStorefile(ctx context.Context) ([]*_storeDtos.StoreWithFileRes, int, error)
	AccountAdminfile(ctx context.Context) ([]*_adminDtos.AdminFileRes, int, error)
	Login(ctx context.Context, req *entities.UsersLogin) (*_accDtos.UserToken, int, error)
	Register(ctx context.Context, req *entities.UserCreatedReq) (*_accDtos.UsersRegisteredRes, int, error)
}