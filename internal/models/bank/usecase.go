package bank

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_bankDtos "github.com/nutikuli/internProject_backend/internal/models/bank/dtos"
	_bankEntities "github.com/nutikuli/internProject_backend/internal/models/bank/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type BankUseCase interface {
	OnCreateBank(c *fiber.Ctx, ctx context.Context, bankDatReq *_bankEntities.BankCreatedReq, filesDatReq []*_fileEntities.FileUploaderReq) (*_bankDtos.BankFileRes, int, error)
	OnGetBanksByStoreId(ctx context.Context, bankId *int64) ([]*_bankDtos.BankFileRes, int, error)
}
