package entities

import "time"

type Account struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Password    string `json:"password" db:"password"`
	Phone       string `json:"phone" db:"phone"`
	Location    string `json:"location" db:"location"`
	Email       string `json:"email" db:"email"`
	ImageAvatar string `json:"image_avatar" db:"imageAvatar"`
	// TODO: เพิ่ม file entites ที่เป็น array
	CreatedAt time.Time `json:"created_at" db:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" db:"updatedAt"`
	Role      string    `json:"role" db:"role"`
	Status    bool      `json:"status" db:"status"`
}

type AccountCredentialGetter interface {
	GetId() *int64
	GetEmail() *string
	GetPassword() *string
}
