package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	"github.com/nutikuli/internProject_backend/internal/models/product-category/repository"
	"github.com/nutikuli/internProject_backend/internal/models/product-category/usecase"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"
)

func UseProductCategoryRoute(db *sqlx.DB, app fiber.Router) {
	prodCatR := app.Group("/product-category", func(c *fiber.Ctx) error {
		log.Infof("product-category : %v", c.Request().URI().String())
		return c.Next()
	})

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	prodCatRepo := repository.NewProductCategoryRepository(db)
	prodCatUse := usecase.NewProductCategoryUsecase(prodCatRepo)

	prodCatConn := NewProductCategoryHandler(prodCatUse)

	prodCatR.Post("/create-product-category/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "หมวดหมู่สินค้า",
				Action: "เพิ่มข้อมูล Product Category ใน Store หมายเลข " + c.Params("store_id"),
			}

			return logger.LogRequest(c, logAction)
		},
		prodCatConn.CreateProductCategory)
	prodCatR.Get("/get-product-category-id/:product_category_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "หมวดหมู่สินค้า",
				Action: fmt.Sprintf("ดูข้อมูล Product Category หมายเลข  %s", c.Params("product_category_id")),
			}

			return logger.LogRequest(c, logAction)

		},
		prodCatConn.GetProductCategoryById)
	prodCatR.Get("/get-product-categories-by-store-id/:store_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ผู้ขาย",
				Action: fmt.Sprintf("ดูข้อมูล Product Category หมายเลข  %s", c.Params("store_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		prodCatConn.GetProductCategoriesByStoreId)
	prodCatR.Delete("/delete-product-category-id/:product_category_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "หมวดหมู่สินค้า",
				Action: fmt.Sprintf("ลบข้อมูล Product Category หมายเลข  %s", c.Params("product_category_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		prodCatConn.DeleteProductCategoryById)
	prodCatR.Patch("/update-product-category-id/:product_category_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "หมวดหมู่สินค้า",
				Action: fmt.Sprintf("อัปเดตข้อมูล Product Category หมายเลข  %s", c.Params("product_category_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		prodCatConn.UpdateProductCategoryById)
}
