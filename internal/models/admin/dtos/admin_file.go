package entities 

import (
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type AdminFileRes struct {
	AdminData *_adminEntities.Admin `json:"admin_data"`
	FilesData []*_fileEntities.File `json:"files_data"`
}