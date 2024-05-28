package entities
type Store struct {
	AccountID       int64  `json:"id" db:"id"`
	StoreName  string `json:"storeName" db:"storeName"`
	StoreLocation  string `json:"storeLocation" db:"storeLocation"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}