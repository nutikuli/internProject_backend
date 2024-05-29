package entities

type Product struct {
	Id            int64   `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	Price         float64 `json:"price" db:"price"`
	Stock         int64   `json:"stock" db:"stock"`
	Detail        string  `json:"detail" db:"detail"`
	Status        bool    `json:"status" db:"status"`
	ProductAvatar string  `json:"product_avatar" db:"productAvatar"`
	CategoryId    int64   `json:"category_id" db:"categoryId"`
	StoreId       int64   `json:"store_id" db:"storeId"`
	CreatedAt     string  `json:"created_at" db:"createdAt"`
	UpdatedAt     string  `json:"updated_at" db:"updatedAt"`
}
