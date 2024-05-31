package entities

type OrderProductCreateReq struct {
	ProductId int64 `json:"product_id" form:"product_id" binding:"required"`
	Quantity  int64 `json:"ord_quantity" form:"order_id" binding:"required"`
}
