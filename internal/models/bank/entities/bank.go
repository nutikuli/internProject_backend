package entities

type Bank struct {
	Id         int64  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	AccNumber  string `json:"acc_number" db:"accNumber"`
	AccName    string `json:"acc_name" db:"accName"`
	AvartarUrl string `json:"avartar_url" db:"avartarUrl"`
	Status     bool   `json:"status" db:"status"`
	StoreId    int64  `json:"store_id" db:"storeId"`
}
