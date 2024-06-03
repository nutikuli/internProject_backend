package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/bank"
	_bankDtos "github.com/nutikuli/internProject_backend/internal/models/bank/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/order"
	order_product "github.com/nutikuli/internProject_backend/internal/models/order-product"
	_orderProductEntities "github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
	_orderDtos "github.com/nutikuli/internProject_backend/internal/models/order/dtos"
	_orderEntities "github.com/nutikuli/internProject_backend/internal/models/order/entities"
	"github.com/nutikuli/internProject_backend/internal/models/product"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type orderUsecase struct {
	orderRepo        order.OrderRepository
	fileRepo         file.FileRepository
	bankUse          bank.BankUseCase
	productUse       product.ProductUsecase
	ordersProductUse order_product.OrderProductUsecase
}

func NewOrderUsecase(
	orderRepo order.OrderRepository,
	fileRepo file.FileRepository,
	bankUse bank.BankUseCase,
	productUse product.ProductUsecase,
	ordersProductUse order_product.OrderProductUsecase) order.OrderUsecase {
	return &orderUsecase{
		orderRepo:        orderRepo,
		fileRepo:         fileRepo,
		bankUse:          bankUse,
		productUse:       productUse,
		ordersProductUse: ordersProductUse,
	}
}

func (s *orderUsecase) OnCreateOrder(c *fiber.Ctx, ctx context.Context, bankId int64, orderDatReq *_orderEntities.OrderCreate, filesDatReq []*_fileEntities.FileUploaderReq, orderProductsReq []*_orderProductEntities.OrderProductCreateReq) (*_orderDtos.OrderBanksFilesRes, int, error) {
	newOrderId, err := s.orderRepo.CreateOrder(ctx, orderDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ORDER",
		EntityId:   *newOrderId,
	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "ORDER",
			OrderId:    newOrderId,
		}

		_, fUrl, status, errOnCreatedFile := file.EncodeBase64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, status, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		_, errOnInsertFile := s.fileRepo.CreateFileByEntityAndId(ctx, fDatReq, fileEntity)
		if errOnInsertFile != nil {
			return nil, http.StatusInternalServerError, errOnInsertFile
		}
	}

	_, status, err := s.ordersProductUse.OnCreateOrderProducts(ctx, *newOrderId, orderProductsReq)
	if err != nil {
		return nil, status, err
	}

	prodRes, status, err := s.productUse.OnGetProductsByOrderId(ctx, *newOrderId)
	if err != nil {
		return nil, status, err
	}

	bankRes, status, err := s.bankUse.OnGetBankByBankId(ctx, bankId)

	filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	orderRes, errOnGetOrder := s.orderRepo.GetOrderById(ctx, newOrderId)
	if errOnGetOrder != nil {
		return nil, http.StatusInternalServerError, errOnGetOrder
	}

	return &_orderDtos.OrderBanksFilesRes{
		OrderData:         orderRes,
		FilesData:         filesRes,
		BanksData:         []*_bankDtos.BankFileRes{bankRes},
		OrdersProductData: prodRes,
	}, http.StatusOK, nil
}

func (s *orderUsecase) OnGetOrderById(ctx context.Context, req *_orderEntities.StoreAndOrderIdReq) (*_orderDtos.OrderBanksFilesRes, int, error) {
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "Order",
		EntityId:   req.OrderId,
	}

	filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	orderRes, errOnGetOrder := s.orderRepo.GetOrderById(ctx, &req.OrderId)
	if errOnGetOrder != nil {
		return nil, http.StatusInternalServerError, errOnGetOrder
	}

	banksRes, status, errOnGetBanks := s.bankUse.OnGetBanksByStoreId(ctx, req.StoreId)
	if errOnGetBanks != nil {
		return nil, status, errOnGetBanks
	}

	ordersProductRes, status, errOnGetOrdersProduct := s.productUse.OnGetProductsByOrderId(ctx, req.OrderId)
	if errOnGetOrdersProduct != nil {
		return nil, status, errOnGetOrdersProduct
	}

	return &_orderDtos.OrderBanksFilesRes{
		OrderData:         orderRes,
		FilesData:         filesRes,
		BanksData:         banksRes,
		OrdersProductData: ordersProductRes,
	}, http.StatusOK, nil
}

