package file

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type FileUsecase interface {
	OnUploadFile(c *fiber.Ctx, ctx context.Context, req *entities.FileUploaderReq) (*entities.File, int, error)
	OnGetSourceFiles(c *fiber.Ctx, ctx context.Context) ([]*entities.File, int, error)
	OnDeleteFileByIdAndEntity(c *fiber.Ctx, ctx context.Context, req *entities.FileEntityReq) (int, error)
}
