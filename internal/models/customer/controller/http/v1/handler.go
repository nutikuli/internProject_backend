package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	"github.com/nutikuli/internProject_backend/internal/models/customer/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	"github.com/nutikuli/internProject_backend/pkg/utils"
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
	req, err := c.ParamsInt("customer_id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request customer_id param",
			"result":      nil,
		})
	}

	req64 := int64(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	customer, status, err := cus.CustomerUse.OnGetCustomerById(ctx, req64)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
		"message":     nil,
		"result":      customer,
	})
}

func (con *customerConn) CreateCustomerAccount(c *fiber.Ctx) error {
	req := new(dtos.CustomerFileReq)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	_, errOnValidate := utils.SchemaValidator(req)
	if errOnValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     errOnValidate.Error(),
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	customer, _, status, err := con.CustomerUse.OnCreateCustomerAccount(c, ctx, req.CustomerRegisterData)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
		"message":     "",
		"result":      customer,
	})
}

func (con *customerConn) UpdateCustomerById(c *fiber.Ctx) error {

	var req = new(entities.CustomerUpdateReq)
	if cE := c.BodyParser(req); cE != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": cE.Error(),
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
		reqP        = c.Params("user_id")
	)
	defer cancel()

	if reqP == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "user_id params is required",
			"raw_message": "",
			"result":      nil,
		})
	}

	userId, err := strconv.ParseInt(reqP, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"raw_message": err.Error(),
			"message":     "error, invalid user_id params",
			"result":      nil,
		})
	}

	status, cE := con.CustomerUse.OnUpdateCustomerById(ctx, userId, req)
	if cE != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     "",
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "user updated successfully",
		"result":      nil,
	})
}

func (con *customerConn) DeletedCustomerByID(c *fiber.Ctx) error {
	req, err := c.ParamsInt("store_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
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
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusNotFound),
			"status_code": http.StatusNotFound,
			"message":     "Customer not found",
			"result":      nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"status_code": http.StatusOK,
		"message":     nil,
		"result":      customer,
	})
}
