package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	_accEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
)

type accountConn struct {
	AccountUse account.AccountUsecase
}

func NewOrderHandler(accountUse account.AccountUsecase) *accountConn {
	return &accountConn{
		AccountUse: accountUse,
	}
}


func (a *accountConn) Login(c *fiber.Ctx) error {
	req := new(_accEntities.UsersCredential)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}


	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res,nil, status, err := a.AccountUse.Login(ctx, req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      res,
	})
}



func (a *accountConn) OTP(c *fiber.Ctx) error {
	req := new(_accEntities.UsersCredential)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}


	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res, status, err := a.AccountUse.CheckOTP(c,ctx,req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      res,
	})
}

func (a *accountConn) UpdatePass(c *fiber.Ctx) error {
	req := new(_accEntities.UsersCredential)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}


	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res,nil, status, err := a.AccountUse.ResetPassword(ctx,req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      res,
	})
}


func (a *accountConn) GetAllDataCustomer(c *fiber.Ctx) error {

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	customerRes, status, err := a.AccountUse.AccountCustomerfile(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      customerRes,
	})
}

func (a *accountConn) GetAllDataStore(c *fiber.Ctx) error {

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	storeRes, status, err := a.AccountUse.AccountStorefile(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      storeRes,
	})
}

func (a *accountConn) GetAllDataAdmin(c *fiber.Ctx) error {

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	adminRes, status, err := a.AccountUse.AccountAdminfile(ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      adminRes,
	})
}