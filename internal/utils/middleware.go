package utils

// @TODO : make middleware like Auth
import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// Middleware JWT function
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		TokenLookup: "header:token",
		SigningKey:  []byte(secret),
	})
}
