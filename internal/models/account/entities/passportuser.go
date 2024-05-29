package entities

type UsersPassport struct {
	Id        int64  `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Role      string `json:"role" db:"role"`
	CreatedAt string `json:"created_at" db:"createdAt"`
	UpdatedAt string `json:"updated_at" db:"createdAt"`
}
