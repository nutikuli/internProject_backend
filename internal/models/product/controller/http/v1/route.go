package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/product/repository"
	"github.com/nutikuli/internProject_backend/internal/models/product/usecase"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
)

func UseProductRoute(db *sqlx.DB, app fiber.Router) {
	prodR := app.Group("/product", func(c *fiber.Ctx) error {
		log.Infof("product : %v", c.Request().URI().String())
		return c.Next()
	})

	fileRepo := _fileRepo.NewFileRepository(db)

	prodRepo := repository.NewproductRepository(db)
	prodUse := usecase.NewProductUsecase(prodRepo, fileRepo)

	prodConn := NewProductHandler(prodUse)

	prodR.Post("/create-product", prodConn.CreateProduct)
	prodR.Get("/get-product-id/:product_id", prodConn.GetProductById)
	prodR.Get("/get-products-by-store-id/:store_id", prodConn.GetProductsByStoreId)
	prodR.Get("/get-products-by-order-id/:order_id", prodConn.GetProductsByOrderId)
	prodR.Delete("/delete-product-id/:product_id", prodConn.DeleteProductById)
	prodR.Patch("/update-product-id/:product_id", prodConn.UpdateProductById)

}
