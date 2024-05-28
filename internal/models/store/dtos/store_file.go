package dtos

import (
	_storeEntities "github.com/nutikuli/internProject_backend/internal/models/store/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type StoreWithFileRes struct {
	StoreData *_storeEntities.Store `json:"store_data"`
	FilesData []*_fileEntities.File `json:"files_data"`
}
