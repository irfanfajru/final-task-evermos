package controller

import (
	"errors"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type WilayahController interface {
	GetAllProvince(ctx *fiber.Ctx) error
	GetAllRegency(ctx *fiber.Ctx) error
	GetProvinceById(ctx *fiber.Ctx) error
	GetRegencyById(ctx *fiber.Ctx) error
}

type WilayahControllerImpl struct {
	WilayahUseCase usecase.WilayahUseCase
}

func NewWilayahController(WilayahUseCase usecase.WilayahUseCase) WilayahController {
	return &WilayahControllerImpl{
		WilayahUseCase: WilayahUseCase,
	}
}

func (uc *WilayahControllerImpl) GetAllProvince(ctx *fiber.Ctx) error {
	c := ctx.Context()
	res, err := uc.WilayahUseCase.GetAllProvince(c)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *WilayahControllerImpl) GetAllRegency(ctx *fiber.Ctx) error {
	c := ctx.Context()
	provId := ctx.Params("prov_id")
	if provId == "" || provId == ":prov_id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}
	res, err := uc.WilayahUseCase.GetAllRegency(c, provId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *WilayahControllerImpl) GetProvinceById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	provId := ctx.Params("prov_id")
	if provId == "" || provId == ":prov_id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}
	res, err := uc.WilayahUseCase.GetProvinceById(c, provId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *WilayahControllerImpl) GetRegencyById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	cityId := ctx.Params("city_id")
	if cityId == "" || cityId == ":city_id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}
	res, err := uc.WilayahUseCase.GetRegencyById(c, cityId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}
