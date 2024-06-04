package dtos

import (
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type CustomerRes struct {
	CustomerData *_customerEntities.Customer `json:"customer_data"`
}
