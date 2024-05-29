package entities

import(
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
)


type Admin struct {
	PermissionID  string `db:"permissionId"`
	entities.Account
}