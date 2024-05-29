package entities

type UsersLogin struct {
	Email    string `json:"email" db:"email" form:"email" binding:"required" validate:"required,min=5,max=50"`
	Password string `json:"password" db:"password" form:"password" binding:"required" validate:"required,min=8"`
}
