package entities

import (
	"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_adminpermissionEntities "github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type AdminPermissionFileRes struct {
	AdminData *_adminEntities.Admin `json:"admin_data"`
	AdminpermissionData *_adminpermissionEntities.Adminpermission `json:"adminpermission_data"`
	FilesData []*_fileEntities.File `json:"files_data"`
	
} 

type AdminTokenFileReqs struct {
	Adminpermission *AdminPermissionFileRes `json:"admin_data"`
	Token *dtos.UserToken   `json:"token"`
}

type AdminPermissionFileReq struct {
	AdminData *_adminEntities.Admin `json:"admin_data"`
	AdminpermissionData *_adminpermissionEntities.AdminPermissionCreatedReq `json:"adminpermission_data"`
	FilesData []*_fileEntities.FileUploaderReq `json:"files_data"`
	
} 

