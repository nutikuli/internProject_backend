package dtos

import (
	_orderProductEntities "github.com/nutikuli/internProject_backend/internal/models/order-product/entities"
	_orderEntities "github.com/nutikuli/internProject_backend/internal/models/order/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type OrderFileBankIdOrderProductsReq struct {
	BankId            int64                                          `json:"bank_id" form:"bank_id" binding:"required"`
	FilesData         []*_fileEntities.FileUploaderReq               `json:"files_data" form:"files_data" binding:"required"`
	OrderData         *_orderEntities.OrderCreate                    `db:"order_data" form:"order_data" binding:"required"`
	OrderProductsData []*_orderProductEntities.OrderProductCreateReq `json:"order_product" form:"order_product" binding:"required"`
}
