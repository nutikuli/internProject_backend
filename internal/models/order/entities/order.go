package entities

import (
	"time"
)

type Order struct {
	Id           int64     `db:"id"`
	OrderId      string    `db:"orderId"`
	TotalAmount  float64   `db:"totalAmount"`
	Topic        string    `db:"topic"`
	SumPrice     float64   `db:"sumPrice"`
	State        string    `db:"state"`
	DeliveryType string    `db:"deliveryType"`
	ParcelNumber string    `db:"parcelNumber"`
	SentDate     time.Time `db:"sentDate"`
	CustomerId   int64     `db:"customerId"`
	StoreId      int64     `db:"storeId"`
	BankId       int64     `db:"bankId"`
	createdAt    time.Time `db:"createdAt"`
	updatedAt    time.Time `db:"UpdatedAt"`
}
