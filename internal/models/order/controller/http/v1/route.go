package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_bankRepo "github.com/nutikuli/internProject_backend/internal/models/bank/repository"
	_bankUse "github.com/nutikuli/internProject_backend/internal/models/bank/usecase"
	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	_orderProdRepo "github.com/nutikuli/internProject_backend/internal/models/order-product/repository"
	_orderProdUse "github.com/nutikuli/internProject_backend/internal/models/order-product/usecase"
	"github.com/nutikuli/internProject_backend/internal/models/order/repository"
	"github.com/nutikuli/internProject_backend/internal/models/order/usecase"
	_prodUse "github.com/nutikuli/internProject_backend/internal/models/product/usecase"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"

	_prodRepo "github.com/nutikuli/internProject_backend/internal/models/product/repository"

	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
	_fileUse "github.com/nutikuli/internProject_backend/internal/services/file/usecase"
)

func UseOrderRoute(db *sqlx.DB, app fiber.Router) {
	orderR := app.Group("/order", func(c *fiber.Ctx) error {
		log.Infof("order : %v", c.Request().URI().String())
		return c.Next()
	})

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	fileRepo := _fileRepo.NewFileRepository(db)
	fileUse := _fileUse.NewFileUsecase(fileRepo)

	prodRepo := _prodRepo.NewproductRepository(db)
	prodUse := _prodUse.NewProductUsecase(prodRepo, fileRepo, fileUse)

	orderProdRepo := _orderProdRepo.NewOrderProductRepository(db)
	orderProdUse := _orderProdUse.NewOrderProductUsecase(orderProdRepo)

	bankRepo := _bankRepo.NewBankRepository(db)
	bankUse := _bankUse.NewBankUsecase(bankRepo, fileRepo)

	orderRepo := repository.NewOrderRepository(db)
	orderUse := usecase.NewOrderUsecase(orderRepo, fileRepo, bankUse, prodUse, orderProdUse)

	orderConn := NewOrderHandler(orderUse)

	orderR.Post("/create-order",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: "สร้างรายการสั่งซื้อ",
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.CreateOrder)
	orderR.Get("/get-store-order-by-id/:store_id/:order_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: fmt.Sprintf("ดูข้อมูล Order ของหมายเลข %s ใน Store เลขหมายเลข  %s", c.Params("order_id"), c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.GetOrderById)
	orderR.Get("/get-orders-by-store-id/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: fmt.Sprintf("ดูข้อมูล Order ของ Store เลขหมายเลข  %s", c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.GetOrdersByStoreId)
	orderR.Get("/get-orders-by-customer-id/:customer_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: fmt.Sprintf("ดูข้อมูล Order ของ Customer เลขหมายเลข  %s", c.Params("customer_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.GetOrdersByCustomerId)
	orderR.Patch("/update-order-state/:order_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "รายการสั่งซื้อ",
				Action: fmt.Sprintf("อัปเดตสถานะ Order หมายเลข  %s", c.Params("order_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		orderConn.UpdateOrderTransportDetailAndState)
}
