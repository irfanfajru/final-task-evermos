package controller

import (
	"errors"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProdukController interface {
	CreateProduk(ctx *fiber.Ctx) error
	GetAllProduk(ctx *fiber.Ctx) error
	GetProdukById(ctx *fiber.Ctx) error
	UpdateProdukById(ctx *fiber.Ctx) error
	DeleteProdukById(ctx *fiber.Ctx) error
}

type ProdukControllerImpl struct {
	ProdukUseCase usecase.ProdukUseCase
}

func NewProdukController(ProdukUseCase usecase.ProdukUseCase) ProdukController {
	return &ProdukControllerImpl{
		ProdukUseCase: ProdukUseCase,
	}
}

func (uc *ProdukControllerImpl) CreateProduk(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	data := new(dto.CreateProdukReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	fotoProduk, errFile := utils.SaveMultiFile(ctx, "photos")
	if errFile != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errFile))
	}

	data.FotoProduk = fotoProduk
	res, err := uc.ProdukUseCase.CreateProduk(c, userId, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusCreated).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *ProdukControllerImpl) GetAllProduk(ctx *fiber.Ctx) error {
	c := ctx.Context()
	filter := new(dto.FilterProduk)
	if err := ctx.QueryParser(filter); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	res, err := uc.ProdukUseCase.GetAllProduk(c, *filter)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *ProdukControllerImpl) GetProdukById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	produkId := ctx.Params("id", "")
	if produkId == "" || produkId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}
	res, err := uc.ProdukUseCase.GetProdukById(c, produkId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *ProdukControllerImpl) UpdateProdukById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	produkId := ctx.Params("id")
	userId := utils.GetUserIdJWT(ctx)
	if produkId == "" || produkId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}

	data := new(dto.UpdateProdukReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	res, err := uc.ProdukUseCase.UpdateProdukById(c, userId, produkId, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *ProdukControllerImpl) DeleteProdukById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	produkId := ctx.Params("id")
	userId := utils.GetUserIdJWT(ctx)
	if produkId == "" || produkId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}

	res, err := uc.ProdukUseCase.DeleteProdukById(c, userId, produkId)

	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}
