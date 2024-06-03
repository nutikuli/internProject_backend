package dtos

import (
	"github.com/nutikuli/internProject_backend/internal/models/product/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type ProductFileUpdateReq struct {
	ProductData *entities.ProductUpdateReq       `json:"product_data"`
	FilesData   []*_fileEntities.FileUploaderReq `json:"files_data"`
}
