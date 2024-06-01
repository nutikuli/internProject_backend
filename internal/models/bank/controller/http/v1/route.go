package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/bank/repository"
	"github.com/nutikuli/internProject_backend/internal/models/bank/usecase"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
)

func UseBankRoute(db *sqlx.DB, app fiber.Router) {
	bankR := app.Group("/bank", func(c *fiber.Ctx) error {
		log.Infof("bank : %v", c.Request().URI().String())
		return c.Next()
	})

	fileRepo := _fileRepo.NewFileRepository(db)

	bankRepo := repository.NewBankRepository(db)
	bankUse := usecase.NewBankUsecase(bankRepo, fileRepo)

	bankConn := NewBankHandler(bankUse)

	bankR.Post("/create-bank", bankConn.CreateBank)
	bankR.Get("/get-bank-id/:bank_id", bankConn.GetBankById)
	bankR.Get("/get-banks-by-store-id/:store_id", bankConn.GetBanksByStoreId)
	bankR.Delete("/delete-bank-id/:bank_id", bankConn.DeleteBankById)
	bankR.Patch("/update-bank-id/:bank_id", bankConn.UpdateBankById)
}
