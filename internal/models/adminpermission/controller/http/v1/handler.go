package v1

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2/log"

	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/nutikuli/internProject_backend/internal/models/adminpermission"
	_dtos "github.com/nutikuli/internProject_backend/internal/models/adminpermission/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission/entities"
) 


type adminPermissionConn struct {
	AdminpermissionUse adminpermission.AdminpermissionUseCase
} 


func NewAdminPermissionHandler(AdminpermissionUse adminpermission.AdminpermissionUseCase) *adminPermissionConn {
	return &adminPermissionConn{
		AdminpermissionUse: AdminpermissionUse,
	}
}  



func (a *adminPermissionConn) GetAdminePermissionById(c *fiber.Ctx) error {
	req, err := c.ParamsInt("adminpermission_id")

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

	adminpermissionRes, status, err := a.AdminpermissionUse.OnGetAdminpermissionById(c,ctx, &req64)
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
		"result":      adminpermissionRes,
	})
}

func (a *adminPermissionConn) CreateAdminPermissionAccount(c *fiber.Ctx) error {
	req := new(entities.AdminPermissionCreatedReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	} 

	log.Debug("req=====>" ,req)

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	adminpermissionRes, status, err := a.AdminpermissionUse.OnCreateAdminpermissionAccount(c, ctx, req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      status,
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	} 

	log.Debug("adres=====>",adminpermissionRes)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":      fiber.StatusCreated,
		"status_code": fiber.StatusCreated,
		"message":     "",
		"result":_dtos.AdminTokenFileReqs{
			Adminpermission: adminpermissionRes,
		},
	})

} 


func (a *adminPermissionConn) UpdateAdminpermissionById(c *fiber.Ctx) error {

	var req = new(entities.AdminPermissionUpdatedReq)
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
		reqP        = c.Params("adminpermission_id")
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

	adminperId, err := strconv.ParseInt(reqP, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"raw_message": err.Error(),
			"message":     "error, invalid user_id params",
			"result":      nil,
		})
	}

	status, cE := a.AdminpermissionUse.OnUpdateAdminPermissionById(ctx, adminperId, req)
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

func (a *adminPermissionConn) DeletedAdminPermissionByID(c *fiber.Ctx) error {
	req, err := c.ParamsInt("adminpermission_id")
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

	adminpermission, err := a.AdminpermissionUse.OnDeletedAdminPermission(ctx, req64)
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
		"result":      adminpermission,
	})
}
