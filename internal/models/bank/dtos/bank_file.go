package entities

import (
	_bankEntities "github.com/nutikuli/internProject_backend/internal/models/bank/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type BankFileRes struct {
	BankData  *_bankEntities.Bank   `json:"bank_data"`
	FilesData []*_fileEntities.File `json:"files_data"`
}
