package middlewares

import (
	
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
	"github.com/nutikuli/internProject_backend/internal/models/logdata/entities"
)

type logger struct {
	logRepo repository.LogRepo
}

func NewLogger(logRepo repository.LogRepo) *logger { 
	return &logger{logRepo: logRepo}
}

func (l *logger) LogRequest(c *fiber.Ctx) error {

    req := new(entities.LogGetReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"result":      nil,
		})
	}
	return c.Next()
}
