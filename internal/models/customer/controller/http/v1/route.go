package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accRepo "github.com/nutikuli/internProject_backend/internal/models/account/repository"
	_accUse "github.com/nutikuli/internProject_backend/internal/models/account/usecase"
	"github.com/nutikuli/internProject_backend/internal/models/customer/repository"
	_customerUse "github.com/nutikuli/internProject_backend/internal/models/customer/usecase"
)

func UseCustomerRoute(db *sqlx.DB, app *fiber.App) {
	authR := app.Group("/customer", func(c *fiber.Ctx) error {
		log.Infof("all : %v", c.Request().URI().String())
		return c.Next()
	})

	customerRepo := repository.NewCustomerRepository(db)

	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, nil, nil, customerRepo, nil)

	customerUse := _customerUse.NewCustomerUsecase(customerRepo, accUse)

	customerConn := NewCustomerHandler(customerUse)

	authR.Get("/:customer_id", customerConn.GetCustomerById)
	authR.Post("/account-register", customerConn.CreateCustomerAccount)
	authR.Put("/update_id", customerConn.UpdateCustomerById)
	authR.Delete("/delete_id", customerConn.DeletedCustomerByID)
}
