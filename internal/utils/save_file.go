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

func SaveMultiFile(ctx *fiber.Ctx, key string) (filenames []string, err error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return filenames, errors.New("Gagal upload file")
	}

	files := form.File[key]

	// check content type
	for _, v := range files {
		if !strings.Contains(v.Header.Get("Content-Type"), "image") {
			return filenames, errors.New("File harus format gambar")
		}
	}

	// save file to storage
	for _, v := range files {
		uniqueName := uuid.New().String()
		fileName := fmt.Sprintf("%s-%s", uniqueName, v.Filename)
		err = ctx.SaveFile(v, fmt.Sprintf("./storage/%s", fileName))
		if err != nil {
			return filenames, errors.New("Gagal upload file")
		}

		filenames = append(filenames, fileName)
	}

	return filenames, nil
}
