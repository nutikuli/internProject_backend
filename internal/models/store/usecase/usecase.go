package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/store"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
	_storeEntities "github.com/nutikuli/internProject_backend/internal/models/store/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type storeUsecase struct {
	storeRepo  store.StoreRepository
	fileRepo   file.FileRepository
	accUsecase account.AccountUsecase
}

func NewStoreUsecase(storeRepo store.StoreRepository, fileRepo file.FileRepository, accUsecase account.AccountUsecase) store.StoreUsecase {
	return &storeUsecase{
		storeRepo:  storeRepo,
		fileRepo:   fileRepo,
		accUsecase: accUsecase,
	}
}

func (s *storeUsecase) OnCreateStoreAccount(c *fiber.Ctx, ctx context.Context, storeDatReq *_storeEntities.StoreRegisterReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_storeDtos.StoreWithFileRes, *_accDtos.UserToken, int, error) {

	accRegister, usrCred, status, errOnRegister := s.accUsecase.Register(ctx, storeDatReq)
	if errOnRegister != nil {
		return nil, nil, status, errOnRegister
	}

	storeDatReq.Password = usrCred.Password

	newStoreId, err := s.storeRepo.CreateStoreAccount(ctx, *storeDatReq)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ACCOUNT",
		EntityId:   *newStoreId,
	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "ACCOUNT",
			AccountId:  *newStoreId,
		}

		_, fUrl, status, errOnCreatedFile := file.EncodeBase64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, nil, status, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		_, errOnInsertFile := s.fileRepo.CreateFileByEntityAndId(ctx, fDatReq, fileEntity)
		if errOnInsertFile != nil {
			return nil, nil, http.StatusInternalServerError, errOnInsertFile
		}
	}
	filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, nil, http.StatusInternalServerError, errOnGetFiles
	}

	storeRes, errOnGetStore := s.storeRepo.GetStoreById(ctx, *newStoreId)
	if errOnGetStore != nil {
		return nil, nil, http.StatusInternalServerError, errOnGetStore
	}

	return &_storeDtos.StoreWithFileRes{
		StoreData: storeRes,
		FilesData: filesRes,
	}, accRegister, http.StatusOK, nil

}

func (s *storeUsecase) OnGetStoreById(ctx context.Context, storeId int64) (*_storeDtos.StoreWithFileRes, int, error) {
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ACCOUNT",
		EntityId:   storeId,
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
