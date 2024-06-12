package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		fmt.Println(accessToken)
		if accessToken == "" {
			log.Println("error, authorization header is empty.")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":      "Unauthorized",
				"status_code": http.StatusUnauthorized,
				"message":     "unauthorized access",
				"result":      nil,
			})
		}

		secretKey := os.Getenv("JWT_SECRET_TOKEN")
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error, unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(secretKey), nil
		})
		if err != nil {
			log.Println(err.Error())
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":      http.StatusText(http.StatusUnauthorized),
				"status_code": http.StatusUnauthorized,
				"message":     "error, unauthorized",
				"result":      nil,
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			c.Locals("email", claims["email"])
			c.Locals("role", claims["role"])
			return c.Next()
		} else {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":      http.StatusText(http.StatusUnauthorized),
				"status_code": http.StatusUnauthorized,
				"message":     "error, unauthorized",
				"result":      nil,
			})
		}
	}
}
