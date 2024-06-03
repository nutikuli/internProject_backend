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

func UseCustomerRoute(db *sqlx.DB, app fiber.Router) {
	authR := app.Group("/customer", func(c *fiber.Ctx) error {
		log.Infof("customer : %v", c.Request().URI().String())
		return c.Next()
	})

	customerRepo := repository.NewCustomerRepository(db)

	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, nil, nil, customerRepo, nil)

	customerUse := _customerUse.NewCustomerUsecase(customerRepo, accUse)

	customerConn := NewCustomerHandler(customerUse)

	authR.Post("/account-register", customerConn.CreateCustomerAccount)
	authR.Get("/:customer_id", customerConn.GetCustomerById)
	authR.Patch("/:customer_id", customerConn.UpdateCustomerById)
	authR.Delete("/:customer_id", customerConn.DeletedCustomerByID)
}
