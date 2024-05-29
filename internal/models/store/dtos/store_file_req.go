package dtos

import (
	_storeEntities "github.com/nutikuli/internProject_backend/internal/models/store/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type StoreFileReq struct {
	StoreRegisterData *_storeEntities.StoreRegisterReq `json:"store_register_data" form:"store_register_data" binding:"required"`
	FilesData         []*_fileEntities.FileUploaderReq `json:"files_data" form:"files_data" binding:"required"`
}
