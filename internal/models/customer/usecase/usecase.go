package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	_customerDtos "github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type customerUsecase struct {
	customerRepo customer.CustomerRepository
	accUsecase   account.AccountUsecase
}

func NewStoreUsecase(customerRepo customer.CustomerRepository, accUsecase account.AccountUsecase) customer.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
		accUsecase:   accUsecase,
	}
}

func (s *customerUsecase) OnCreateCustomerAccount(c *fiber.Ctx, ctx context.Context, customerDatReq *_customerEntities.CustomerRegisterReq) (*_customerDtos.CustomerAccountFileRes, *_accDtos.UsersRegisteredRes, int, error) {

	accRegister, usrCred, status, errOnRegister := s.accUsecase.Register(ctx, customerDatReq)
	if errOnRegister != nil {
		return nil, nil, status, errOnRegister
	}

	customerDatReq.Password = usrCred.Password

	newCustomerId, err := s.customerRepo.CreateCustomerAccount(ctx, customerDatReq)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	customerRes, errOnGetCustomer := s.customerRepo.GetCustomerById(ctx, newCustomerId)
	if errOnGetCustomer != nil {
		return nil, nil, http.StatusInternalServerError, errOnGetCustomer
	}

	return &_customerDtos.CustomerAccountFileRes{
		Customer: *customerRes,
	}, accRegister, http.StatusOK, nil

}

func (s *customerUsecase) OnGetCustomerById(c *fiber.Ctx, ctx context.Context, Id *int64) (*_customerDtos.CustomerAccountFileRes, int, error) {

	customerRes, errOnGetCustomer := s.customerRepo.GetCustomerById(ctx, Id)
	if errOnGetCustomer != nil {
		return nil, http.StatusInternalServerError, errOnGetCustomer
	}
	return &_customerDtos.CustomerAccountFileRes{
		Customer: *customerRes,
	}, http.StatusOK, nil
}
