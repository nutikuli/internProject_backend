package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
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
	log.Debug("args: ", args)

	var SQL_FileInsertByEntityAndId string
	switch fileEntity.EntityType {
	case "ACCOUNT":
		SQL_FileInsertByEntityAndId = "INSERT INTO File (name, pathUrl, type,  entityType, accountId) VALUES (?, ?, ?, ?, ?)"
	case "PRODUCT":
		SQL_FileInsertByEntityAndId = "INSERT INTO File (name, pathUrl, type,  entityType, productId) VALUES (?, ?, ?, ?, ?)"
	case "ORDER":
		SQL_FileInsertByEntityAndId = "INSERT INTO File (name, pathUrl, type,  entityType, orderId) VALUES (?, ?, ?, ?, ?)"
	case "BANK":
		SQL_FileInsertByEntityAndId = "INSERT INTO File (name, pathUrl, type,  entityType, bankId) VALUES (?, ?, ?, ?, ?)"
	default:
		panic("invalid entity type")
	}

	result, err := f.db.ExecContext(ctx, SQL_FileInsertByEntityAndId, args...)
	if err != nil {
		log.Debug("err on CreateFileByEntityAndId: ", err)
		return nil, err
	}

	newFileID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &newFileID, nil

}

func (f *fileRepo) GetFiles(ctx context.Context) ([]*entities.File, error) {
	var files = make([]*entities.File, 0)

	err := f.db.SelectContext(ctx, &files, repository_query.QueryFileSelectAll)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (f *fileRepo) GetFilesByIdAndEntity(ctx context.Context, req *entities.FileEntityReq) ([]*entities.File, error) {
	var files = make([]*entities.File, 0)

	args := utils.Array{
		req.EntityType,
		req.EntityId,
	}

	log.Debug("args: ", args)

	var SQL_QueryFileSelectByIdAndEntity string
	switch req.EntityType {
	case "ACCOUNT":
		SQL_QueryFileSelectByIdAndEntity = "SELECT * FROM File WHERE entityType = ? AND accountId = ?"
	case "PRODUCT":
		SQL_QueryFileSelectByIdAndEntity = "SELECT * FROM File WHERE entityType = ? AND productId = ?"
	case "ORDER":
		SQL_QueryFileSelectByIdAndEntity = "SELECT * FROM File WHERE entityType = ? AND orderId = ?"
	case "BANK":
		SQL_QueryFileSelectByIdAndEntity = "SELECT * FROM File WHERE entityType = ? AND bankId = ?"
	default:
		panic("invalid entity type")
	}

	err := f.db.SelectContext(ctx, &files, SQL_QueryFileSelectByIdAndEntity, args...)
	if err != nil {
		log.Debug("err on GetFilesByIdAndEntity: ", err)
		return nil, err
	}

	return files, nil
}

func (f *fileRepo) DeleteFileByIdAndEntity(ctx context.Context, fileId int64, req *entities.FileEntityReq) error {
	args := utils.Array{
		req.EntityType,
		fileId,
	}

	// var SQL_ExecFileDeleteByIdAndEntity string
	// switch req.EntityType {
	// case "ACCOUNT":
	// 	SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND accountId = ?"
	// case "PRODUCT":
	// 	SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND productId = ?"
	// case "ORDER":
	// 	SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND orderId = ?"
	// case "BANK":
	// 	SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND bankId = ?"
	// default:
	// 	panic("invalid entity type")
	// }

	var SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND id = ?"

	_, err := f.db.ExecContext(ctx, SQL_ExecFileDeleteByIdAndEntity, args...)
	if err != nil {
		return err
	}

	return nil
}

//- Don't used right now
// func (f *fileRepo) UpdateFileByIdAndEntity(ctx context.Context, req *entities.FileEntityReq, file *entities.FileUploaderReq) error {
// 	args := utils.Array{
// 		file.FileName,
// 		file.FileData,
// 		file.FileType,
// 		req.EntityType,
// 		req.EntityId,
// 	}

// 	var SQL_ExecFileUpdateByIdAndEntity string
// 	switch req.EntityType {
// 	case "ACCOUNT":
// 		SQL_ExecFileUpdateByIdAndEntity = "UPDATE File SET name = ?, pathUrl = ?, type = ? WHERE entityType = ? AND accountId = ?"
// 	case "PRODUCT":
// 		SQL_ExecFileUpdateByIdAndEntity = "UPDATE File SET name = ?, pathUrl = ?, type = ? WHERE entityType = ? AND productId = ?"
// 	case "ORDER":
// 		SQL_ExecFileUpdateByIdAndEntity = "UPDATE File SET name = ?, pathUrl = ?, type = ? WHERE entityType = ? AND orderId = ?"
// 	case "BANK":
// 		SQL_ExecFileUpdateByIdAndEntity = "UPDATE File SET name = ?, pathUrl = ?, type = ? WHERE entityType = ? AND bankId = ?"
// 	default:
// 		panic("invalid entity type")
// 	}

// 	_, err := f.db.ExecContext(ctx, SQL_ExecFileUpdateByIdAndEntity, args...)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// package repository

// import (
// 	"context"

// 	"github.com/gofiber/fiber/v2/log"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/nutikuli/internProject_backend/internal/services/file"
// 	"github.com/nutikuli/internProject_backend/internal/services/file/entities"
// 	"github.com/nutikuli/internProject_backend/internal/services/file/repository/repository_query"
// 	"github.com/nutikuli/internProject_backend/pkg/utils"
// )

// type fileRepo struct {
// 	db *sqlx.DB
// }

// func NewFileRepository(db *sqlx.DB) file.FileRepository {
// 	return &fileRepo{
// 		db: db,
// 	}
// }

// func (f *fileRepo) CreateFileByEntityAndId(ctx context.Context, file *entities.FileUploaderReq, fileEntity *entities.FileEntityReq) (*int64, error) {
// 	var args utils.Array

// 	var SQL_FileInsertByEntityAndId string
// 	switch fileEntity.EntityType {
// 	case "ACCOUNT":
// 		SQL_FileInsertByEntityAndId = "INSERT INTO File (name, pathUrl, type,  entityType, accountId) VALUES (?, ?, ?, ?, ?)"
// 		args = utils.Array{
// 			file.FileName,
// 			file.FileData,
// 			file.FileType,
// 			fileEntity.EntityType,
// 			struct {
// 				AccountId int64 `db:"accountId"`
// 			}{AccountId: fileEntity.EntityId},
// 		}
// 	case "PRODUCT":
// 		SQL_FileInsertByEntityAndId = "INSERT INTO File (name, pathUrl, type,  entityType, productId) VALUES (?, ?, ?, ?, ?)"
// 		args = utils.Array{file.FileName, file.FileData, file.FileType, fileEntity.EntityType, struct {
// 			ProductId int64 `db:"productId"`
// 		}{ProductId: fileEntity.EntityId}}
// 	case "ORDER":
// 		SQL_FileInsertByEntityAndId = "INSERT INTO File (name, pathUrl, type,  entityType, orderId) VALUES (?, ?, ?, ?, ?)"
// 		args = utils.Array{file.FileName, file.FileData, file.FileType, fileEntity.EntityType, struct {
// 			OrderId int64 `db:"orderId"`
// 		}{OrderId: fileEntity.EntityId}}
// 	case "BANK":
// 		SQL_FileInsertByEntityAndId = "INSERT INTO File (name, pathUrl, type,  entityType, bankId) VALUES (?, ?, ?, ?, ?)"
// 		args = utils.Array{file.FileName, file.FileData, file.FileType, fileEntity.EntityType, struct {
// 			BankId int64 `db:"bankId"`
// 		}{BankId: fileEntity.EntityId}}
// 	default:
// 		panic("invalid entity type")
// 	}
// 	log.Debug("args: ", args)

// 	result, err := f.db.ExecContext(ctx, SQL_FileInsertByEntityAndId, args...)
// 	if err != nil {
// 		log.Debug("err on CreateFileByEntityAndId: ", err)
// 		return nil, err
// 	}

// 	newFileID, err := result.LastInsertId()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &newFileID, nil

// }

// func (f *fileRepo) GetFiles(ctx context.Context) ([]*entities.File, error) {
// 	var files = make([]*entities.File, 0)

// 	err := f.db.SelectContext(ctx, &files, repository_query.QueryFileSelectAll)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return files, nil
// }

// func (f *fileRepo) GetFilesByIdAndEntity(ctx context.Context, req *entities.FileEntityReq) ([]*entities.File, error) {
// 	var files = make([]*entities.File, 0)

// 	var args utils.Array

// 	var SQL_QueryFileSelectByIdAndEntity string
// 	switch req.EntityType {
// 	case "ACCOUNT":
// 		SQL_QueryFileSelectByIdAndEntity = "SELECT * FROM File WHERE entityType = ? AND accountId = ?"
// 		args = utils.Array{
// 			req.EntityType,
// 			struct {
// 				AccountId int64 `db:"accountId"`
// 			}{AccountId: req.EntityId},
// 		}
// 	case "PRODUCT":
// 		SQL_QueryFileSelectByIdAndEntity = "SELECT * FROM File WHERE entityType = ? AND productId = ?"
// 		args = utils.Array{
// 			req.EntityType,
// 			struct {
// 				ProductId int64 `db:"productId"`
// 			}{ProductId: req.EntityId},
// 		}
// 	case "ORDER":
// 		SQL_QueryFileSelectByIdAndEntity = "SELECT * FROM File WHERE entityType = ? AND orderId = ?"
// 		args = utils.Array{
// 			req.EntityType,
// 			struct {
// 				OrderId int64 `db:"orderId"`
// 			}{OrderId: req.EntityId},
// 		}
// 	case "BANK":
// 		SQL_QueryFileSelectByIdAndEntity = "SELECT * FROM File WHERE entityType = ? AND bankId = ?"
// 		args = utils.Array{
// 			req.EntityType,
// 			struct {
// 				BankId int64 `db:"bankId"`
// 			}{BankId: req.EntityId},
// 		}
// 	default:
// 		panic("invalid entity type")
// 	}

// 	err := f.db.SelectContext(ctx, &files, SQL_QueryFileSelectByIdAndEntity, args...)
// 	if err != nil {
// 		log.Debug("err on GetFilesByIdAndEntity: ", err)
// 		return nil, err
// 	}

// 	return files, nil
// }

// func (f *fileRepo) DeleteFileByIdAndEntity(ctx context.Context, req *entities.FileEntityReq) error {

// 	var args utils.Array

// 	var SQL_ExecFileDeleteByIdAndEntity string
// 	switch req.EntityType {
// 	case "ACCOUNT":
// 		SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND accountId = ?"
// 		args = utils.Array{
// 			req.EntityType,
// 			struct {
// 				AccountId int64 `db:"accountId"`
// 			}{AccountId: req.EntityId},
// 		}
// 	case "PRODUCT":
// 		SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND productId = ?"
// 		args = utils.Array{
// 			req.EntityType,
// 			struct {
// 				ProductId int64 `db:"productId"`
// 			}{ProductId: req.EntityId},
// 		}
// 	case "ORDER":
// 		SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND orderId = ?"
// 		args = utils.Array{
// 			req.EntityType,
// 			struct {
// 				OrderId int64 `db:"orderId"`
// 			}{OrderId: req.EntityId},
// 		}
// 	case "BANK":
// 		SQL_ExecFileDeleteByIdAndEntity = "DELETE FROM File WHERE entityType = ? AND bankId = ?"
// 		args = utils.Array{
// 			req.EntityType,
// 			struct {
// 				BankId int64 `db:"bankId"`
// 			}{BankId: req.EntityId},
// 		}
// 	default:
// 		panic("invalid entity type")
// 	}

// 	_, err := f.db.ExecContext(ctx, SQL_ExecFileDeleteByIdAndEntity, args...)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UpdateFileByIdAndEntity implements file.FileRepository.
// func (f *fileRepo) UpdateFileByIdAndEntity(ctx context.Context, req *entities.FileEntityReq, file *entities.FileUploaderReq) error {
// 	var args utils.Array

// 	var SQL_ExecFileUpdateByIdAndEntity string
// 	switch req.EntityType {
// 	case "ACCOUNT":
// 		SQL_ExecFileUpdateByIdAndEntity = "UPDATE File SET name = ?, pathUrl = ?, type = ? WHERE entityType = ? AND accountId = ?"
// 		args = utils.Array{
// 			file.FileName,
// 			file.FileData,
// 			file.FileType,
// 			req.EntityType,
// 			struct {
// 				AccountId int64 `db:"accountId"`
// 			}{AccountId: req.EntityId},
// 		}
// 	case "PRODUCT":
// 		SQL_ExecFileUpdateByIdAndEntity = "UPDATE File SET name = ?, pathUrl = ?, type = ? WHERE entityType = ? AND productId = ?"
// 		args = utils.Array{
// 			file.FileName,
// 			file.FileData,
// 			file.FileType,
// 			req.EntityType,
// 			struct {
// 				ProductId int64 `db:"productId"`
// 			}{ProductId: req.EntityId},
// 		}
// 	case "ORDER":
// 		SQL_ExecFileUpdateByIdAndEntity = "UPDATE File SET name = ?, pathUrl = ?, type = ? WHERE entityType = ? AND orderId = ?"
// 		args = utils.Array{
// 			file.FileName,
// 			file.FileData,
// 			file.FileType,
// 			req.EntityType,

// 			struct {
// 				OrderId int64 `db:"orderId"`
// 			}{OrderId: req.EntityId},
// 		}
// 	case "BANK":
// 		SQL_ExecFileUpdateByIdAndEntity = "UPDATE File SET name = ?, pathUrl = ?, type = ? WHERE entityType = ? AND bankId = ?"
// 		args = utils.Array{
// 			file.FileName,
// 			file.FileData,
// 			file.FileType,
// 			req.EntityType,
// 			struct {
// 				BankId int64 `db:"bankId"`
// 			}{BankId: req.EntityId},
// 		}
// 	default:
// 		panic("invalid entity type")
// 	}

// 	_, err := f.db.ExecContext(ctx, SQL_ExecFileUpdateByIdAndEntity, args...)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
