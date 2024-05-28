package entities
type Store struct {
	StoreName  string `db:"storeName"`
	StoreLocation  string `db:"storeLocation"`
	Account
}