package usercase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	"github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
	
)

type AccountUsecase struct {
	accountRepo account.AccountRepository
	fileRepo    file.FileRepository
}

func NewFileUsecase(accountRepo account.AccountRepository, filesRepo file.FileRepository) account.AccountUsecase {
	return &AccountUsecase{
		accountRepo: accountRepo,
		fileRepo:    filesRepo,
	}
}


func (a *AccountUsecase) AccountCustomerfile(ctx context.Context) ([]*dtos.CustomerAccountFileRes, int, error) {

	customers, err := a.accountRepo.GetAccountCustomers(ctx)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Debug("+*+*+*+**+++++*++*+++**+*+*+**++**+*+*+**", customers)

	var res []*dtos.CustomerAccountFileRes

	for _, customer := range customers {
		fileEntity := &_fileEntities.FileEntityReq{
			EntityType: "CUSTOMER",
			EntityId:   *&customer.Id,
		}

		filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
		if errOnGetFiles != nil {
			return nil, http.StatusInternalServerError, errOnGetFiles
		}

		cFile := &dtos.CustomerAccountFileRes{
			Customer: *customer,
			Files:    filesRes,
		}

		res = append(res, cFile)
	}
	return res, http.StatusOK, nil


}

func (a *AccountUsecase) AccountStorefile(ctx context.Context) ([]*_storeDtos.StoreWithFileRes, int, error) {

	stores, err := a.accountRepo.GetAccountStores(ctx)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Debug("store : ", stores)

	var res []*_storeDtos.StoreWithFileRes

	for _, store := range stores {
		fileEntity := &_fileEntities.FileEntityReq{
			EntityType: "STORE",
			EntityId:   *&store.Id,
		}

		filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
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

func (a *AccountUsecase) AccountAdminfile(ctx context.Context) ([]*_adminDtos.AdminFileRes, int, error) {

	admins, err := a.accountRepo.GetAccountAdmins(ctx)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Debug("store : ", admins)

	var res []*_adminDtos.AdminFileRes

	for _, admin := range admins {
		fileEntity := &_fileEntities.FileEntityReq{
			EntityType: "ADMIN",
			EntityId:   *&admin.Id,
		}

		filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
		if errOnGetFiles != nil {
			return nil, http.StatusInternalServerError, errOnGetFiles
		}

		aFile := &_adminDtos.AdminFileRes{
			AdminData: admin,
			FilesData: filesRes,
		}

		res = append(res, aFile)
	}
	return res, http.StatusOK, nil


}

