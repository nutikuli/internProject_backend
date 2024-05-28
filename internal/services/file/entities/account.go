package entities

type Account struct {
	Id        int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Password  string `json:"password" db:"password"`
	Phone string `json:"phone" db:"phone"`
	Location string `json:"location" db:"location"`
	Email string `json:"email" db:"email"`
	ImageAvatar string `json:"imageAvatar" db:"imageAvatar"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Role string `json:"role"`
	Status bool `json:"status"`
}
