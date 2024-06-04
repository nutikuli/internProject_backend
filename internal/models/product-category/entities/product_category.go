package entities

type ProductCategory struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Status    bool   `json:"status" db:"status"`
	Code      string `json:"code" db:"code"`
	Detail    string `json:"detail" db:"detail"`
	StoreId   int    `json:"store_id" db:"storeId"`
	CreatedAt string `json:"created_at" db:"createdAt"`
	UpdatedAt string `json:"updated_at" db:"updatedAt"`
}
