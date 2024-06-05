package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accRepo "github.com/nutikuli/internProject_backend/internal/models/account/repository"
	_accUse "github.com/nutikuli/internProject_backend/internal/models/account/usecase"
	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	"github.com/nutikuli/internProject_backend/internal/models/store/repository"
	_storeUse "github.com/nutikuli/internProject_backend/internal/models/store/usecase"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"
)

func UseStoreRoute(db *sqlx.DB, app fiber.Router) {
	storeR := app.Group("/store", func(c *fiber.Ctx) error {
		log.Infof("store : %v", c.Request().URI().String())
		return c.Next()
	})

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	fileRepo := _fileRepo.NewFileRepository(db)

	storeRepo := repository.NewStoreRepository(db)

	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, fileRepo, nil, nil, storeRepo)

	storeUse := _storeUse.NewStoreUsecase(storeRepo, fileRepo, accUse)

	storeConn := NewStoreHandler(storeUse)

	storeR.Get("/get-stores", func(c *fiber.Ctx) error {
		logAction := &middlewares.LoggerAction{
			Menu:   "สินค้า",
			Action: "ดูข้อมูล",
		}

		return logger.LogRequest(c, logAction)
	}, storeConn.GetStoreAccounts)
	storeR.Post("/account-register", func(c *fiber.Ctx) error {
		logAction := &middlewares.LoggerAction{
			Menu:   "ผู้ขาย",
			Action: "เพิ่มข้อมูล",
		}

		return logger.LogRequest(c, logAction)
	}, storeConn.RegisterStoreAccount)
	storeR.Get("/get-store-by-id/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ผู้ขาย",
				Action: fmt.Sprintf("ดูข้อมูล Store Account หมายเลข  %s", c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		storeConn.GetStoreById)
	storeR.Patch("/update-store-by-id/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ผู้ขาย",
				Action: fmt.Sprintf("อัปเดตข้อมูล Store Account หมายเลข  %s", c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		storeConn.UpdateStoreById)
	storeR.Delete("/delete-store-by-id/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ผู้ขาย",
				Action: fmt.Sprintf("ลบข้อมูล Store Account หมายเลข  %s", c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		storeConn.DeleteStoreById)
}
