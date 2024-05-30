package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/admin"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type adminUseCase struct {
	adminRepo admin.AdminRepository
	fileRepo file.FileRepository
}


func NewAdminUsecase(adminRepo admin.AdminRepository, fileRepo file.FileRepository) admin.AdminUseCase {
	return &adminUseCase{
		adminRepo: adminRepo,
		fileRepo:  fileRepo,
	}
}


func (a *adminUseCase) OnCreateAdminAccount(c *fiber.Ctx, ctx context.Context, adminDatReq *_adminEntities.AdminCreatedReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_adminDtos.AdminFileRes, int, error) {

	newAdminId, err := a.adminRepo.CreateAdmin(ctx, adminDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ADMIN",
		EntityId:   *newAdminId,
	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "ADMIN",
			EntityId:   *newAdminId,
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

	adminRes, errOnGetAdmin := a.adminRepo.GetAccountAdminById(ctx, newAdminId)
	if errOnGetAdmin != nil {
		return nil, http.StatusInternalServerError, errOnGetAdmin
	}

	return &_adminDtos.AdminFileRes{
		AdminData: adminRes,
		FilesData: filesRes,
	}, http.StatusOK, nil

}



func (a *adminUseCase) OnGetAdminById(c *fiber.Ctx, ctx context.Context, adminId *int64) (*_adminDtos.AdminFileRes, int, error) {
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ADMIN",
		EntityId:   *adminId,
	}

	filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	adminRes, errOnGetAdmin := a.adminRepo.GetAccountAdminById(ctx , adminId)
	if errOnGetAdmin != nil {
		return nil, http.StatusInternalServerError, errOnGetAdmin
	}

	return &_adminDtos.AdminFileRes{
		AdminData: adminRes,
		FilesData: filesRes,
	}, http.StatusOK, nil
}


func (a *adminUseCase) OnUpdateUserById(ctx context.Context, Id int64, req *entities.AdminUpdateReq) (int, error) {

	err := a.adminRepo.UpdateAdminById(ctx, Id, req)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update user by ID: %w", err)
	}

	return http.StatusOK, nil
} 


func (a *adminUseCase) AdminDeleted(ctx context.Context, Id int64) (int,error) {

	err := a.adminRepo.DeleteAdminById(ctx, Id)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete admin by ID: %w", err)
	}

	return http.StatusOK, nil
}