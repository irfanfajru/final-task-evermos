package controller

import (
	"errors"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TokoController interface {
	GetMyToko(ctx *fiber.Ctx) error
	GetAllToko(ctx *fiber.Ctx) error
	GetTokoById(ctx *fiber.Ctx) error
	UpdateMyTokoById(ctx *fiber.Ctx) error
}

type TokoControllerImpl struct {
	TokoUseCase usecase.TokoUseCase
}

func NewTokoController(TokoUseCase usecase.TokoUseCase) TokoController {
	return &TokoControllerImpl{
		TokoUseCase: TokoUseCase,
	}
}

func (uc *TokoControllerImpl) GetMyToko(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)

	res, err := uc.TokoUseCase.GetMyToko(c, userId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *TokoControllerImpl) GetAllToko(ctx *fiber.Ctx) error {
	c := ctx.Context()
	filter := new(dto.FilterToko)
	if err := ctx.QueryParser(filter); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}
	res, err := uc.TokoUseCase.GetAllToko(c, dto.FilterToko{
		Page:     filter.Page,
		Limit:    filter.Limit,
		NamaToko: filter.NamaToko,
	})
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *TokoControllerImpl) GetTokoById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	tokoId := ctx.Params("id_toko", "")
	if tokoId == "" || tokoId == ":id_toko" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}
	res, err := uc.TokoUseCase.GetTokoById(c, tokoId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *TokoControllerImpl) UpdateMyTokoById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	tokoId := ctx.Params("id_toko", "")

	if tokoId == "" || tokoId == ":id_toko" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}

	data := new(dto.UpdateTokoReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	// save file
	fileName, errFile := utils.SaveFile(ctx, "photo")
	if errFile != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errFile))
	}

	data.Photo = fileName

	res, err := uc.TokoUseCase.UpdateMyTokoById(c, userId, tokoId, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}
