package entities 


type AdminUpdateReq struct {
	Name         string `json:"name" from:"name" binding:"required"`
	Password     string `json:"password" from:"password" binding:"required"`
	Phone        string `json:"phone" from:"phone" binding:"required"`
	Location     string `json:"location" from:"location" binding:"required"`
	Email        string `json:"email" from:"email" binding:"required"`
	Role         string `json:"role" from:"role" binding:"required"`
	Status       bool   `json:"status" from:"status" binding:"required"`
	PermissionID string `json:"permissino_id" from:"permissino_id" binding:"required"`
	

	
}

