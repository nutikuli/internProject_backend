package entities

type OrderStateReq struct {
	OrderID int64  `json:"order_id" form:"order_id" binding:"required" validate:"required"`
	State   string `json:"state" form:"state" binding:"required" validate:"required"`
}
