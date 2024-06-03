package v1

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nutikuli/internProject_backend/internal/models/account"
	"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	_accEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/admin"
	"github.com/nutikuli/internProject_backend/internal/models/customer"
	"github.com/nutikuli/internProject_backend/internal/models/store"
)

type accountConn struct {
	AccountUse  account.AccountUsecase
	StoreUse    store.StoreUsecase
	CustomerUse customer.CustomerUsecase
	AdminUse    admin.AdminUseCase
}

func NewAccountHandler(accountUse account.AccountUsecase) *accountConn {
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
	log.Debug("loging")
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	usrPassport, userToken, status, err := a.AccountUse.Login(ctx, req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	var accRes interface{}

	switch userToken.Role {
	case "CUSTOMER":
		acc, status, err := a.CustomerUse.OnGetCustomerById(ctx, usrPassport.Id)
		if err != nil {
			return c.Status(status).JSON(fiber.Map{
				"status":      status,
				"status_code": status,
				"message":     err.Error(),
				"result":      nil,
			})
		}
		accRes = acc
	case "STORE":
		acc, status, err := a.StoreUse.OnGetStoreById(ctx, usrPassport.Id)
		if err != nil {
			return c.Status(status).JSON(fiber.Map{
				"status":      status,
				"status_code": status,
				"message":     err.Error(),
				"result":      nil,
			})
		}
		accRes = acc
	case "ADMIN":
		acc, status, err := a.AdminUse.OnGetAdminById(ctx, usrPassport.Id)
		if err != nil {
			return c.Status(status).JSON(fiber.Map{
				"status":      status,
				"status_code": status,
				"message":     err.Error(),
				"result":      nil,
			})
		}
		accRes = acc
	default:
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     errors.New("Can't query the Account Table, Invalid role"),
			"result":      nil,
		})

	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result": dtos.AccountLoginRes{
			AccountData: accRes,
			UserToken:   *userToken,
		},
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

	userToken, status, err := a.AccountUse.CheckOTP(c, ctx, req)
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
		"result":      userToken,
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
	log.Debug(req)
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()
	userPass, status, err := a.AccountUse.ResetPassword(ctx, req)

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
		"result":      userPass,
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
