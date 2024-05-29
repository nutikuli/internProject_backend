package customer

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_customerDtos "github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerUsecase interface {
	OnCreateCustomerAccount(c *fiber.Ctx, ctx context.Context, customerDatReq *_customerEntities.CustomerRegister) (*_customerDtos.CustomerAccountFileRes, int, error)
	OnGetCustomerById(c *fiber.Ctx, ctx context.Context, Id *int64) (*_customerDtos.CustomerAccountFileRes, int, error)
}
