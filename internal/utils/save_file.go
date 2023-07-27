package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveFile(ctx *fiber.Ctx, key string) (filename string, err error) {
	uniqueName := uuid.New().String()
	file, err := ctx.FormFile(key)

	if !strings.Contains(file.Header.Get("Content-Type"), "image") {
		return "", errors.New("File harus format gambar")
	}

	if err != nil {
		return "", errors.New("Gagal upload file")
	}

	err = ctx.SaveFile(file, fmt.Sprintf("./storage/%s-%s", uniqueName, file.Filename))
	if err != nil {
		return "", errors.New("Gagal upload file")
	}

	return fmt.Sprintf("%s-%s", uniqueName, file.Filename), nil
}
