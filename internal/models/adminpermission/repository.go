package adminpermission

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
)

type AdminPermissionRepository interface {

	//INSERT MENU ADMIN PERMISSION
	CreateAdminPermission(ctx context.Context, adminpermissiondata *entities.AdminPermissionCreatedReq) (*int64, error)

	//GET MENU ADMIN PERMISSION
	GetAdminPermissions(ctx context.Context) (*entities.AdminPermissionCreatedReq, error)

	// GET ADMIN PERMISSION BY ID
	GetAdminpermissiomById(ctx context.Context, id *int64) (*entities.Adminpermission, error) 

	//UPDATE 
	UpdateAdminPermissionById(ctx context.Context, Id int64, adminpermissiondata *entities.AdminPermissionUpdatedReq) error

	//DELETE
	 DeleteAdminPermissionById(ctx context.Context, Id int64) error
}
