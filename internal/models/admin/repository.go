package admin


import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/models/admin/entities"
)

type AdminRepository interface {
	GetAccountAdmins(ctx context.Context) (*entities.Admin, error)
	GetAccountAdminById(ctx context.Context, id *int64) (*entities.Admin, error)
	//INSERT ADMIN
	CreateAdmin(ctx context.Context, admindata *entities.AdminCreatedReq) (*int64, error)
	// UPDATE ADMIN
	UpdateAdminById(ctx context.Context, Id int64, admindata *entities.AdminUpdateReq) error
	

}