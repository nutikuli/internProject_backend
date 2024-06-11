package dtos

import (
	_bankDtos "github.com/nutikuli/internProject_backend/internal/models/bank/dtos"
	_customerEntities "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	_orderEntities "github.com/nutikuli/internProject_backend/internal/models/order/entities"
	_productDtos "github.com/nutikuli/internProject_backend/internal/models/product/dtos"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type OrderBankFilesRes struct {
	OrderData         *_orderEntities.Order          `json:"order_data"`
	CustomerData      *_customerEntities.Customer    `json:"customer_data"`
	FilesData         []*_fileEntities.File          `json:"files_data"`
	BanksData         *_bankDtos.BankFileRes         `json:"bank_payment"`
	OrdersProductData []*_productDtos.ProductFileRes `json:"product_data"`
}

type OrderBanksFilesRes struct {
	OrderData    *_orderEntities.Order       `json:"order_data"`
	CustomerData *_customerEntities.Customer `json:"customer_data"`
}
