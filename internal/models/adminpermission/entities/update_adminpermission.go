package entities 
type AdminPermissionUpdatedReq struct {
	MenuPermission []string `json:"menuPermission" from:"menu_permission" binding:"required"`
	Rolename string `json:"roleName" from:"role_name" , binding:"required"`
}
