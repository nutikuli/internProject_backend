package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission"
	_adminpermissionDtos "github.com/nutikuli/internProject_backend/internal/models/adminpermission/dtos"
	_adminpermissionEntities "github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
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

func (a *adminpermissionUseCase) OnCreateAdminpermissionAccount(c *fiber.Ctx, ctx context.Context, adminpermissionDatReq *_adminpermissionEntities.AdminPermissionCreatedReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_adminpermissionDtos.AdminPermissionFileRes, int, error) {

	newAdminpermissionId, err := a.adminpermissionRepo.CreateAdminPermission(ctx, adminpermissionDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ADMIN",
		EntityId:   *newAdminpermissionId,
	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "ADMIN",
			AccountId:  *newAdminpermissionId,
		}

		_, fUrl, errOnCreatedFile := file.Base64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, http.StatusConflict, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		_, errOnInsertFile := a.fileRepo.CreateFileByEntityAndId(ctx, fDatReq, fileEntity)
		if errOnInsertFile != nil {
			return nil, http.StatusInternalServerError, errOnInsertFile
		}
	}
	filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	adminpermissionRes, errOnGetAdminpermission := a.adminpermissionRepo.GetAdminpermissiomById(ctx, *newAdminpermissionId)
	if errOnGetAdminpermission != nil {
		return nil, http.StatusInternalServerError, errOnGetAdminpermission
	}

	return &_adminpermissionDtos.AdminPermissionFileRes{
		AdminpermissionData: adminpermissionRes,
		FilesData:           filesRes,
	}, http.StatusOK, nil

}

func (a *adminpermissionUseCase) OnGetAdminpermissionById(c *fiber.Ctx, ctx context.Context, adminpermissionId *int64) (*_adminpermissionDtos.AdminPermissionFileRes, int, error) {
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ADMIN",
		EntityId:   *adminpermissionId,
	}

	filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	adminpermissionRes, errOnGetAdminpermission := a.adminpermissionRepo.GetAdminpermissiomById(ctx, *adminpermissionId)
	if errOnGetAdminpermission != nil {
		return nil, http.StatusInternalServerError, errOnGetAdminpermission
	}

	return &_adminpermissionDtos.AdminPermissionFileRes{
		AdminpermissionData: adminpermissionRes,
		FilesData:           filesRes,
	}, http.StatusOK, nil
}
