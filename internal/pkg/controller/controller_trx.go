package controller

import (
	"errors"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TrxController interface {
	CreateTrx(ctx *fiber.Ctx) error
	GetAllTrx(ctx *fiber.Ctx) error
	GetTrxById(ctx *fiber.Ctx) error
}

type TrxControllerImpl struct {
	TrxUseCase usecase.TrxUseCase
}

func NewTrxController(TrxUseCase usecase.TrxUseCase) TrxController {
	return &TrxControllerImpl{
		TrxUseCase: TrxUseCase,
	}
}

func (uc *TrxControllerImpl) CreateTrx(ctx *fiber.Ctx) error {
	c := ctx.Context()
	data := new(dto.CreateTrxReq)
	userId := utils.GetUserIdJWT(ctx)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}
	res, err := uc.TrxUseCase.CreateTrx(c, userId, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusCreated).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *TrxControllerImpl) GetAllTrx(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	filter := new(dto.FilterTrx)
	if err := ctx.QueryParser(filter); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	res, err := uc.TrxUseCase.GetAllTrx(c, userId, *filter)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *TrxControllerImpl) GetTrxById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	trxId := ctx.Params("id", "")
	userId := utils.GetUserIdJWT(ctx)
	if trxId == "" || trxId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}
	res, err := uc.TrxUseCase.GetTrxById(c, userId, trxId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}
