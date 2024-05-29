package entities

import (
	"time"
)

type OrderCreate struct {
	Id           int64     `json:"order_id" form:"order_id" binding:"required"`
	TotalAmount  float64   `json:"total_amount" form:"total_amount" binding:"required"`
	Topic        string    `json:"order_topic" form:"order_topic" binding:"required"`
	SumPrice     float64   `json:"sum_price" form:"sum_price" binding:"required"`
	State        string    `json:"order_state" form:"order_state" binding:"required"`
	DeliveryType string    `json:"delivery_type" form:"delivery_type" binding:"required"`
	ParcelNumber string    `json:"parcel_number" form:"parcel_number" binding:"required"`
	SentDate     time.Time `json:"sent_date" form:"sent_date" binding:"required"`
	CustomerId   int64     `json:"customer_id" form:"customer_id" binding:"required"`
	StoreId      int64     `json:"store_id" form:"store_id" binding:"required"`
	BankId       int64     `json:"bank_id" form:"bank_id" binding:"required"`
	createdAt    time.Time `json:"created_at" form:"created_at" binding:"required"`
	updatedAt    time.Time `json:"updated_at" form:"updated_at" binding:"required"`
}
