package dtos 

import (
	"github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type AdminFileUpdateReq struct {
	AdminData *entities.AdminUpdateReq       `json:"admin_data"`
	FilesData   []*_fileEntities.FileUploaderReq `json:"files_data"`
}