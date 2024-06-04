package v1

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/product"
	"github.com/nutikuli/internProject_backend/internal/models/product/dtos"
)

type prodConn struct {
	prodUse product.ProductUsecase
}

func NewProductHandler(prodUse product.ProductUsecase) *prodConn {
	return &prodConn{
		prodUse: prodUse,
	}
}

func (p *prodConn) GetAllProducts(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)
	defer cancel()

	products, status, err := p.prodUse.OnGetAllProducts(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      fiber.StatusOK,
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      products,
	})
}

func (p *prodConn) GetProductsByStoreId(c *fiber.Ctx) error {
	req, err := c.ParamsInt("store_id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request store_id param",
			"result":      nil,
		})
	}

	req64 := int64(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	products, status, err := p.prodUse.OnGetProductsByStoreId(ctx, req64)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      fiber.StatusOK,
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      products,
	})
}

func (p *prodConn) GetProductById(c *fiber.Ctx) error {
	req, err := c.ParamsInt("product_id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request product_id param",
			"result":      nil,
		})
	}

	req64 := int64(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	product, status, err := p.prodUse.OnGetProductById(ctx, req64)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      fiber.StatusOK,
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      product,
	})
}

func (p *prodConn) GetProductsByOrderId(c *fiber.Ctx) error {
	req, err := c.ParamsInt("order_id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request order_id param",
			"result":      nil,
		})
	}

	req64 := int64(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	products, status, err := p.prodUse.OnGetProductsByOrderId(ctx, req64)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      fiber.StatusOK,
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      products,
	})
}

func (p *prodConn) CreateProduct(c *fiber.Ctx) error {
	req := new(dtos.ProductFileReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request body",
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	product, status, err := p.prodUse.OnCreateProduct(c, ctx, req.ProductData, req.FilesData)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":      fiber.StatusCreated,
		"status_code": fiber.StatusCreated,
		"message":     nil,
		"result":      product,
	})
}

func (p *prodConn) DeleteProductById(c *fiber.Ctx) error {
	req, err := c.ParamsInt("product_id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request product_id param",
			"result":      nil,
		})
	}

	req64 := int64(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	status, err := p.prodUse.OnDeleteProductById(ctx, req64)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      fiber.StatusOK,
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      nil,
	})
}

func (p *prodConn) UpdateProductById(c *fiber.Ctx) error {
	prodId, err := c.ParamsInt("product_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request product_id param",
			"result":      nil,
		})
	}

	prodId64 := int64(prodId)

	req := new(dtos.ProductFileUpdateReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request body",
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	product, status, err := p.prodUse.OnUpdateProductById(c, ctx, prodId64, req.ProductData, req.FilesData)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      fiber.StatusOK,
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      product,
	})
}
