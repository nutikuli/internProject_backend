package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/bank/repository"
	"github.com/nutikuli/internProject_backend/internal/models/bank/usecase"
	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"
)

func UseBankRoute(db *sqlx.DB, app fiber.Router) {
	bankR := app.Group("/bank", func(c *fiber.Ctx) error {
		log.Infof("bank : %v", c.Request().URI().String())
		return c.Next()
	})

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	fileRepo := _fileRepo.NewFileRepository(db)

	bankRepo := repository.NewBankRepository(db)
	bankUse := usecase.NewBankUsecase(bankRepo, fileRepo)

	bankConn := NewBankHandler(bankUse)

	bankR.Post("/create-bank",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ธนาคาร",
				Action: "เพิ่มข้อมูลธนาคาร",
			}

			return logger.LogRequest(c, logAction)
		},
		bankConn.CreateBank)
	bankR.Get("/get-bank-id/:bank_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ธนาคาร",
				Action: "ดูข้อมูลธนาคาร หมายเลข " + c.Params("bank_id"),
			}

			return logger.LogRequest(c, logAction)
		},
		bankConn.GetBankById)
	bankR.Get("/get-banks-by-store-id/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ธนาคาร",
				Action: "ดูข้อมูลธนาคาร ใน Store หมายเลข " + c.Params("store_id"),
			}

			return logger.LogRequest(c, logAction)
		},
		bankConn.GetBanksByStoreId)
	bankR.Delete("/delete-bank-id/:bank_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ธนาคาร",
				Action: "ลบข้อมูลธนาคาร หมายเลข " + c.Params("bank_id"),
			}

			return logger.LogRequest(c, logAction)
		},
		bankConn.DeleteBankById)
	bankR.Patch("/update-bank-id/:bank_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ธนาคาร",
				Action: "อัปเดตข้อมูลธนาคาร หมายเลข " + c.Params("bank_id"),
			}

			return logger.LogRequest(c, logAction)
		},
		bankConn.UpdateBankById)
}
