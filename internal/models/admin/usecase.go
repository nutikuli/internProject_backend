package admin

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type AdminUseCase interface {
	OnCreateAdminAccount(c *fiber.Ctx, ctx context.Context, adminDatReq *_adminEntities.AdminCreatedReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_adminDtos.AdminFileRes, int, error)
	OnGetAdminById(c *fiber.Ctx, ctx context.Context, adminId *int64) (*_adminDtos.AdminFileRes, int, error)
	OnUpdateUserById(ctx context.Context, Id int64, req *_adminEntities.AdminUpdateReq) (int, error)
}