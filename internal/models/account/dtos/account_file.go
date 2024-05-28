package entities

import (
	_accEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type AccountFileRes struct {
	files []_fileEntities.File
	_accEntities.Account
}
