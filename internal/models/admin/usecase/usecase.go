package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	_accDtos "github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/admin"
	"github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	_adminDtos "github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	_adminEntities "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type adminUseCase struct {
	adminRepo           admin.AdminRepository
	fileRepo            file.FileRepository
	accUsecase          account.AccountUsecase
	adminpermissionRepo adminpermission.AdminPermissionRepository
}

func NewAdminUsecase(adminRepo admin.AdminRepository, fileRepo file.FileRepository, accUsecase account.AccountUsecase, adminpermissionRepo adminpermission.AdminPermissionRepository) admin.AdminUseCase {
	return &adminUseCase{
		adminRepo:           adminRepo,
		fileRepo:            fileRepo,
		adminpermissionRepo: adminpermissionRepo,
		accUsecase:          accUsecase,
	}
}

func (a *adminUseCase) OnCreateAdminAccount(c *fiber.Ctx, ctx context.Context, adminDatReq *_adminEntities.AdminRegisterReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_adminDtos.AdminFileRes, *_accDtos.UserToken, int, error) {

	accRegister, usrCred, status, errOnRegister := a.accUsecase.Register(ctx, adminDatReq)
	if errOnRegister != nil {
		return nil, nil, status, errOnRegister
	}

	adminDatReq.Password = usrCred.Password

	newAdminId, err := a.adminRepo.CreateAdmin(ctx, adminDatReq)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ACCOUNT",
		EntityId:   *newAdminId,
	}

	for _, fDatReq := range filesDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "ACCOUNT",
			AccountId:   newAdminId,
		}

		_, fUrl, status ,errOnCreatedFile := file.EncodeBase64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, nil, status, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		_, errOnInsertFile := a.fileRepo.CreateFileByEntityAndId(ctx, fDatReq, fileEntity)
		if errOnInsertFile != nil {
			return nil, nil, http.StatusInternalServerError, errOnInsertFile
		}
	}
	filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, nil, http.StatusInternalServerError, errOnGetFiles
	}

	adminRes, errOnGetAdmin := a.adminRepo.GetAccountAdminById(ctx, *newAdminId)
	if errOnGetAdmin != nil {
		return nil, nil, http.StatusInternalServerError, errOnGetAdmin
	}

	return &_adminDtos.AdminFileRes{
		AdminData: adminRes,
		FilesData: filesRes,
	}, accRegister, http.StatusOK, nil

}

func (a *adminUseCase) OnGetAdminById(ctx context.Context, adminId int64) (*_adminDtos.AdminFileRes, int, error) {
	
	log.Debug("adminid=====>",adminId)
	
	fileEntity := &_fileEntities.FileEntityReq{
		EntityType: "ACCOUNT",
		EntityId:   adminId,
	}

	filesRes, errOnGetFiles := a.fileRepo.GetFilesByIdAndEntity(ctx, fileEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	} 

	log.Debug("fileres=====>",filesRes)

	adminRes, errOnGetAdmin := a.adminRepo.GetAccountAdminById(ctx, adminId)
	if errOnGetAdmin != nil {
		return nil, http.StatusInternalServerError, errOnGetAdmin
	}
	log.Debug("adminres=====>",adminRes)

	adminPermissionRes, errOnGetAdminPermission := a.adminpermissionRepo.GetAdminpermissiomById(ctx, adminId)
	if errOnGetAdmin != nil {
		return nil, http.StatusInternalServerError, errOnGetAdminPermission
	} 

	

	return &_adminDtos.AdminFileRes{
		AdminData:           adminRes,
		AdminpermissionData: adminPermissionRes,
		FilesData:           filesRes,
	}, http.StatusOK, nil
}

func (a *adminUseCase) OnUpdateUserById(ctx context.Context, Id int64, req *entities.AdminUpdateReq) (int, error) {

	err := a.adminRepo.UpdateAdminById(ctx, Id, req)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update user by ID: %w", err)
	}

	return http.StatusOK, nil
}

func (a *adminUseCase) AdminDeleted(ctx context.Context, Id int64) (int, error) {

	err := a.adminRepo.DeleteAdminById(ctx, Id)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete admin by ID: %w", err)
	}

	return http.StatusOK, nil
} 



func (a *adminUseCase) OnGetAllUserAdmin(ctx context.Context) ([]*dtos.AdminFileRes, int, error) {
	admins, err := a.adminRepo.GetAccountAdmins(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var adminFileRes []*dtos.AdminFileRes

	for _, admin := range admins {
		fEntity := &_fileEntities.FileEntityReq{
			EntityType: "ADMIN",
			EntityId:   admin.Id, // Assuming product has an ID field
		}

		files, err := a.fileRepo.GetFilesByIdAndEntity(ctx, fEntity)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		adminFileRes = append(adminFileRes, &dtos.AdminFileRes{
			AdminData: admin,
			FilesData:   files,
		})
	}

	return adminFileRes, http.StatusOK, nil
}  









