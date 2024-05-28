package datasource

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type RouteRepository interface {
	// TODO: Implemented model routers
}

func InitRoute(db *sqlx.DB, rG *fiber.Router) {
	// TODO: Init routers for using in main.go here
}
