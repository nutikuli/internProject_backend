package entities

type ProductCategoryCreatedReq struct {
	Name   string `json:"name" form:"name" validate:"required" binding:"required"`
	Status bool   `json:"status" form:"status" validate:"required" binding:"required"`
	Code   string `json:"code" form:"code"`
	Detail string `json:"detail" form:"detail" validate:"required" binding:"required"`
}
