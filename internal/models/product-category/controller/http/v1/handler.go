package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	product_category "github.com/nutikuli/internProject_backend/internal/models/product-category"
	"github.com/nutikuli/internProject_backend/internal/models/product-category/entities"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type prodCateConn struct {
	prodCateUse product_category.ProductCategoryUsecase
}

func NewProductCategoryHandler(
	prodCateUse product_category.ProductCategoryUsecase,

) *prodCateConn {
	return &prodCateConn{
		prodCateUse: prodCateUse,
	}
}

func (p *prodCateConn) CreateProductCategory(c *fiber.Ctx) error {
	req := new(entities.ProductCategoryCreatedReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"result":      nil,
		})
	}

	_, errOnValidate := utils.SchemaValidator(req)
	if errOnValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     errOnValidate,
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res, status, err := p.prodCateUse.OnCreateProductCategory(ctx, req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     nil,
		"result":      res,
	})

}

func (p *prodCateConn) GetProductCategoryById(c *fiber.Ctx) error {
	categoryId, err := c.ParamsInt("product_category_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid category id",
			"result":      nil,
		})
	}

	categoryIdInt64 := int64(categoryId)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res, status, err := p.prodCateUse.OnGetProductCategoryById(ctx, categoryIdInt64)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     nil,
		"result":      res,
	})
}

func (p *prodCateConn) GetProductCategoriesByStoreId(c *fiber.Ctx) error {
	storeId, err := c.ParamsInt("store_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid store id",
			"result":      nil,
		})
	}

	storeIdInt64 := int64(storeId)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res, status, err := p.prodCateUse.OnGetProductCategoriesByStoreId(ctx, storeIdInt64)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     nil,
		"result":      res,
	})
}

func (p *prodCateConn) DeleteProductCategoryById(c *fiber.Ctx) error {
	categoryId, err := c.ParamsInt("product_category_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid category id",
			"result":      nil,
		})
	}

	categoryIdInt64 := int64(categoryId)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	status, err := p.prodCateUse.OnDeleteProductCategoryById(ctx, categoryIdInt64)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     nil,
		"result":      nil,
	})
}

func (p *prodCateConn) UpdateProductCategoryById(c *fiber.Ctx) error {
	categoryId, err := c.ParamsInt("product_category_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid category id",
			"result":      nil,
		})
	}

	categoryIdInt64 := int64(categoryId)

	req := new(entities.ProductCategoryCreatedReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"result":      nil,
		})
	}

	_, errOnValidate := utils.SchemaValidator(req)
	if errOnValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     errOnValidate,
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	status, err := p.prodCateUse.OnUpdateProductCategoryById(ctx, categoryIdInt64, req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     nil,
		"result":      nil,
	})
}
