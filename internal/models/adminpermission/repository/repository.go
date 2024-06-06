package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"

	// "github.com/nutikuli/internProject_backend/internal/models/account/repository/repository_query"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
	repositoryquery "github.com/nutikuli/internProject_backend/internal/models/adminpermission/repository/repository_query"
	// "github.com/nutikuli/internProject_backend/pkg/utils"
)

type AdminPermissionRepo struct {
	db *sqlx.DB
}

func NewAdminPermissionRepository(db *sqlx.DB) adminpermission.AdminPermissionRepository {
	return &AdminPermissionRepo{
		db: db,
	}
}

func (a *AdminPermissionRepo) GetAdminpermissiomById(ctx context.Context, id int64) ([]*entities.Adminpermission, error) {
	var adminpermission = make([]*entities.Adminpermission, 0)

	err := a.db.SelectContext(ctx, &adminpermission, repositoryquery.SQL_get_adminpermission_by_id, id)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return adminpermission, nil
}

func (a *AdminPermissionRepo) CreateAdminPermission(ctx context.Context, adminpermissiondata *entities.AdminPermissionCreatedReq) (*int64, error) {

	// แปลง slice MenuPermission ไปเป็น JSON string
	menuPermissionJSON, err := json.Marshal(adminpermissiondata.MenuPermission)
	if err != nil {
		log.Info("Error marshaling MenuPermission:", err)
		return nil, err
	}

	// ดำเนินการ SQL query ด้วย JSON string และ Rolename
	res, err := a.db.ExecContext(ctx, repositoryquery.SQL_insert_permission_admin, menuPermissionJSON, adminpermissiondata.Rolename)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	// รับค่า ID ของแถวที่ถูกแทรกล่าสุด
	adminPermissionID, err := res.LastInsertId()
	if err != nil {
		log.Info("Error getting last insert ID:", err)
		return nil, err
	}

	return &adminPermissionID, nil

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

	// res, err := a.db.ExecContext(ctx, repositoryquery.SQL_get_adminpermission_by_id, adminpermissiondata.MenuPermission,Id)
	menuPermissionJSON, err := json.Marshal(adminpermissiondata.MenuPermission)

	// Perform the update operation
	res, err := a.db.ExecContext(ctx, repositoryquery.SQL_update_adminperrmision_by_id, menuPermissionJSON, adminpermissiondata.Rolename, Id)
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


func (a *AdminPermissionRepo) GetAdminpermissionALL(ctx context.Context) ([]*entities.Adminpermission, error) {
	var adminpermission = make([]*entities.Adminpermission, 0)

	err := a.db.SelectContext(ctx, &adminpermission, repositoryquery.SQL_getall_adminpermission)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return adminpermission, nil
}