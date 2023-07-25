package helper

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @TODO : make helper response

func ErrorResponse(method string, err error) fiber.Map {
	return fiber.Map{
		"status":  false,
		"message": fmt.Sprintf("Failed to %s data", strings.ToUpper(method)),
		"errors":  strings.Split(err.Error(), "\n"),
		"data":    nil,
	}
}

func SuccessResponse(method string, data interface{}) fiber.Map {
	return fiber.Map{
		"status":  true,
		"message": fmt.Sprintf("Succeed to %s data", strings.ToUpper(method)),
		"errors":  nil,
		"data":    data,
	}
}
