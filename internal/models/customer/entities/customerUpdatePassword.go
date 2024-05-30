package entities

type CustomerUpdatePasswordReq struct {
	Password string `json:"password" form:"password" binding:"required"`
}
