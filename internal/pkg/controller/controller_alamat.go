package controller

import (
	"errors"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AlamatController interface {
	GetMyAlamat(ctx *fiber.Ctx) error
	GetMyAlamatById(ctx *fiber.Ctx) error
	CreateAlamat(ctx *fiber.Ctx) error
	UpdateAlamat(ctx *fiber.Ctx) error
	DeleteAlamat(ctx *fiber.Ctx) error
}

type AlamatControllerImpl struct {
	AlamatUseCase usecase.AlamatUseCase
}

func NewAlamatController(AlamatUseCase usecase.AlamatUseCase) AlamatController {
	return &AlamatControllerImpl{
		AlamatUseCase: AlamatUseCase,
	}
}

func (uc *AlamatControllerImpl) GetMyAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	filter := new(dto.AlamatFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse("GET", err))
	}

	res, err := uc.AlamatUseCase.GetMyAlamat(c, userId, dto.AlamatFilter{
		JudulAlamat: filter.JudulAlamat,
	})

	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse("get", err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse("get", res))
}

func (uc *AlamatControllerImpl) GetMyAlamatById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	alamatId := ctx.Params("id", "")
	if alamatId == "" || alamatId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse("get", errors.New("Bad request")))
	}

	res, err := uc.AlamatUseCase.GetMyAlamatById(c, userId, alamatId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse("get", err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse("get", res))
}

func (uc *AlamatControllerImpl) CreateAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	data := new(dto.CreateAlamatReq)

	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse("post", err))
	}

	res, err := uc.AlamatUseCase.CreateAlamat(c, userId, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse("post", err.Err))
	}

	return ctx.Status(fiber.StatusCreated).JSON(helper.SuccessResponse("post", res))
}

func (uc *AlamatControllerImpl) UpdateAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	alamatId := ctx.Params("id", "")
	if alamatId == "" || alamatId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse("put", errors.New("Bad request")))
	}

	data := new(dto.UpdateAlamatReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse("put", err))
	}

	res, err := uc.AlamatUseCase.UpdateAlamat(c, alamatId, userId, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse("put", err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse("put", res))
}

func (uc *AlamatControllerImpl) DeleteAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	alamatId := ctx.Params("id")
	userId := utils.GetUserIdJWT(ctx)
	if alamatId == "" || alamatId == "id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse("delete", errors.New("Bad request")))
	}

	res, err := uc.AlamatUseCase.DeleteAlamat(c, alamatId, userId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse("delete", err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse("delete", res))

}
