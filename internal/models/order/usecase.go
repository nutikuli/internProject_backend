package order

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_orderProductEntities "github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
	_orderDtos "github.com/nutikuli/internProject_backend/internal/models/order/dtos"
	_orderEntities "github.com/nutikuli/internProject_backend/internal/models/order/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type OrderUsecase interface {
	OnCreateOrder(c *fiber.Ctx, ctx context.Context, orderDatReq *_orderEntities.OrderCreate, filesDatReq []*_fileEntities.FileUploaderReq, orderProductsReq []*_orderProductEntities.OrderProductCreateReq) (*_orderDtos.OrderBanksFilesRes, int, error)
	OnGetOrderById(ctx context.Context, req *_orderEntities.StoreAndOrderIdReq) (*_orderDtos.OrderBanksFilesRes, int, error)
	OnGetOrdersByCustomerId(ctx context.Context, customerId *int64) ([]*_orderDtos.OrderBanksFilesRes, int, error)
	OnGetOrdersByStoreId(ctx context.Context, storeId *int64) ([]*_orderDtos.OrderBanksFilesRes, int, error)
	OnUpdateOrderStatus(ctx context.Context, orderId int64, req *_orderEntities.OrderStateReq) (int, error)
	OnUpdateOrderTransportDetail(ctx context.Context, orderId int64, req *_orderEntities.OrderTransportDetailReq) (int, error)
}
