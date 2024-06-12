package entities

type CustomerUpdateReq struct {
	Name     string `json:"acc_name" from:"acc_name" binding:"required"`
	Phone    string `json:"acc_phone" from:"acc_phone" binding:"required"`
	Location string `json:"acc_location" from:"acc_location" binding:"required"`
	Email    string `json:"acc_email" from:"acc_email" binding:"required"`
	Status   bool   `json:"acc_status" form:"acc_status" binding:"required"`
}
