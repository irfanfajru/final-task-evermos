package utils

// @TODO : make middleware like Auth
import (
	"errors"
	"tugas_akhir_example/internal/helper"

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

// Middleware role admin
func AdminRoleMiddleware(ctx *fiber.Ctx) error {
	if IsAdminJWT(ctx) {
		ctx.Next()
		return nil
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(
		helper.ErrorResponse(string(ctx.Context().Method()), errors.New("Unauthorized")),
	)
}
