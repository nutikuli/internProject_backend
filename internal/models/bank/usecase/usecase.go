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
			EntityId:   *newBankId,
		}

		_, fUrl, errOnCreatedFile := file.Base64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, http.StatusConflict, errOnCreatedFile
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

	bankRes, errOnGetStore := a.bankRepo.GetBankById(ctx, newBankId)
	if errOnGetStore != nil {
		return nil, http.StatusInternalServerError, errOnGetStore
	}

	return &_bankDtos.BankFileRes{
		BankData:  bankRes,
		FilesData: filesRes,
	}, http.StatusOK, nil

}

func (a *bankUseCase) OnGetBankById(c *fiber.Ctx, ctx context.Context, bankId *int64) (*_bankDtos.BankFileRes, int, error) {
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "BANK",
		EntityId:   *bankId,
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
