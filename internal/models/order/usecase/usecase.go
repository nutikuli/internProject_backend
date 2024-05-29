package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/order"
	_orderDtos "github.com/nutikuli/internProject_backend/internal/models/order/dtos"
	_orderEntities "github.com/nutikuli/internProject_backend/internal/models/order/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type orderUsecase struct {
	orderRepo order.OrderRepository
	fileRepo  file.FileRepository
}

func NewStoreUsecase(orderRepo order.OrderRepository, fileRepo file.FileRepository) order.OrderUsecase {
	return &orderUsecase{
		orderRepo: orderRepo,
		fileRepo:  fileRepo,
	}
}

func (s *orderUsecase) OnCreateOrder(c *fiber.Ctx, ctx context.Context, orderDatReq *_orderEntities.OrderCreate, filesDatReq []*_fileEntities.FileUploaderReq) (*_orderDtos.OrderWithFileRes, int, error) {
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
			EntityId:   *newOrderId,
		}

		_, fUrl, errOnCreatedFile := file.Base64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, http.StatusConflict, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		_, errOnInsertFile := s.fileRepo.CreateFileByEntityAndId(ctx, fDatReq, fileEntity)
		if errOnInsertFile != nil {
			return nil, http.StatusInternalServerError, errOnInsertFile
		}
	}

	filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	orderRes, errOnGetOrder := s.orderRepo.GetOrderById(ctx, newOrderId)
	if errOnGetOrder != nil {
		return nil, http.StatusInternalServerError, errOnGetOrder
	}

	return &_orderDtos.OrderWithFileRes{
		OrderData: orderRes,
		FilesData: filesRes,
	}, http.StatusOK, nil
}

func (s *orderUsecase) OnGetOrderById(c *fiber.Ctx, ctx context.Context, Id *int64) (*_orderDtos.OrderWithFileRes, int, error) {
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "Order",
		EntityId:   *Id,
	}

	filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	orderRes, errOnGetOrder := s.orderRepo.GetOrderById(ctx, Id)
	if errOnGetOrder != nil {
		return nil, http.StatusInternalServerError, errOnGetOrder
	}

	return &_orderDtos.OrderWithFileRes{
		OrderData: orderRes,
		FilesData: filesRes,
	}, http.StatusOK, nil
}
