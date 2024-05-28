package entities
type Customer struct {
	AccountID       int64  `json:"id" db:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}