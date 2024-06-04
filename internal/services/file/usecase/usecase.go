package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	"github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type fileUse struct {
	fileRepo file.FileRepository
}

func NewFileUsecase(fileRepo file.FileRepository) file.FileUsecase {
	return &fileUse{
		fileRepo,
	}
}

func (f *fileUse) OnGetSourceFiles(c *fiber.Ctx, ctx context.Context) ([]*entities.File, int, error) {

	files, err := f.fileRepo.GetFiles(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return files, http.StatusOK, nil
}

func (f *fileUse) OnUploadFile(c *fiber.Ctx, ctx context.Context, fileReq *entities.FileUploaderReq, fileEntityReq *entities.FileEntityReq) (*int64, int, error) {
	file := entities.File{
		Name:    fileReq.FileName,
		Type:    fileReq.FileType,
		PathUrl: fileReq.FileData,
	}

	_, fPathDat, status, errOnDecode := file.EncodeBase64toFile(c, true)
	if errOnDecode != nil {
		return nil, status, errOnDecode
	}

	fileReq.FileData = *fPathDat

	fileModel, err := f.fileRepo.CreateFileByEntityAndId(ctx, fileReq, fileEntityReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return fileModel, http.StatusOK, nil
}

func (f *fileUse) OnDeleteFileByIdAndEntity(c *fiber.Ctx, ctx context.Context, fileId int64, req *entities.FileEntityReq) (int, error) {
	err := f.fileRepo.DeleteFileByIdAndEntity(ctx, fileId, req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (f *fileUse) OnUpdateFileByIdAndEntity(c *fiber.Ctx, ctx context.Context, req *entities.FileEntityReq, fileEntityReq *entities.FileUploaderReq) (int, error) {
	_, errOnUpload := f.fileRepo.CreateFileByEntityAndId(ctx, fileEntityReq, req)
	if errOnUpload != nil {
		return http.StatusInternalServerError, errOnUpload
	}

	return http.StatusOK, nil
}
