package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type UsersClaims struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
