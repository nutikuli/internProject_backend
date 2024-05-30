package entities

import (
	"time"
)

type OrderProduct struct {
	OrderId   int64     `db:"orderId"`
	ProductId int64     `db:"productId"`
	Quantity  int64     `db:"quantity"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}
