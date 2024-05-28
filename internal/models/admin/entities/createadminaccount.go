package entities

//TODO:นำตัวไปใช้ใน Log Model
type AdminCreatedReq struct {
	Id           int64  `db:"id"`
	Name         string `db:"name"`
	Password     string `db:"password"`
	Phone        string `db:"phone"`
	Location     string `db:"location"`
	Email        string `db:"email"`
	ImageAvatar  string `db:"imageAvatar"`
	Role         string `db:"role"`
	Status       bool   `db:"status"`
	PermissionID string `db:"permissionId"`
}
