package entities 

import (
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type AdminFileRes struct {
	files []_fileEntities.File
	_adminEntities.Admin
	AdminData *_adminEntities.Admin `json:"admin_data"`
	FilesData []*_fileEntities.File `json:"files_data"`
}