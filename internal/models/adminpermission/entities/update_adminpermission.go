package entities 
type AdminPermissionUpdatedReq struct {
	MenuPermission string `json:"menuPermission" from:"menu_permission" binding:"required"`
}
