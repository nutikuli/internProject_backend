package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	_prodCate "github.com/nutikuli/internProject_backend/internal/models/product-category/repository"
	"github.com/nutikuli/internProject_backend/internal/models/product/repository"
	"github.com/nutikuli/internProject_backend/internal/models/product/usecase"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
	_fileUse "github.com/nutikuli/internProject_backend/internal/services/file/usecase"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"
)

func UseProductRoute(db *sqlx.DB, app fiber.Router) {
	prodR := app.Group("/product", func(c *fiber.Ctx) error {
		log.Infof("product : %v", c.Request().URI().String())
		return c.Next()
	})

	fileRepo := _fileRepo.NewFileRepository(db)
	fileUse := _fileUse.NewFileUsecase(fileRepo)

	prodCateRepo := _prodCate.NewProductCategoryRepository(db)

	prodRepo := repository.NewproductRepository(db)
	prodUse := usecase.NewProductUsecase(prodRepo, fileRepo, fileUse, prodCateRepo)

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	prodConn := NewProductHandler(prodUse)

	prodR.Post("/create-product",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "สินค้า",
				Action: "เพิ่มข้อมูล",
			}

			return logger.LogRequest(c, logAction)
		},
		prodConn.CreateProduct)
	prodR.Get("/get-product-id/:product_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "สินค้า",
				Action: fmt.Sprintf("ดูข้อมูล Product หมายเลข  %s", c.Params("product_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		prodConn.GetProductById)
	prodR.Get("/get-products-by-store-id/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "สินค้า",
				Action: fmt.Sprintf("ดูข้อมูล Stores หมายเลข  %s", c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		prodConn.GetProductsByStoreId)
	prodR.Get("/get-products-by-order-id/:order_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "สินค้า",
				Action: fmt.Sprintf("ดูข้อมูล Orders หมายเลข  %s", c.Params("order_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		prodConn.GetProductsByOrderId)
	prodR.Delete("/delete-product-id/:product_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "สินค้า",
				Action: fmt.Sprintf("ลบข้อมูล Product หมายเลข  %s", c.Params("product_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		prodConn.DeleteProductById)
	prodR.Patch("/update-product-id/:product_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "สินค้า",
				Action: fmt.Sprintf("อัปเดตข้อมูล Product หมายเลข  %s", c.Params("product_id")),
			}

			return logger.LogRequest(c, logAction)

		},
		prodConn.UpdateProductById)
	prodR.Get("/get-products",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "สินค้า",
				Action: "ดูข้อมูล",
			}

			return logger.LogRequest(c, logAction)
		},
		prodConn.GetAllProducts)
}
