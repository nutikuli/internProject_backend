package middlewares

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
)

type logger struct {
	logRepo repository.LogRepo
}

func NewLogger(logRepo repository.LogRepo) *logger {
	return &logger{logRepo: logRepo}
}

func (l *logger) LogRequest() fiber.Handler {
	// นำข้อมูลจาก log created req ไปสร้าง Record เพิ่มใน Log table

	ctx := context.Background()
	newLog, err := l.logRepo.CreateLogData(ctx)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusForbidden),
			"status_code": http.StatusForbidden,
			"message":     "You don't have permission to access this resource",
			"raw_message": "",
			"result":      nil,
		})
	}

	return c.Next()
}
