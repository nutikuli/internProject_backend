package entities

import(
	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
)


type Admin struct {
	
	entities.Account
	PermissionID  string `db:"permissionId"`
}

