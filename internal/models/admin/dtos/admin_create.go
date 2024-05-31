package dtos 

import (
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities" 
	"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	 
	
)

type AdminCreateFileReq struct {
	AdminData *_adminEntities.AdminRegisterReq `json:"admin_register_data" form:"admin_register_data" binding:"required"`
	FilesData []*_fileEntities.FileUploaderReq `json:"files_data" form:"files_data" binding:"required"`

} 

type AdminTokenFileReqs struct {
	Admin *AdminFileRes `json:"admin_data"`
	Token *dtos.UserToken   `json:"token"`
}


