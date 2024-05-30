package admin

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
)

type AdminUseCase interface {
	OnCreateAdminAccount(c *fiber.Ctx, ctx context.Context, adminDatReq *_adminEntities.AdminRegisterReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_adminDtos.AdminFileRes,*_accDtos.UserToken, int, error)
	OnGetAdminById( ctx context.Context, adminId *int64) (*_adminDtos.AdminFileRes, int, error)
	OnUpdateUserById(ctx context.Context, Id int64, req *_adminEntities.AdminUpdateReq) (int, error)
	AdminDeleted(ctx context.Context, Id int64) (int,error)
	
}