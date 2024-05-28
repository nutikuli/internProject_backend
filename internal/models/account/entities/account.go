package entities

import "time"

type Account struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Password    string `db:"password"`
	Phone       string `db:"phone"`
	Location    string `db:"location"`
	Email       string `db:"email"`
	ImageAvatar string `db:"imageAvatar"`
	// TODO: เพิ่ม file entites ที่เป็น array
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
	Role      string    `db:"role"`
	Status    bool      `db:"status"`
}
