package v1

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/store"
	"github.com/nutikuli/internProject_backend/internal/models/store/dtos"
)

type storeConn struct {
	storeUse store.StoreUsecase
}

func NewStoreHandler(storeUse store.StoreUsecase) *storeConn {
	return &storeConn{
		storeUse: storeUse,
	}
}

func (s *storeConn) RegisterStoreAccount(c *fiber.Ctx) error {
	req := new(dtos.StoreFileReq)
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

	storeRes, userToken, status, err := s.storeUse.OnCreateStoreAccount(c, ctx, req.StoreRegisterData, req.FilesData)
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
		"result": dtos.StoreTokenFileRes{
			Store: storeRes,
			Token: userToken,
		},
	})

}

func (s *storeConn) GetStoreById(c *fiber.Ctx) error {
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

	storeRes, status, err := s.storeUse.OnGetStoreById(ctx, &req64)
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
		"result":      storeRes,
	})
}
