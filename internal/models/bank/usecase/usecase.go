package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/bank"
	_bankDtos "github.com/nutikuli/internProject_backend/internal/models/bank/dtos"
	_bankEntities "github.com/nutikuli/internProject_backend/internal/models/bank/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type bankUseCase struct {
	bankRepo bank.BankRepository
	fileRepo file.FileRepository
}

func NewBankUsecase(bankRepo bank.BankRepository, fileRepo file.FileRepository) bank.BankUseCase {
	return &bankUseCase{
		bankRepo: bankRepo,
		fileRepo: fileRepo,
	}
}

func (a *bankUseCase) OnCreateBank(c *fiber.Ctx, ctx context.Context, bankDatReq *_bankEntities.BankCreatedReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_bankDtos.BankFileRes, int, error) {

	newBankId, err := a.bankRepo.CreateBank(ctx, bankDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "BANK",
		EntityId:   *newBankId,
	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "BANK",
			BankId:     newBankId,
		}

		_, fUrl, status, errOnCreatedFile := file.EncodeBase64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, status, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		_, errOnInsertFile := a.fileRepo.CreateFileByEntityAndId(ctx, fDatReq, fileEntity)
		if errOnInsertFile != nil {
			return nil, http.StatusInternalServerError, errOnInsertFile
		}
	}

	filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	bankRes, errOnGetStore := a.bankRepo.GetBankById(ctx, *newBankId)
	if errOnGetStore != nil {
		return nil, http.StatusInternalServerError, errOnGetStore
	}

	return &_bankDtos.BankFileRes{
		BankData:  bankRes,
		FilesData: filesRes,
	}, http.StatusOK, nil

}

func (a *bankUseCase) OnGetBanksByStoreId(ctx context.Context, storeId int64) ([]*_bankDtos.BankFileRes, int, error) {

	bankRes, errOnGetStore := a.bankRepo.GetBanksByStoreId(ctx, storeId)
	if errOnGetStore != nil {
		return nil, http.StatusInternalServerError, errOnGetStore
	}

	var bankWithFileRes = make([]*_bankDtos.BankFileRes, 0)

	for _, b := range bankRes {
		fileEntity := &_fileEntities.FileEntityReq{
			EntityType: "BANK",
			EntityId:   b.Id,
		}

		filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
		if errOnGetFiles != nil {
			return nil, http.StatusInternalServerError, errOnGetFiles
		}

		res := &_bankDtos.BankFileRes{
			BankData:  b,
			FilesData: filesRes,
		}

		bankWithFileRes = append(bankWithFileRes, res)
	}

	return bankWithFileRes, http.StatusOK, nil
}

// OnGetBankByBankId implements bank.BankUseCase.
func (a *bankUseCase) OnGetBankByBankId(ctx context.Context, bankId int64) (*_bankDtos.BankFileRes, int, error) {

	bankRes, errOnGetStore := a.bankRepo.GetBankById(ctx, bankId)
	if errOnGetStore != nil {
		return nil, http.StatusInternalServerError, errOnGetStore
	}

	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "BANK",
		EntityId:   bankId,
	}

	filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	return &_bankDtos.BankFileRes{
		BankData:  bankRes,
		FilesData: filesRes,
	}, http.StatusOK, nil
}

// OnDeleteBankById implements bank.BankUseCase.
func (a *bankUseCase) OnDeleteBankById(ctx context.Context, bankId int64) (int, error) {
	err := a.bankRepo.DeleteBankById(ctx, bankId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil

}

// OnUpdateBankById implements bank.BankUseCase.
func (a *bankUseCase) OnUpdateBankById(c *fiber.Ctx, ctx context.Context, bankId int64, bankDatReq *_bankEntities.BankUpdateReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_bankDtos.BankFileRes, int, error) {

	err := a.bankRepo.UpdateBankById(ctx, bankId, bankDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "BANK",
		EntityId:   bankId,
	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "BANK",
			BankId:     &bankId,
		}

		_, fUrl, status, errOnCreatedFile := file.EncodeBase64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, status, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		_, errOnInsertFile := a.fileRepo.CreateFileByEntityAndId(ctx, fDatReq, fileEntity)
		if errOnInsertFile != nil {
			return nil, http.StatusInternalServerError, errOnInsertFile
		}
	}

	filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	bankRes, errOnGetStore := a.bankRepo.GetBankById(ctx, bankId)
	if errOnGetStore != nil {
		return nil, http.StatusInternalServerError, errOnGetStore
	}

	return &_bankDtos.BankFileRes{
		BankData:  bankRes,
		FilesData: filesRes,
	}, http.StatusOK, nil
}
