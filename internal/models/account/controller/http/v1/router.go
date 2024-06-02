package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accRepo "github.com/nutikuli/internProject_backend/internal/models/account/repository"
	_accUse "github.com/nutikuli/internProject_backend/internal/models/account/usecase"
	_adminRepo "github.com/nutikuli/internProject_backend/internal/models/admin/repository"
	"github.com/nutikuli/internProject_backend/internal/models/customer/repository"
	_storeRepo "github.com/nutikuli/internProject_backend/internal/models/store/repository"

	_cutomerHand "github.com/nutikuli/internProject_backend/internal/models/customer/controller/http/v1"
	_cutomerUse "github.com/nutikuli/internProject_backend/internal/models/customer/usecase"
)

func UseAccountRoute(db *sqlx.DB, app fiber.Router) {
	authR := app.Group("/account", func(c *fiber.Ctx) error {
		log.Infof("store : %v", c.Request().URI().String())
		return c.Next()
	})

	//register
	adminRepo := _adminRepo.NewFileRepository(db)
	storeRep := _storeRepo.NewStoreRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, nil, adminRepo, customerRepo, storeRep)
	customerUse := _cutomerUse.NewCustomerUsecase(customerRepo, accUse)
	customerConn := _cutomerHand.NewCustomerHandler(customerUse)
	authR.Post("/register", customerConn.CreateCustomerAccount)

	//login
	accConn := NewAccountHandler(accUse)
	authR.Post("/login", accConn.Login)

	//OTP
	authR.Post("/otp", accConn.OTP)
	//resetPassword
	authR.Post("/resetpass", accConn.UpdatePass)
}
