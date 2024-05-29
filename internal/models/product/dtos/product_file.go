package dtos

import (
	_productEntities "github.com/nutikuli/internProject_backend/internal/models/product/entities"
	_filesEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type ProductFileRes struct {
	Files                     []*_filesEntities.File `json:"files_data" form:"files_data" binding:"required"`
	*_productEntities.Product `json:"product_data" form:"product_data" binding:"required"`
}
