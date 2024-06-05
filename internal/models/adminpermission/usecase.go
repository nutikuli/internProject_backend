package adminpermission

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_adminpermissionDtos "github.com/nutikuli/internProject_backend/internal/models/adminpermission/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
	_adminpermissionEntities "github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
)

type AdminpermissionUseCase interface {
	OnCreateAdminpermissionAccount(c *fiber.Ctx, ctx context.Context, adminpermissionDatReq *_adminpermissionEntities.AdminPermissionCreatedReq) (*_adminpermissionDtos.AdminPermissionFileRes, int, error)
	OnGetAdminpermissionById(c *fiber.Ctx, ctx context.Context, adminpermissionId *int64) (*_adminpermissionDtos.AdminPermissionFileRes, int, error)
	OnUpdateAdminPermissionById(ctx context.Context, adminperId int64, req *entities.AdminPermissionUpdatedReq) (int, error)
	OnDeletedAdminPermission(ctx context.Context, Id int64) (int, error)
}
