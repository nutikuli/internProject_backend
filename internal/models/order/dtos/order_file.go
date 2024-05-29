package dtos

import (
	_orderEntities "github.com/nutikuli/internProject_backend/internal/models/order/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type OrderWithFileRes struct {
	OrderData *_orderEntities.Order `json:"order_data"`
	FilesData []*_fileEntities.File `json:"files_data"`
}
