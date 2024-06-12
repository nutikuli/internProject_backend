package dtos

import (
	_orderProductEntities "github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
	_productCateEntities "github.com/nutikuli/internProject_backend/internal/models/product-category/entities"
	_productEntities "github.com/nutikuli/internProject_backend/internal/models/product/entities"
	_filesEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type ProductFileRes struct {
	Files                                 []*_filesEntities.File `json:"files_data" `
	*_productEntities.Product             `json:"product_data" `
	*_productCateEntities.ProductCategory `json:"product_category_data"`
}

type ProductWithOrderProductFileRes struct {
	Files                                 []*_filesEntities.File `json:"files_data" `
	*_productEntities.Product             `json:"product_data" `
	*_productCateEntities.ProductCategory `json:"product_category_data"`
	OrderProductData                      *_orderProductEntities.OrderProduct `json:"order_product_data"`
}
