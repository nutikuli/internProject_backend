package dtos

import (
	"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerAccountFileRes struct {
	CustomerData *_customerEntities.Customer `json:"customer_data"`
}

type StoreTokenRes struct {
	Customer *CustomerAccountFileRes `json:"customer"`
	Token    *dtos.UserToken         `json:"token"`
}
