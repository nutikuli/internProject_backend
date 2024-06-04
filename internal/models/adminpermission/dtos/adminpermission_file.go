package entities

import (
	"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_adminpermissionEntities "github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
)

type AdminPermissionFileRes struct {
	
	AdminpermissionData *_adminpermissionEntities.Adminpermission `json:"adminpermission_data"`
	
	
} 

type AdminTokenFileReqs struct {
	Adminpermission *AdminPermissionFileRes `json:"admin_data"`
	Token *dtos.UserToken   `json:"token"`
}

type AdminPermissionFileReq struct {
	AdminData *_adminEntities.Admin `json:"admin_data"`
	AdminpermissionData *_adminpermissionEntities.AdminPermissionCreatedReq `json:"adminpermission_data"`
	
	
} 

