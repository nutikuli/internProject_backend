package v1

import (
	"context"

	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/nutikuli/internProject_backend/internal/models/adminpermission"
	_dtos "github.com/nutikuli/internProject_backend/internal/models/adminpermission/dtos"
) 


type adminPermissionConn struct {
	AdminpermissionUse adminpermission.AdminpermissionUseCase
} 


func NewAdminHandler(AdminpermissionUse adminpermission.AdminpermissionUseCase) *adminPermissionConn {
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
	req := new(_dtos.AdminPermissionFileReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	adminpermissionRes, status, err := a.AdminpermissionUse.OnCreateAdminpermissionAccount(c, ctx, req.AdminpermissionData, req.FilesData)
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
		"result":_dtos.AdminTokenFileReqs{
			Adminpermission: adminpermissionRes,
		},
	})

}
