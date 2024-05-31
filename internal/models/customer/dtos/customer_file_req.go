package dtos

import (
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerFileReq struct {
	CustomerRegisterData *_customerEntities.CustomerRegisterReq `json:"customer_register_data" form:"customer_register_data" binding:"required"`
}
