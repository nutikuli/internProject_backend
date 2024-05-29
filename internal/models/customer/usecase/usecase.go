package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	_customerDtos "github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type customerUsecase struct {
	customerRepo customer.CustomerRepository
}

func NewStoreUsecase(customerRepo customer.CustomerRepository) customer.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (s *customerUsecase) OnCreateCustomerAccount(c *fiber.Ctx, ctx context.Context, customerDatReq *_customerEntities.CustomerRegister) (*_customerDtos.CustomerAccountFileRes, int, error) {

	newCustomerId, err := s.customerRepo.CreateCustomerAccount(ctx, customerDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	customerRes, errOnGetCustomer := s.customerRepo.GetCustomerById(ctx, newCustomerId)
	if errOnGetCustomer != nil {
		return nil, http.StatusInternalServerError, errOnGetCustomer
	}

	return &_customerDtos.CustomerAccountFileRes{
		Customer: *customerRes,
	}, http.StatusOK, nil

}

func (s *customerUsecase) OnGetStoreById(c *fiber.Ctx, ctx context.Context, Id *int64) (*_customerDtos.CustomerAccountFileRes, int, error) {

	customerRes, errOnGetCustomer := s.customerRepo.GetCustomerById(ctx, Id)
	if errOnGetCustomer != nil {
		return nil, http.StatusInternalServerError, errOnGetCustomer
	}
	return &_customerDtos.CustomerAccountFileRes{
		Customer: *customerRes,
	}, http.StatusOK, nil
}
