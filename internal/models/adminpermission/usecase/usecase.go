package usecase

import (
	"context"
	"fmt"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission"
	_adminpermissionDtos "github.com/nutikuli/internProject_backend/internal/models/adminpermission/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
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


func (a *adminpermissionUseCase) OnUpdateAdminPermissionById(ctx context.Context, adminperId int64, req *entities.AdminPermissionUpdatedReq) (int, error) {

	err := a.adminpermissionRepo.UpdateAdminPermissionById(ctx, adminperId, req)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update user by ID: %w", err)
	}

	return http.StatusOK, nil
} 

func (a *adminpermissionUseCase) OnDeletedAdminPermission(ctx context.Context, Id int64) (int, error) {

	err := a.adminpermissionRepo.DeleteAdminPermissionById(ctx, Id)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete user by ID: %w", err)
	}

	return http.StatusOK, nil
}

