package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

// OnDeleteStoreById implements store.StoreUsecase.
func (s *storeUsecase) OnDeleteStoreById(ctx context.Context, storeId int64) (int, error) {
	err := s.storeRepo.DeleteStoreById(ctx, storeId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
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
			AccountId:  newStoreId,
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

func (s *storeUsecase) OnUpdateStoreById(c *fiber.Ctx, ctx context.Context, storeId int64, storeDatReq *_storeEntities.StoreUpdatedReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_storeDtos.StoreWithFileRes, int, error) {
	err := s.storeRepo.UpdateStoreById(ctx, storeId, *storeDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err

	}

	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ACCOUNT",
		EntityId:   storeId,
	}

	oldFilesProd, err := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if err != nil {
		log.Debug("error get file ", err)
		return nil, http.StatusInternalServerError, err
	}

	for _, f := range oldFilesProd {
		errOnDeleteFile := s.fileRepo.DeleteFileByIdAndEntity(ctx, f.Id, fileEntity)
		if errOnDeleteFile != nil {
			return nil, http.StatusInternalServerError, errOnDeleteFile
		}

	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "ACCOUNT",
			AccountId:  &storeId,
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

// OnGetStoreAccounts implements store.StoreUsecase.
func (s *storeUsecase) OnGetStoreAccounts(ctx context.Context) ([]*_storeDtos.StoreWithFileRes, int, error) {

	stores, err := s.storeRepo.GetStoreAccounts(ctx)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Debug("store : ", stores)

	var res []*_storeDtos.StoreWithFileRes

	for _, store := range stores {
		fileEntity := &_fileEntities.FileEntityReq{
			EntityType: "ACCOUNT",
			EntityId:   store.Id,
		}

		filesRes, errOnGetFiles := s.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
		if errOnGetFiles != nil {
			return nil, http.StatusInternalServerError, errOnGetFiles
		}

		sFile := &_storeDtos.StoreWithFileRes{
			StoreData: store,
			FilesData: filesRes,
		}

		res = append(res, sFile)
	}
	return res, http.StatusOK, nil
}
