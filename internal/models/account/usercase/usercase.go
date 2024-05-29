package usercase

import (
	"context"
	"fmt"
	"net/http"
	
	"github.com/gofiber/fiber/v2/log"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_storeDtos "github.com/nutikuli/internProject_backend/internal/models/store/dtos"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
	"golang.org/x/crypto/bcrypt"
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

func (a *AccountUsecase) Login(ctx context.Context, req *entities.UsersLogin) (*_accDtos.UserToken, int, error) {

	user, err := a.accountRepo.FindUserAsPassport(ctx, req.Email)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusInternalServerError, err
	}

	userToken, err := a.accountRepo.SignUsersAccessToken(&entities.UserSignToken{
		Id:       user.Id,
		Email: req.Email,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return userToken, http.StatusOK, nil
}

func (a *AccountUsecase) Register(ctx context.Context, req *entities.UserCreatedReq) (*_accDtos.UsersRegisteredRes, int, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	cred := entities.UserCreatedReq{
		Password: string(hashedPassword),
		Email: req.Email,
		
	}
	req.Password = cred.Password
	log.Info("req", req)
	user, err := a.accountRepo.CreateUser(ctx, req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	log.Info("res", user)

	userToken, err := a.accountRepo.SignUsersAccessToken(&entities.UserSignToken{
		Id:       *user,
		Email: req.Email,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	// //Receiver email address.
	//   to := req.Email // <-------------- (3) แก้ไขอีเมลของผู้รับ หากใส่หลายเมล จะไปอยู่ที่ cc

	// //Message.
	//   message := []byte("Register Success")

	//   // Authentication.
	//   auth := smtp.PlainAuth("",viper.GetString("emailFrom"), viper.GetString("passwordMail"), viper.GetString("smtpHost"))
	//   log.Debug("+++++++++++++++**********-----------------",auth)
	//   // Sending email.
	//   err = smtp.SendMail(viper.GetString("smtpHost")+":"+viper.GetString("smtpPort"), auth, viper.GetString("emailFrom"),[]string{to}, message)

	res := &_accDtos.UsersRegisteredRes{
		AccessToken: userToken.AccessToken,
		CreatedAt:   userToken.IssuedAt,
		ExpiredAt:   userToken.ExpiresIn,
	}

	return res, http.StatusOK, nil

}