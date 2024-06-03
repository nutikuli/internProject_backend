package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	_customerDtos "github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type customerUsecase struct {
	customerRepo customer.CustomerRepository
	accUsecase   account.AccountUsecase
}

func NewCustomerUsecase(customerRepo customer.CustomerRepository, accUsecase account.AccountUsecase) customer.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
		accUsecase:   accUsecase,
	}
}

func (s *customerUsecase) OnCreateCustomerAccount(c *fiber.Ctx, ctx context.Context, customerDatReq *_customerEntities.CustomerRegisterReq) (*_customerDtos.CustomerAccountFileRes, *_accDtos.UserToken, int, error) {

	accRegister, usrCred, status, errOnRegister := s.accUsecase.Register(ctx, customerDatReq)
	if errOnRegister != nil {
		return nil, nil, status, errOnRegister
	}

	customerDatReq.Password = usrCred.Password

	newCustomerId, err := s.customerRepo.CreateCustomerAccount(ctx, customerDatReq)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	customerRes, errOnGetCustomer := s.customerRepo.GetCustomerById(ctx, *newCustomerId)
	if errOnGetCustomer != nil {
		return nil, nil, http.StatusInternalServerError, errOnGetCustomer
	}

	return &_customerDtos.CustomerAccountFileRes{
		CustomerData: customerRes,
	}, accRegister, http.StatusOK, nil

}

func (s *customerUsecase) OnGetCustomerById(ctx context.Context, customerId int64) (*_customerDtos.CustomerAccountFileRes, int, error) {
	customerRes, errOnGetCustomer := s.customerRepo.GetCustomerById(ctx, customerId)
	if errOnGetCustomer != nil {
		return nil, http.StatusInternalServerError, errOnGetCustomer
	}
	return &_customerDtos.CustomerAccountFileRes{
		CustomerData: customerRes,
	}, http.StatusOK, nil
}

func (u *customerUsecase) OnUpdateCustomerById(ctx context.Context, userId int64, req *entities.CustomerUpdateReq) (int, error) {

	err := u.customerRepo.UpdateCustomerById(ctx, userId, req)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update user by ID: %w", err)
	}

	return http.StatusOK, nil
}

func (u *customerUsecase) OnDeletedCustomer(ctx context.Context, Id int64) (int, error) {

	err := u.customerRepo.DeleteCustomerById(ctx, Id)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete user by ID: %w", err)
	}

	return http.StatusOK, nil
}
