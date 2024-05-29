package customer

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	_customerDtos "github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerUsecase interface {
	OnCreateCustomerAccount(c *fiber.Ctx, ctx context.Context, customerDatReq *_customerEntities.CustomerRegisterReq) (*_customerDtos.CustomerAccountFileRes, *_accDtos.UsersRegisteredRes, int, error)
	OnGetCustomerById(c *fiber.Ctx, ctx context.Context, Id *int64) (*_customerDtos.CustomerAccountFileRes, int, error)
}
