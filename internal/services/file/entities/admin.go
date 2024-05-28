package entities

type Admin struct {
	PermissionID  string `db:"permissionId"`
	Account
}