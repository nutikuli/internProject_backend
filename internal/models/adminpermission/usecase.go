package adminpermission

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_adminpermissionDtos "github.com/nutikuli/internProject_backend/internal/models/adminpermission/dtos"
	_adminpermissionEntities "github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type AdminpermissionUseCase interface {
	OnCreateAdminpermissionAccount(c *fiber.Ctx, ctx context.Context, adminpermissionDatReq *_adminpermissionEntities.AdminPermissionCreatedReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_adminpermissionDtos.AdminPermissionFileRes, int, error)
}
