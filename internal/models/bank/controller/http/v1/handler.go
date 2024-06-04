package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/bank"
	"github.com/nutikuli/internProject_backend/internal/models/bank/dtos"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type bankConn struct {
	BankUse bank.BankUseCase
}

func NewBankHandler(BankUse bank.BankUseCase) *bankConn {
	return &bankConn{
		BankUse: BankUse,
	}
}

func (o *bankConn) GetBanksByStoreId(c *fiber.Ctx) error {
	req, err := strconv.ParseInt(c.Params("store_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request id param",
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	customer, status, err := o.BankUse.OnGetBanksByStoreId(ctx, req)
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

func (o *bankConn) CreateBank(c *fiber.Ctx) error {
	req := new(dtos.BankFileReq)
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

	// Call OnCreateBank with all four arguments
	bank, status, err := o.BankUse.OnCreateBank(c, ctx, req.BankData, req.FilesData)
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
		"result":      bank,
	})
}

func (o *bankConn) GetBankById(c *fiber.Ctx) error {
	req, err := strconv.ParseInt(c.Params("bank_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request id param",
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	bank, status, err := o.BankUse.OnGetBankByBankId(ctx, req)
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
		"result":      bank,
	})
}

func (o *bankConn) UpdateBankById(c *fiber.Ctx) error {
	req := new(dtos.BankFileReq)
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

	req64, err := strconv.ParseInt(c.Params("bank_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request id param",
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	bank, status, err := o.BankUse.OnUpdateBankById(c, ctx, req64, req.BankData, req.FilesData)
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
		"result":      bank,
	})
}

func (o *bankConn) DeleteBankById(c *fiber.Ctx) error {
	req, err := strconv.ParseInt(c.Params("bank_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request id param",
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	status, err := o.BankUse.OnDeleteBankById(ctx, req)
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
		"result":      nil,
	})
}
