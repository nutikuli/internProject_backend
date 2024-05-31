package dtos

import (
	"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	_storeEntities "github.com/nutikuli/internProject_backend/internal/models/store/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type StoreWithFileRes struct {
	StoreData *_storeEntities.Store `json:"store_data"`
	FilesData []*_fileEntities.File `json:"files_data"`
}

type StoreTokenFileRes struct {
	Store *StoreWithFileRes `json:"store"`
	Token *dtos.UserToken   `json:"token"`
}
