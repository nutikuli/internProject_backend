package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_bankRepo "github.com/nutikuli/internProject_backend/internal/models/bank/repository"
	_bankUse "github.com/nutikuli/internProject_backend/internal/models/bank/usecase"
	_orderProdRepo "github.com/nutikuli/internProject_backend/internal/models/order-product/repository"
	_orderProdUse "github.com/nutikuli/internProject_backend/internal/models/order-product/usecase"
	"github.com/nutikuli/internProject_backend/internal/models/order/repository"
	"github.com/nutikuli/internProject_backend/internal/models/order/usecase"
	_prodUse "github.com/nutikuli/internProject_backend/internal/models/product/usecase"

	_prodRepo "github.com/nutikuli/internProject_backend/internal/models/product/repository"

	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
)

func UseOrderRoute(db *sqlx.DB, app fiber.Router) {
	orderR := app.Group("/order", func(c *fiber.Ctx) error {
		log.Infof("order : %v", c.Request().URI().String())
		return c.Next()
	})

	fileRepo := _fileRepo.NewFileRepository(db)

	prodRepo := _prodRepo.NewproductRepository(db)
	prodUse := _prodUse.NewProductUsecase(prodRepo, fileRepo)

	orderProdRepo := _orderProdRepo.NewOrderProductRepository(db)
	orderProdUse := _orderProdUse.NewOrderProductUsecase(orderProdRepo)

	bankRepo := _bankRepo.NewBankRepository(db)
	bankUse := _bankUse.NewBankUsecase(bankRepo, fileRepo)

	orderRepo := repository.NewOrderRepository(db)
	orderUse := usecase.NewOrderUsecase(orderRepo, fileRepo, bankUse, prodUse, orderProdUse)

	orderConn := NewOrderHandler(orderUse)

	orderR.Post("/create-order", orderConn.CreateOrder)
	orderR.Get("/get-order-id/:order_id", orderConn.GetOrderById)
	orderR.Get("/get-orders-by-store-id/:store_id", orderConn.GetOrdersByStoreId)
	orderR.Get("/get-orders-by-customer-id/:customer_id", orderConn.GetOrdersByCustomerId)
	orderR.Patch("/update-order-state/:order_id", orderConn.UpdateOrderTransportDetailAndState)

}
