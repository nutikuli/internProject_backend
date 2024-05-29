package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/store"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
	_storeEntities "github.com/nutikuli/internProject_backend/internal/models/store/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type storeUsecase struct {
	storeRepo store.StoreRepository
	fileRepo  file.FileRepository
}

func NewStoreUsecase(storeRepo store.StoreRepository, fileRepo file.FileRepository) store.StoreUsecase {
	return &storeUsecase{
		storeRepo: storeRepo,
		fileRepo:  fileRepo,
	}
}

func (s *storeUsecase) OnCreateStoreAccount(c *fiber.Ctx, ctx context.Context, storeDatReq *_storeEntities.StoreRegisterReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_storeDtos.StoreWithFileRes, int, error) {

	newStoreId, err := s.storeRepo.CreateStoreAccount(ctx, *storeDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "STORE",
		EntityId:   *newStoreId,
	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "STORE",
			EntityId:   *newStoreId,
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

	storeRes, errOnGetStore := s.storeRepo.GetStoreById(ctx, newStoreId)
	if errOnGetStore != nil {
		return nil, http.StatusInternalServerError, errOnGetStore
	}

	return &_storeDtos.StoreWithFileRes{
		StoreData: storeRes,
		FilesData: filesRes,
	}, http.StatusOK, nil

}

func (s *storeUsecase) OnGetStoreById(c *fiber.Ctx, ctx context.Context, storeId *int64) (*_storeDtos.StoreWithFileRes, int, error) {
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "STORE",
		EntityId:   *storeId,
	}

	filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	storeRes, errOnGetStore := s.storeRepo.GetStoreById(ctx, storeId)
	if errOnGetStore != nil {
		return nil, http.StatusInternalServerError, errOnGetStore
	}

	return &_storeDtos.StoreWithFileRes{
		StoreData: storeRes,
		FilesData: filesRes,
	}, http.StatusOK, nil
}
