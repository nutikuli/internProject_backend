package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	"github.com/nutikuli/internProject_backend/internal/services/file/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type fileRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) file.FileRepository {
	return &fileRepo{
		db: db,
	}
}

func (f *fileRepo) CreateFileByEntityAndId(ctx context.Context, file *entities.FileUploaderReq) (*entities.File, error) {
	var newFile entities.File
	args := utils.Array{
		file.FileName,
		file.FileData,
		file.FileType,
		file.EntityType,
		file.EntityId,
	}

	err := f.db.GetContext(ctx, &newFile, repository_query.FileInsertByEntityAndId, args...)
	if err != nil {
		return nil, err
	}

	return &newFile, nil

}

func (f *fileRepo) GetFiles(ctx context.Context) ([]*entities.File, error) {
	var files []*entities.File

	err := f.db.SelectContext(ctx, &files, repository_query.QueryFileSelectAll)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (f *fileRepo) GetFileByIdAndEntity(ctx context.Context, req *entities.FileEntityReq) (*entities.File, error) {
	var file entities.File

	args := utils.Array{
		req.EntityType,
		req.EntityId,
	}

	err := f.db.GetContext(ctx, &file, repository_query.QueryFileSelectByIdAndEntity, args...)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (f *fileRepo) DeleteFileByIdAndEntity(ctx context.Context, req *entities.FileEntityReq) error {
	args := utils.Array{
		req.EntityType,
		req.EntityId,
	}

	_, err := f.db.ExecContext(ctx, repository_query.ExecFileDeleteByIdAndEntity, args...)
	if err != nil {
		return err
	}

	return nil
}
