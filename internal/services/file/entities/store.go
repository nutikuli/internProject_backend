package entities
type Store struct {
	StoreName  string `json:"storeName" db:"storeName"`
	StoreLocation  string `json:"storeLocation" db:"storeLocation"`
	Account
}