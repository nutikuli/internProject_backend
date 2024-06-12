package entities

import "time"

type OrderTransportDetailReq struct {
	DeliveryType string     `json:"delivery_type" form:"delivery_type" binding:"required" `
	ParcelNumber string     `json:"parcel_number" form:"parcel_number" binding:"required" `
	SentDate     *time.Time `json:"sent_date" form:"sent_date" `
}
