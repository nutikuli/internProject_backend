package account

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
)

type AccountUsecase interface {
	AccountCustomerfile(ctx context.Context) ([]*dtos.CustomerAccountFileRes, int, error)
	AccountStorefile(ctx context.Context) ([]*_storeDtos.StoreWithFileRes, int, error)
	AccountAdminfile(ctx context.Context) ([]*_adminDtos.AdminFileRes, int, error)
	Login(c *fiber.Ctx, ctx context.Context, req *entities.UsersCredential) (*_accDtos.UserToken, interface{}, int, error)
	Register(ctx context.Context, req entities.AccountCredentialGetter) (*_accDtos.UsersRegisteredRes, *entities.UsersCredential, int, error)
}