func (s *orderUsecase) OnGetOrdersByStoreId(ctx context.Context, storeId *int64) ([]*_orderDtos.OrderBanksFilesRes, int, error) {
	orderRes, errOnGetOrders := s.orderRepo.GetOrdersByStoreId(ctx, storeId)
	if errOnGetOrders != nil {
		return nil, http.StatusInternalServerError, errOnGetOrders
	}

	var orderWithFileRes = make([]*_orderDtos.OrderBanksFilesRes, 0)

	for _, o := range orderRes {
		fileEntity := &_fileEntities.FileEntityReq{
			EntityType: "ORDER",
			EntityId:   o.Id,
		}

		filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
		if errOnGetFiles != nil {
			return nil, http.StatusInternalServerError, errOnGetFiles
		}

		banksRes, status, errOnGetBanks := s.bankUse.OnGetBanksByStoreId(ctx, o.StoreId)
		if errOnGetBanks != nil {
			return nil, status, errOnGetBanks
		}

		ordersProductRes, status, errOnGetOrdersProduct := s.productUse.OnGetProductsByOrderId(ctx, o.Id)
		if errOnGetOrdersProduct != nil {
			return nil, status, errOnGetOrdersProduct
		}

		res := &_orderDtos.OrderBanksFilesRes{
			OrderData:         o,
			FilesData:         filesRes,
			BanksData:         banksRes,
			OrdersProductData: ordersProductRes,
		}

		orderWithFileRes = append(orderWithFileRes, res)
	}

	return orderWithFileRes, http.StatusOK, nil
}

func (s *orderUsecase) OnGetOrdersByCustomerId(ctx context.Context, customerId *int64) ([]*_orderDtos.OrderBanksFilesRes, int, error) {
	orderRes, errOnGetOrders := s.orderRepo.GetOrdersByCustomerId(ctx, customerId)
	if errOnGetOrders != nil {
		return nil, http.StatusInternalServerError, errOnGetOrders
	}

	var orderWithFileRes = make([]*_orderDtos.OrderBanksFilesRes, 0)

	for _, o := range orderRes {
		fileEntity := &_fileEntities.FileEntityReq{
			EntityType: "ORDER",
			EntityId:   o.Id,
		}

		filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
		if errOnGetFiles != nil {
			return nil, http.StatusInternalServerError, errOnGetFiles
		}

		banksRes, status, errOnGetBanks := s.bankUse.OnGetBanksByStoreId(ctx, o.StoreId)
		if errOnGetBanks != nil {
			return nil, status, errOnGetBanks
		}

		ordersProductRes, status, errOnGetOrdersProduct := s.productUse.OnGetProductsByOrderId(ctx, o.Id)
		if errOnGetOrdersProduct != nil {
			return nil, status, errOnGetOrdersProduct
		}

		res := &_orderDtos.OrderBanksFilesRes{
			OrderData:         o,
			FilesData:         filesRes,
			BanksData:         banksRes,
			OrdersProductData: ordersProductRes,
		}

		orderWithFileRes = append(orderWithFileRes, res)
	}

	return orderWithFileRes, http.StatusOK, nil
}

// OnUpdateOrderStatus implements order.OrderUsecase.
func (s *orderUsecase) OnUpdateOrderStatus(ctx context.Context, req *_orderEntities.OrderStateReq) (int, error) {
	err := s.orderRepo.UpdateOrderStatus(ctx, req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// OnUpdateOrderTransportDetail implements order.OrderUsecase.
func (s *orderUsecase) OnUpdateOrderTransportDetail(ctx context.Context, req *_orderEntities.OrderTransportDetailReq) (int, error) {
	err := s.orderRepo.UpdateOrderTransportDetail(ctx, req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
