package utils

// @TODO : make middleware like Auth
import (
	"errors"
	"tugas_akhir_example/internal/helper"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// Middleware JWT function
func NewAuthMiddleware(secret string, roles ...string) fiber.Handler {
	for _, role := range roles {
		if role == "admin" {
			return jwtware.New(jwtware.Config{
				TokenLookup:    "header:token",
				SigningKey:     []byte(secret),
				SuccessHandler: AdminRoleMiddleware,
				ErrorHandler:   ErrorJWTHandler,
			})
		}
	}
	return jwtware.New(jwtware.Config{
		TokenLookup:  "header:token",
		SigningKey:   []byte(secret),
		ErrorHandler: ErrorJWTHandler,
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

// error jwt handler
func ErrorJWTHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(
		helper.ErrorResponse(string(ctx.Context().Method()), err),
	)
}
