package dtos 

import (
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_adminpermissionEntities "github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
	
	
)

type AdminFileRes struct {
	AdminData *_adminEntities.Admin `json:"admin_data"`
	AdminpermissionData []*_adminpermissionEntities.Adminpermission `json:"adminpermission_data"`
	FilesData []*_fileEntities.File `json:"files_data"`

} 

