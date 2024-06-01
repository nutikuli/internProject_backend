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

func (f *fileRepo) CreateFileByEntityAndId(ctx context.Context, file *entities.FileUploaderReq, fileEntity *entities.FileEntityReq) (*int64, error) {
	args := utils.Array{
		file.FileName,
		file.FileData,
		file.FileType,
		fileEntity.EntityType,
		fileEntity.EntityId,
	}

	result, err := f.db.ExecContext(ctx, repository_query.FileInsertByEntityAndId, args...)
	if err != nil {
		return nil, err
	}

	newFileID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &newFileID, nil

}

func (f *fileRepo) GetFiles(ctx context.Context) ([]*entities.File, error) {
	var files []*entities.File

	err := f.db.SelectContext(ctx, &files, repository_query.QueryFileSelectAll)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (f *fileRepo) GetFilesByIdAndEntity(ctx context.Context, req *entities.FileEntityReq) ([]*entities.File, error) {
	var files []*entities.File

	args := utils.Array{
		req.EntityType,
		req.EntityId,
	}

	err := f.db.SelectContext(ctx, &files, repository_query.QueryFileSelectByIdAndEntity, args...)
	if err != nil {
		return nil, err
	}

	return files, nil
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

// UpdateFileByIdAndEntity implements file.FileRepository.
func (f *fileRepo) UpdateFileByIdAndEntity(ctx context.Context, req *entities.FileEntityReq, file *entities.FileUploaderReq) error {
	args := utils.Array{
		file.FileName,
		file.FileData,
		file.FileType,
		req.EntityType,
		req.EntityId,
	}

	_, err := f.db.ExecContext(ctx, repository_query.ExecFileUpdateByIdAndEntity, args...)
	if err != nil {
		return err
	}

	return nil
}
