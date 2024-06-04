package admin

import (
	"context"

	_accountEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/admin/entities"
)

type AdminRepository interface {
	GetAccountAdmins(ctx context.Context) ([]*entities.Admin, error)
	GetAccountAdminById(ctx context.Context, id int64) (*entities.Admin, error)
	//INSERT ADMIN
	CreateAdmin(ctx context.Context, admindata *entities.AdminRegisterReq) (*int64, error)
	// UPDATE ADMIN
	UpdateAdminById(ctx context.Context, Id int64, admindata *entities.AdminUpdateReq) error
	//DELETE ADMIN
	DeleteAdminById(ctx context.Context, Id int64) error
	//UPDATE PASSWORD ADMIN
	UpdateAdminPasswordById(ctx context.Context, admindata *_accountEntities.UpdatePass) error
}
