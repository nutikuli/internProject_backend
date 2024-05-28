package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/admin"
	"github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	"github.com/nutikuli/internProject_backend/internal/models/admin/repository/repository_query"

)

type AdminRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) admin.AdminRepository {
	return &AdminRepo{
		db: db,
	}
} 


func (a *AdminRepo) GetAccountAdmins(ctx context.Context) (*entities.Admin, error) {
	var admin entities.Admin

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account_admin, "admin")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &admin, nil
}

func (a *AdminRepo) GetAccountAdminById(ctx context.Context, id *int64) (*entities.Admin, error) {
	var admin entities.Admin

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account_admin_by_id, "admin", *id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &admin, nil
}



func (a *AdminRepo) CreateFile(ctx context.Context, file *entities.createadmin) (*int64, error) {
	if err != nil {
		log.Info(err)
		return nil, err
	}

	args := utils.Array{
		file.FileName,
		b64,
		file.FileType,
	}

	res, err := r.db.ExecContext(ctx, repository_query.CreateFile, args...)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	createdId, err := res.LastInsertId()
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &createdId, nil
}
