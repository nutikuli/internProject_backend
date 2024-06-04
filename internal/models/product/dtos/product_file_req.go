package dtos

import (
	"github.com/nutikuli/internProject_backend/internal/models/product/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type ProductFileReq struct {
	ProductData *entities.ProductCreateReq       `json:"product_data" form:"product_data" binding:"required"`
	FilesData   []*_fileEntities.FileUploaderReq `json:"files_data" form:"files_data" binding:"required"`
}
