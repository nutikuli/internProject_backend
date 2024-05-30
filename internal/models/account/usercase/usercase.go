package usercase

import (
	"context"
	"crypto/rand"
	"os"
	"sync"
	"time"

	"errors"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
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
	accountRepo     account.AccountRepository
	fileRepo        file.FileRepository
	adminUse        admin.AdminUseCase
	adminUseRepo    admin.AdminRepository
	customerUseRepo customer.CustomerRepository
	customerUse     customer.CustomerUsecase
	storeUse        store.StoreUsecase
	storeUseRepo    store.StoreRepository
}

func NewFileUsecase(
	accountRepo account.AccountRepository,
	filesRepo file.FileRepository,
	adminUse admin.AdminUseCase,
	customerUse customer.CustomerUsecase,
	adminUseRepo admin.AdminRepository,
	customerUseRepo customer.CustomerRepository,
	storeUseRepo store.StoreRepository,
	storeUse store.StoreUsecase) account.AccountUsecase {
	return &AccountUsecase{
		accountRepo:     accountRepo,
		fileRepo:        filesRepo,
		adminUse:        adminUse,
		customerUse:     customerUse,
		storeUse:        storeUse,
		adminUseRepo:    adminUseRepo,
		customerUseRepo: customerUseRepo,
		storeUseRepo:    storeUseRepo,
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

func (a *AccountUsecase) Login(ctx context.Context, req *entities.UsersCredential) (*_accDtos.UserToken, interface{}, int, error) {

	user, err := a.accountRepo.FindUserAsPassport(ctx, req.Email)

	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(err.Error())
		return nil, nil, http.StatusInternalServerError, err
	}

	userToken, err := a.accountRepo.SignUsersAccessToken(&entities.UserSignToken{
		Role:  user.Role,
		Email: req.Email,
	})
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	var roleAccount interface{}
	switch user.Role {
	case "CUSTOMER":
		acc, status, err := a.customerUse.OnGetCustomerById(ctx, &user.Id)
		if err != nil {
			return nil, nil, status, err
		}
		roleAccount = acc
	case "STORE":
		acc, status, err := a.storeUse.OnGetStoreById(ctx, &user.Id)
		if err != nil {
			return nil, nil, status, err
		}
		roleAccount = acc
	case "ADMIN":
		acc, status, err := a.adminUse.OnGetAdminById(ctx, &user.Id)
		if err != nil {
			return nil, nil, status, err
		}
		roleAccount = acc
	default:
		return nil, nil, http.StatusInternalServerError, errors.New("Can't query the Account Table, Invalid role")
	}

	return userToken, roleAccount, http.StatusOK, nil
}

func (a *AccountUsecase) Register(ctx context.Context, req entities.AccountCredentialGetter) (*_accDtos.UserToken, *entities.UsersCredential, int, error) {

	if req.GetEmail() == nil || req.GetPassword() == nil || req.GetRole() == nil {
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
		Role:  *req.GetRole(),
		Email: *req.GetEmail(),
	})

	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	//Receiver email address.
	to := req.GetEmail() // <-------------- (3) แก้ไขอีเมลของผู้รับ หากใส่หลายเมล จะไปอยู่ที่ cc

	//Message.
	message := []byte("Register Success")

	// Authentication.
	auth := smtp.PlainAuth("", os.Getenv("emailFrom"), os.Getenv("passwordMail"), os.Getenv("smtpHost"))
	log.Debug("+++++++++++++++**********-----------------", auth)
	// Sending email.
	err = smtp.SendMail(os.Getenv("smtpHost")+":"+os.Getenv("smtpPort"), auth, os.Getenv("emailFrom"), []string{*to}, message)

	return userToken, cred, http.StatusOK, nil

}

func (a *AccountUsecase) CheckOTP(c *fiber.Ctx, ctx context.Context, req *entities.UsersCredential) (*_accDtos.OTPres, int, error) {
	user, err := a.accountRepo.FindUserAsPassport(ctx, req.Email)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	log.Debug(user)

	const otpLength = 6
	const charset = "0123456789"
	otp := make([]byte, otpLength)
	_, err = rand.Read(otp)

	for i := 0; i < otpLength; i++ {
		otp[i] = charset[otp[i]%byte(len(charset))]
	}

	// เก็บ OTP ในหน่วยความจำพร้อมกำหนดเวลาหมดอายุ
	otpStore.Lock()
	otpStore.store[req.Email] = entities.OTPDetails{OTP: string(otp), CreatedAt: time.Now()}
	otpStore.Unlock()

	// ตั้งเวลาให้ OTP หมดอายุ
	go func() {
		time.Sleep(10 * time.Minute)
		otpStore.Lock()
		delete(otpStore.store, req.Email)
		otpStore.Unlock()
	}()

	to := req.Email // <-------------- (3) แก้ไขอีเมลของผู้รับ หากใส่หลายเมล จะไปอยู่ที่ cc

	//Message.
	message := []byte("OTP : " + string(otp))

	// Authentication.
	auth := smtp.PlainAuth("", os.Getenv("emailFrom"), os.Getenv("passwordMail"), os.Getenv("smtpHost"))
	log.Debug("+++++++++++++++**********-----------------", auth)
	// Sending email.
	err = smtp.SendMail(os.Getenv("smtpHost")+":"+os.Getenv("smtpPort"), auth, os.Getenv("emailFrom"), []string{to}, message)
	OTPres := &_accDtos.OTPres{
		OTP:       string(otp),
		Email:     req.Email,
		CreatedAt: time.Now(),
	}
	return OTPres, http.StatusOK, err
}

var otpStore = struct {
	sync.RWMutex
	store map[string]entities.OTPDetails
}{store: make(map[string]entities.OTPDetails)}

func (a *AccountUsecase) ResetPassword(ctx context.Context, req *entities.UsersCredential) (*entities.UpdatePass, interface{}, int, error) {
	user, err := a.accountRepo.FindUserAsPassport(ctx, req.Email)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	log.Debug(user.Email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	repassRes := &entities.UpdatePass{
		Id:       user.Id,
		Password: string(hashedPassword),
		Role:     user.Role,
	}

	var roleAccount interface{}
	switch user.Role {
	case "CUSTOMER":
		err := a.customerUseRepo.UpdateCustomerPasswordById(ctx, repassRes)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
	case "STORE":
		err := a.storeUseRepo.UpdateStoreAccountPassword(ctx, *repassRes)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
	case "ADMIN":
		err := a.adminUseRepo.UpdateAdminPasswordById(ctx, repassRes)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
	default:
		return nil, nil, http.StatusInternalServerError, errors.New("Can't query the Account Table, Invalid role")
	}

	return repassRes, roleAccount, http.StatusOK, err
}
