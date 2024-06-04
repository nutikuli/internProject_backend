package file

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type FileUsecase interface {
	OnUploadFile(c *fiber.Ctx, ctx context.Context, fileReq *entities.FileUploaderReq, fileEntityReq *entities.FileEntityReq) (*int64, int, error)
	OnGetSourceFiles(c *fiber.Ctx, ctx context.Context) ([]*entities.File, int, error)
	OnDeleteFileByIdAndEntity(c *fiber.Ctx, ctx context.Context, fileId int64, req *entities.FileEntityReq) (int, error)
	OnUpdateFileByIdAndEntity(c *fiber.Ctx, ctx context.Context, req *entities.FileEntityReq, fileEntityReq *entities.FileUploaderReq) (int, error)
}
