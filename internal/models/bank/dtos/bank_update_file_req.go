package dtos

import (
	_bankEntities "github.com/nutikuli/internProject_backend/internal/models/bank/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type BankUpdateFileReq struct {
	BankData  *_bankEntities.BankUpdateReq     `json:"bank_data" form:"bank_data" binding:"required"`
	FilesData []*_fileEntities.FileUploaderReq `json:"files_data" form:"files_data" binding:"required"`
}
