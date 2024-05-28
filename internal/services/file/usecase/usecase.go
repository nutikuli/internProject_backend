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

func (f *fileUse) OnUploadFile(c *fiber.Ctx, ctx context.Context, req *entities.FileUploaderReq) (*entities.File, int, error) {
	file := entities.File{
		Name:    req.FileName,
		Type:    req.FileType,
		PathUrl: req.FileData,
	}

	_, fPathDat, status, errOnDecode := file.EncodeBase64toFile(c, true)
	if errOnDecode != nil {
		return nil, status, errOnDecode
	}

	req.FileData = *fPathDat

	fileModel, err := f.fileRepo.CreateFileByEntityAndId(ctx, req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return fileModel, http.StatusOK, nil
}

func (f *fileUse) OnDeleteFileByIdAndEntity(c *fiber.Ctx, ctx context.Context, req *entities.FileEntityReq) (int, error) {
	err := f.fileRepo.DeleteFileByIdAndEntity(ctx, req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
