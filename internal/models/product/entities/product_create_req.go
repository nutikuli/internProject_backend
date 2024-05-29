package entities

type ProductCreateReq struct {
	Name          string  `json:"name" form:"name" validate:"required" binding:"required"`
	Price         float64 `json:"price" form:"price" validate:"required" binding:"required"`
	Stock         int64   `json:"stock" form:"stock" validate:"required" binding:"required"`
	Detail        string  `json:"detail" form:"detail" validate:"required" binding:"required"`
	Status        bool    `json:"status" form:"status" validate:"required" binding:"required"`
	ProductAvatar string  `json:"product_avatar" form:"product_avatar" validate:"required" binding:"required"`
	CategoryId    int64   `json:"category_id" form:"category_id" validate:"required" binding:"required"`
	StoreId       int64   `json:"store_id" form:"store_id" validate:"required" binding:"required"`
}
