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

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account_admin, "ADMIN")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &admin, nil
}

func (a *AdminRepo) GetAccountAdminById(ctx context.Context, id *int64) (*entities.Admin, error) {
	var admin entities.Admin

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account_admin_by_id, "ADMIN", *id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return &admin, nil
}

func (a *AdminRepo) CreateAdmin(ctx context.Context, admindata *entities.AdminCreatedReq) (*int64, error) {

	adminrole := "ADMIN"
	adminstatus := 1
	res, err := a.db.ExecContext(ctx, repository_query.SQL_insert_account_admin, admindata.Name, admindata.Password,
		admindata.Phone, admindata.Location, admindata.Email, adminrole, adminstatus, admindata.PermissionID)
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


