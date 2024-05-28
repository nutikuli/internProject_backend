package entities
type StoreAccount struct {
	StoreName  string `db:"storeName"`
	StoreLocation  string `db:"storeLocation"`
	Account
}