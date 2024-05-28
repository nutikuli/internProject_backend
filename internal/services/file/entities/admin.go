package entities

type Admin struct {
	AccountID       int64  `json:"id" db:"id"`
	PermissionID  string `json:"permissionId" db:"permissionId"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}