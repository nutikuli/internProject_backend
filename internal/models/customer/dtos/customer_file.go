package dtos

import (
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type CustomerAccountFileRes struct {
	Files []*_fileEntities.File
	_customerEntities.Customer
}
