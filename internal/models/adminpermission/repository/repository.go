package repository

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"

	"github.com/nutikuli/internProject_backend/internal/models/adminpermission"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
	repositoryquery "github.com/nutikuli/internProject_backend/internal/models/adminpermission/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type AdminPermissionRepo struct {
	db *sqlx.DB
}

func NewAdminPermissionRepository(db *sqlx.DB) adminpermission.AdminPermissionRepository {
	return &AdminPermissionRepo{
		db: db,
	}
}

func (a *AdminPermissionRepo) GetAdminpermissiomById(ctx context.Context, id int64) (*entities.Adminpermission, error) {
	var adminpermission entities.Adminpermission

	err := a.db.GetContext(ctx, &adminpermission, repositoryquery.SQL_get_adminpermission_by_id, "ADMIN", id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &adminpermission, nil
}

func (a *AdminPermissionRepo) CreateAdminPermission(ctx context.Context, adminpermissiondata *entities.AdminPermissionCreatedReq) (*int64, error) {

	res, err := a.db.ExecContext(ctx, repositoryquery.SQL_get_permisson_admin, adminpermissiondata.MenuPermission)

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

func (a *AdminPermissionRepo) GetAdminPermissions(ctx context.Context) (*entities.AdminPermissionCreatedReq, error) {
	var adminpermission entities.AdminPermissionCreatedReq

	err := a.db.GetContext(ctx, &adminpermission, repositoryquery.SQL_get_permisson_admin, "ADMIN")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &adminpermission, nil
}

func (a *AdminPermissionRepo) UpdateAdminPermissionById(ctx context.Context, Id int64, adminpermissiondata *entities.AdminPermissionUpdatedReq) error {
	args := utils.Array{
		adminpermissiondata.MenuPermission,
		Id,
	}

	log.Info(args)

	res, err := a.db.ExecContext(ctx, repositoryquery.SQL_get_adminpermission_by_id, args...)
	if err != nil {
		log.Error(err)
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (a *AdminPermissionRepo) DeleteAdminPermissionById(ctx context.Context, Id int64) error {
	res, err := a.db.ExecContext(ctx, repositoryquery.SQL_delete_adminpermission_by_id, Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
