package entities

type Admin struct {
	PermissionID  string `json:"permissionId" db:"permissionId"`
	Account
}