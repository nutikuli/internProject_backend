package entities

type OrderTransportDetailReq struct {
	OrderID      int64  `json:"order_id" form:"order_id" binding:"required" validate:"required"`
	DeliveryType string `json:"delivery_type" form:"delivery_type" binding:"required" validate:"required"`
	ParcelNumber string `json:"parcel_number" form:"parcel_number" binding:"required" validate:"required"`
	SentDate     string `json:"sent_date" form:"sent_date" binding:"required" validate:"required"`
}
