package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

func PermissionRoleGuard(roleGuard string, opt ...[]interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accRoleFromToken, _ := c.Locals("role").(string)
		log.Debug(accRoleFromToken)

		if accRoleFromToken != roleGuard {
			if len(opt) > 0 {
				opt := opt[0]
				log.Debug("opt: ", opt)
				if utils.Contains(opt, "PREVENT_DEFAULT_ACTION") {
					return c.Next()
				}
			}
			return c.Status(http.StatusForbidden).JSON(fiber.Map{
				"status":      http.StatusText(http.StatusForbidden),
				"status_code": http.StatusForbidden,
				"message":     "You don't have permission to access this resource",
				"raw_message": "",
				"result":      nil,
			})
		} else {
			if len(opt) > 0 {
				opt := opt[0]
				if utils.Contains(opt, "OWNER_ACTION_FORBIDDEN") {
					return c.Status(http.StatusForbidden).JSON(fiber.Map{
						"status":      http.StatusText(http.StatusForbidden),
						"status_code": http.StatusForbidden,
						"message":     "Owner action is forbidden",
						"raw_message": "",
						"result":      nil,
					})

				}
			}
			return c.Next()
		}
	}
}
