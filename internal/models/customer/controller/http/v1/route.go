package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accRepo "github.com/nutikuli/internProject_backend/internal/models/account/repository"
	_accUse "github.com/nutikuli/internProject_backend/internal/models/account/usecase"
	"github.com/nutikuli/internProject_backend/internal/models/customer/repository"
	_customerUse "github.com/nutikuli/internProject_backend/internal/models/customer/usecase"
	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"
)

func UseCustomerRoute(db *sqlx.DB, app fiber.Router) {
	authR := app.Group("/customer", func(c *fiber.Ctx) error {
		log.Infof("customer : %v", c.Request().URI().String())
		return c.Next()
	})

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	customerRepo := repository.NewCustomerRepository(db)

	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, nil, nil, customerRepo, nil)

	customerUse := _customerUse.NewCustomerUsecase(customerRepo, accUse)

	customerConn := NewCustomerHandler(customerUse)

	authR.Post("/account-register",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ลูกค้า",
				Action: "เพิ่มข้อมูล",
			}

			return logger.LogRequest(c, logAction)
		},

		customerConn.CreateCustomerAccount)
	authR.Get("/getallcustomer",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ลูกค้า",
				Action: "ดูข้อมูลทั้งหมด",
			}

			return logger.LogRequest(c, logAction)
		},
		customerConn.GetAllCustomer)
	authR.Get("/:customer_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ลูกค้า",
				Action: "ดูข้อมูล Customer Account หมายเลข " + c.Params("customer_id"),
			}

			return logger.LogRequest(c, logAction)
		},
		customerConn.GetCustomerById)
	authR.Patch("/:customer_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ลูกค้า",
				Action: "อัปเดตข้อมูล Customer Account หมายเลข " + c.Params("customer_id"),
			}

			return logger.LogRequest(c, logAction)
		},
		customerConn.UpdateCustomerById)
	authR.Delete("/:customer_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ลูกค้า",
				Action: "ลบข้อมูล Customer Account หมายเลข " + c.Params("customer_id"),
			}

			return logger.LogRequest(c, logAction)
		},
		customerConn.DeletedCustomerByID)

}
