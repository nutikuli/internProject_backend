package entities

type OrderStateReq struct {
	State string `json:"state" form:"state" binding:"required" validate:"required"`
}
