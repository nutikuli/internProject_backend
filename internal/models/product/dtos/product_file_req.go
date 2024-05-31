package dtos

import (
	"github.com/nutikuli/internProject_backend/internal/models/product/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type ProductFileReq struct {
	ProductData *entities.ProductCreateReq       `json:"product_data"`
	FileData    []*_fileEntities.FileUploaderReq `json:"file_data"`
}
