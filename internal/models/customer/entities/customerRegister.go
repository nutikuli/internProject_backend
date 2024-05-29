package entities

type CustomerRegisterReq struct {
	Id       int64  `json:"id" form:"id" binding:"required"`
	Name     string `json:"acc_name" form:"acc_name" binding:"required"`
	Password string ` json:"acc_password" form:"acc_password" binding:"required"`
	Phone    string ` json:"acc_phone" form:"acc_phone" binding:"required"`
	Location string ` json:"acc_location" form:"acc_location" binding:"required"`
	Email    string ` json:"acc_email" form:"acc_email" binding:"required"`
	Role     string ` json:"acc_role" form:"acc_role" binding:"required"`
	Status   bool   ` json:"acc_status" form:"acc_status" binding:"required"`
}

func (u *CustomerRegisterReq) GetEmail() *string {
	return &u.Email
}

func (u *CustomerRegisterReq) GetPassword() *string {
	return &u.Password
}

func (u *CustomerRegisterReq) GetRole() *string {
	return &u.Role
}
