package file

import (
	"context"

	"github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type FileRepository interface {
	CreateFileByEntityAndId(ctx context.Context, file *entities.FileUploaderReq, fileEntity *entities.FileEntityReq) (*int64, error)
	GetFiles(ctx context.Context) ([]*entities.File, error)
	GetFilesByIdAndEntity(ctx context.Context, req *entities.FileEntityReq) ([]*entities.File, error)
	DeleteFileByIdAndEntity(ctx context.Context, req *entities.FileEntityReq) error
}
