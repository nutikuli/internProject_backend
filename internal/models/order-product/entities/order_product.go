package entities

type OrderProduct struct {
	OrderId   int64  `json:"order_id" db:"orderId"`
	ProductId int64  `json:"product_id" db:"productId"`
	Quantity  int64  `json:"quantity" db:"quantity"`
	CreatedAt string `json:"created_at" db:"createdAt"`
	UpdatedAt string `json:"updated_at" db:"updatedAt"`
}
