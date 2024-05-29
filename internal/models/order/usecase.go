package order

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_orderDtos "github.com/nutikuli/internProject_backend/internal/models/order/dtos"
	_orderEntities "github.com/nutikuli/internProject_backend/internal/models/order/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type OrderUsecase interface {
	OnGetOrderById(c *fiber.Ctx, ctx context.Context, Id *int64) (*_orderDtos.OrderWithFileRes, int, error)
	OnCreateOrder(c *fiber.Ctx, ctx context.Context, orderDatReq *_orderEntities.OrderCreate, filesDatReq []*_fileEntities.FileUploaderReq) (*_orderDtos.OrderWithFileRes, int, error)
}
