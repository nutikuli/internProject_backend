package dtos

import (
	_storeEntities "github.com/nutikuli/internProject_backend/internal/models/store/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type StoreUpdateReq struct {
	StoreData *_storeEntities.StoreUpdatedReq  `json:"store_data" form:"store_data" binding:"required"`
	FilesData []*_fileEntities.FileUploaderReq `json:"files_data" form:"files_data"`
}
