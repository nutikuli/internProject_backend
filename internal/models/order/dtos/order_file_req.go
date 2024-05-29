package dtos

import (
	_orderEntities "github.com/nutikuli/internProject_backend/internal/models/order/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type OrderFileReq struct {
	OrderData *_orderEntities.Order            `db:"order_data" form:"order_data"`
	FilesData []*_fileEntities.FileUploaderReq `json:"files_data" form:"files_data" binding:"required"`
}
