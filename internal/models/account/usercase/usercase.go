package usercase

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/admin"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	"github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/store"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
	"golang.org/x/crypto/bcrypt"
)

type AccountUsecase struct {
	accountRepo account.AccountRepository
	fileRepo    file.FileRepository
	adminUse    admin.AdminUseCase
	customerUse customer.CustomerUsecase
	storeUse    store.StoreUsecase
}

func NewFileUsecase(accountRepo account.AccountRepository,
	filesRepo file.FileRepository,
	adminUse admin.AdminUseCase,
	customerUse customer.CustomerUsecase,
	storeUse store.StoreUsecase) account.AccountUsecase {
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

		c := &dtos.CustomerAccountFileRes{
			Customer: *customer,
		}

		res = append(res, c)
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
			EntityId:   store.Id,
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
			EntityId:   admin.Id,
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

func (a *AccountUsecase) Login(ctx context.Context, req *entities.UsersCredential) (*_accDtos.UserToken, int, error) {

	user, err := a.accountRepo.FindUserAsPassport(ctx, req.Email)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusInternalServerError, err
	}

	userToken, err := a.accountRepo.SignUsersAccessToken(&entities.UserSignToken{
		Id:    user.Id,
		Email: req.Email,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	switch user.Role {
		case "CUSTOMER":
		_, _, err := a.customerUse.
	}

	return userToken, http.StatusOK, nil
}

func (a *AccountUsecase) Register(ctx context.Context, req entities.AccountCredentialGetter) (*_accDtos.UsersRegisteredRes, *entities.UsersCredential, int, error) {

	if req.GetEmail() == nil || req.GetPassword() == nil || req.GetId() == nil {
		return nil, nil, http.StatusBadRequest, errors.New("Invalid request, not found AccountCredential when registering Account.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	cred := &entities.UsersCredential{
		Password: string(hashedPassword),
		Email:    *req.GetEmail(),
	}

	userToken, err := a.accountRepo.SignUsersAccessToken(&entities.UserSignToken{
		Id:    *req.GetId(),
		Email: *req.GetEmail(),
	})

	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	res := &_accDtos.UsersRegisteredRes{
		AccessToken: userToken.AccessToken,
		CreatedAt:   userToken.IssuedAt,
		ExpiredAt:   userToken.ExpiresIn,
	}

	return res, cred, http.StatusOK, nil

}
