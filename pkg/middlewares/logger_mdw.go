package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nutikuli/internProject_backend/internal/models/logdata"
	"github.com/nutikuli/internProject_backend/internal/models/logdata/entities"
)

type LoggerAction struct {
	Menu   string
	Action string
}

type logger struct {
	logRepo logdata.LogRepository
}

func NewLogger(logRepo logdata.LogRepository) LoggerUsecase {
	return &logger{logRepo: logRepo}
}

type LoggerUsecase interface {
	LogRequest(c *fiber.Ctx, logAction *LoggerAction) error
}

func (l *logger) LogRequest(c *fiber.Ctx, logAction *LoggerAction) error {

	username := string(c.Request().Header.Peek("username"))
	if username == "" {
		username = "anonymous"
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	logReq := &entities.LogCreateReq{
		Fullname:      username,
		MenuRequest:   logAction.Menu,
		ActionRequest: logAction.Action,
	}

	_, err := l.logRepo.CreateLogData(ctx, logReq)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusInternalServerError),
			"status_code": http.StatusInternalServerError,
			"message":     fmt.Errorf("Error while trying to create logAction %v, Error: %v", *logReq, err.Error()),
			"result":      nil,
		})
	}

	log.Info("LogAction: ", logAction, " by: ", username)

	return c.Next()
}
