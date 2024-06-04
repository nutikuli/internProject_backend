package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/order"
	"github.com/nutikuli/internProject_backend/internal/models/order/dtos"
	"github.com/nutikuli/internProject_backend/internal/models/order/entities"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type orderConn struct {
	OrderUse order.OrderUsecase
}

func NewOrderHandler(orderUse order.OrderUsecase) *orderConn {
	return &orderConn{
		OrderUse: orderUse,
	}
}

func (o *orderConn) GetOrdersByStoreId(c *fiber.Ctx) error {
	req, err := strconv.ParseInt(c.Params("store_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request store_id param",
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	orders, status, err := o.OrderUse.OnGetOrdersByStoreId(ctx, &req)
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
		"result":      orders,
	})

}

func (o *orderConn) GetOrdersByCustomerId(c *fiber.Ctx) error {
	req, err := strconv.ParseInt(c.Params("customer_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request customer_id param",
			"result":      nil,
		})

	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	orders, status, err := o.OrderUse.OnGetOrdersByCustomerId(ctx, &req)
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
		"result":      orders,
	})
}

func (o *orderConn) GetOrderById(c *fiber.Ctx) error {

	storeIdReq, err := strconv.ParseInt(c.Params("store_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request store_id param",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	orderIdReq, err := strconv.ParseInt(c.Params("order_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request order_id param",
			"raw_message": err.Error(),
			"result":      nil,
		})

	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	req := &entities.StoreAndOrderIdReq{
		StoreId: storeIdReq,
		OrderId: orderIdReq,
	}

	orders, status, err := o.OrderUse.OnGetOrderById(ctx, req)
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
		"result":      orders,
	})
}

func (o *orderConn) CreateOrder(c *fiber.Ctx) error {
	req := new(dtos.OrderFileBankIdOrderProductsReq)
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

	files, status, err := o.OrderUse.OnCreateOrder(c, ctx, req.OrderData, req.FilesData, req.OrderProductsData)
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
		"result":      files,
	})
}

func (o *orderConn) UpdateOrderTransportDetailAndState(c *fiber.Ctx) error {
	req := new(dtos.OrderTransportStateReq)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	orderIdReq, err := strconv.ParseInt(c.Params("order_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request order_id param",
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

	status, err := o.OrderUse.OnUpdateOrderStatus(ctx, orderIdReq, req.StateData)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	if req.StateData.State == "SEND" && req.TransportData != nil {
		status, err := o.OrderUse.OnUpdateOrderTransportDetail(ctx, orderIdReq, req.TransportData)
		if err != nil {
			return c.Status(status).JSON(fiber.Map{
				"status":      http.StatusText(status),
				"status_code": status,
				"message":     err.Error(),
				"result":      nil,
			})
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
		"message":     "",
		"result":      nil,
	})
}
