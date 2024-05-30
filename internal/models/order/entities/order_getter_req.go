package entities

type StoreAndOrderIdReq struct {
	StoreId int64 `json:"store_id" form:"store_id" binding:"required"`
	OrderId int64 `json:"order_id" form:"order_id" binding:"required"`
}
