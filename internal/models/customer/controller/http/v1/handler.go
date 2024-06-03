package v1

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	"github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
)

type customerConn struct {
	CustomerUse customer.CustomerUsecase
}

func NewCustomerHandler(CustomerUse customer.CustomerUsecase) *customerConn {
	return &customerConn{
		CustomerUse: CustomerUse,
	}
}

func (cus *customerConn) GetCustomerById(c *fiber.Ctx) error {
	req, err := strconv.ParseInt(c.Params("customer_id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request customer_id param",
			"result":      nil,
		})
	}

	log.Debug(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	customer, status, err := cus.CustomerUse.OnGetCustomerById(ctx, req)
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
		"result":      customer,
	})
}

func (con *customerConn) CreateCustomerAccount(c *fiber.Ctx) error {
	req := new(entities.CustomerRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}
	log.Debug(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	customerRes, userToken, status, err := con.CustomerUse.OnCreateCustomerAccount(c, ctx, req)
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
		"result": dtos.CustomerTokenRes{
			Customer: customerRes,
			Token:    userToken,
		},
	})
}

func (con *customerConn) UpdateCustomerById(c *fiber.Ctx) error {

	var req = new(entities.CustomerUpdateReq)
	if cE := c.BodyParser(req); cE != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": cE.Error(),
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
		reqP        = c.Params("customer_id")
	)
	defer cancel()

	if reqP == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "user_id params is required",
			"raw_message": "",
			"result":      nil,
		})
	}

	customerId, err := strconv.ParseInt(reqP, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"raw_message": err.Error(),
			"message":     "error, invalid user_id params",
			"result":      nil,
		})
	}

	status, cE := con.CustomerUse.OnUpdateCustomerById(ctx, customerId, req)
	if cE != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     "",
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      status,
		"status_code": status,
		"message":     "user updated successfully",
		"result":      nil,
	})
}

func (con *customerConn) DeletedCustomerByID(c *fiber.Ctx) error {
	req, err := c.ParamsInt("customer_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request id param",
			"result":      nil,
		})
	}

	req64 := int64(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	customer, err := con.CustomerUse.OnDeletedCustomer(ctx, req64)
	if err != nil {
		// Assuming OnGetDeletedCustomerByID returns an error if customer not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      fiber.StatusNotFound,
			"status_code": fiber.StatusNotFound,
			"message":     "Customer not found",
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      fiber.StatusOK,
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      customer,
	})
}
