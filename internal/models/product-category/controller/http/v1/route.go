package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/product-category/repository"
	"github.com/nutikuli/internProject_backend/internal/models/product-category/usecase"
)

func UseProductCategoryRoute(db *sqlx.DB, app fiber.Router) {
	prodCatR := app.Group("/product-category", func(c *fiber.Ctx) error {
		log.Infof("product-category : %v", c.Request().URI().String())
		return c.Next()
	})

	prodCatRepo := repository.NewProductCategoryRepository(db)
	prodCatUse := usecase.NewProductCategoryUsecase(prodCatRepo)

	prodCatConn := NewProductCategoryHandler(prodCatUse)

	prodCatR.Post("/create-product-category/:store_id", prodCatConn.CreateProductCategory)
	prodCatR.Get("/get-product-category-id/:product_category_id", prodCatConn.GetProductCategoryById)
	prodCatR.Get("/get-product-categories-by-store-id/:store_id", prodCatConn.GetProductCategoriesByStoreId)
	prodCatR.Delete("/delete-product-category-id/:product_category_id", prodCatConn.DeleteProductCategoryById)
	prodCatR.Patch("/update-product-category-id/:product_category_id", prodCatConn.UpdateProductCategoryById)
}
