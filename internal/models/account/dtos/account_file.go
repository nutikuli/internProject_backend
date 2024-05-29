package dtos

import (
	_accCustomerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)



type AccountAdminFileRes struct {
	Files []_fileEntities.File
	_accCustomerEntities.Customer
}
