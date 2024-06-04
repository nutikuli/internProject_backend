package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accRepo "github.com/nutikuli/internProject_backend/internal/models/account/repository"
	_accUse "github.com/nutikuli/internProject_backend/internal/models/account/usecase"
	"github.com/nutikuli/internProject_backend/internal/models/store/repository"
	_storeUse "github.com/nutikuli/internProject_backend/internal/models/store/usecase"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
)

func UseStoreRoute(db *sqlx.DB, app fiber.Router) {
	storeR := app.Group("/store", func(c *fiber.Ctx) error {
		log.Infof("store : %v", c.Request().URI().String())
		return c.Next()
	})

	fileRepo := _fileRepo.NewFileRepository(db)

	storeRepo := repository.NewStoreRepository(db)

	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, fileRepo, nil, nil, storeRepo)

	storeUse := _storeUse.NewStoreUsecase(storeRepo, fileRepo, accUse)

	storeConn := NewStoreHandler(storeUse)

	storeR.Post("/account-register", storeConn.RegisterStoreAccount)
	storeR.Get("/get-store-by-id/:store_id", storeConn.GetStoreById)
	storeR.Patch("/update-store-by-id/:store_id", storeConn.UpdateStoreById)
	storeR.Delete("/delete-store-by-id/:store_id", storeConn.DeleteStoreById)
}
