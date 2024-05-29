package entities

type AdminPermissionCreatedReq struct {
	MenuPermission string `json:"menuPermission" from:"menu_permission" binding:"required"`
}