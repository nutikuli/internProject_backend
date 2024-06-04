package usecase

import (
	"context"
	
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission"
	_adminpermissionDtos "github.com/nutikuli/internProject_backend/internal/models/adminpermission/dtos"
	_adminpermissionEntities "github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	// _fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
	"github.com/gofiber/fiber/v2/log"
)

type adminpermissionUseCase struct {
	adminpermissionRepo adminpermission.AdminPermissionRepository
	fileRepo            file.FileRepository
}

func NewAdminpermissionUsecase(adminpermissionRepo adminpermission.AdminPermissionRepository, fileRepo file.FileRepository) adminpermission.AdminpermissionUseCase {
	return &adminpermissionUseCase{
		adminpermissionRepo: adminpermissionRepo,
		fileRepo:            fileRepo,
	}
}

func (a *adminpermissionUseCase) OnCreateAdminpermissionAccount(c *fiber.Ctx, ctx context.Context, adminpermissionDatReq *_adminpermissionEntities.AdminPermissionCreatedReq) (*_adminpermissionDtos.AdminPermissionFileRes, int, error) {

	newAdminpermissionId, err := a.adminpermissionRepo.CreateAdminPermission(ctx, adminpermissionDatReq)
	log.Debug(err)
	log.Debug("new add=====>",newAdminpermissionId)
	

	adminpermissionRes, errOnGetAdminpermission := a.adminpermissionRepo.GetAdminpermissiomById(ctx, *newAdminpermissionId)
	if errOnGetAdminpermission != nil {
		return nil, http.StatusInternalServerError, errOnGetAdminpermission
	}
	log.Debug("adres=====>",adminpermissionRes)

	return &_adminpermissionDtos.AdminPermissionFileRes{
		AdminpermissionData: adminpermissionRes,
		
	}, http.StatusOK, nil

}

func (a *adminpermissionUseCase) OnGetAdminpermissionById(c *fiber.Ctx, ctx context.Context, adminpermissionId *int64) (*_adminpermissionDtos.AdminPermissionFileRes, int, error) {
	
	adminpermissionRes, errOnGetAdminpermission := a.adminpermissionRepo.GetAdminpermissiomById(ctx, *adminpermissionId)
	if errOnGetAdminpermission != nil {
		return nil, http.StatusInternalServerError, errOnGetAdminpermission
	}

	return &_adminpermissionDtos.AdminPermissionFileRes{
		AdminpermissionData: adminpermissionRes,
		
	}, http.StatusOK, nil
}
