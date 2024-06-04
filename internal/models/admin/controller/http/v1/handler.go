package v1

import (
	"context"
	"net/http"
	"strconv"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/nutikuli/internProject_backend/internal/models/admin"
	"github.com/nutikuli/internProject_backend/internal/models/admin/dtos"
	
) 


type adminConn struct {
	AdminUse admin.AdminUseCase
} 


func NewAdminHandler(AdminUse admin.AdminUseCase) *adminConn {
	return &adminConn{
		AdminUse: AdminUse,
	}
} 



func (a *adminConn) GetAdmineById(c *fiber.Ctx) error {
	req, err := c.ParamsInt("admin_id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request admin_id param",
			"result":      nil,
		})
	}

	req64 := int64(req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	adminRes, status, err := a.AdminUse.OnGetAdminById(ctx, req64)
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
		"message":     "",
		"result":      adminRes,
	})
}

func (a *adminConn) RegisterAdminAccount(c *fiber.Ctx) error {
	req := new(dtos.AdminCreateFileReq)
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

	adminRes, userToken, status, err := a.AdminUse.OnCreateAdminAccount(c, ctx , req.AdminData , req.FilesData)
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
		"message":     "",
		"result": dtos.AdminTokenFileReqs{
			Admin: adminRes,
			Token: userToken,
		},
	})

}



func (a *adminConn) UpdateAdminById(c *fiber.Ctx) error {
	adminId, err := c.ParamsInt("admin_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, invalid request admin_id param",
			"result":      nil,
		})
	}

	adId64 := int64(adminId)

	req := new(dtos.AdminFileUpdateReq)
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

	admin, status, err := a.AdminUse.OnUpdateAdminById(c, ctx, adId64,req.AdminData,req.FilesData)
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
		"result":      admin,
	})
}


func (a *adminConn) DeletedAdminByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("admin_id"), 10, 64)
	log.Debug("id=====>",id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request id param",
			"result":      nil,
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 30*time.Second)
	defer cancel()

	admin, err := a.AdminUse.AdminDeleted(ctx, id)
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
		"status":      http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
		"message":     "",
		"result":      admin,
	})
} 

func (a *adminConn) GetAllAdmin(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)
	defer cancel()

	admins, status, err := a.AdminUse.OnGetAllUserAdmin(ctx)
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
		"result":      admins,
	})
}
