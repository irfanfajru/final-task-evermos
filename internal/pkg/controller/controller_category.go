package controller

import (
	"errors"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	GetAllCategory(ctx *fiber.Ctx) error
	GetCategoryById(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategoryById(ctx *fiber.Ctx) error
	DeleteCategoryById(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryController(CategoryUseCase usecase.CategoryUseCase) CategoryController {
	return &CategoryControllerImpl{
		CategoryUseCase: CategoryUseCase,
	}
}

func (uc *CategoryControllerImpl) GetAllCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()
	res, err := uc.CategoryUseCase.GetAllCategory(c)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *CategoryControllerImpl) GetCategoryById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId := ctx.Params("id", "")
	if categoryId == "" || categoryId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}

	res, err := uc.CategoryUseCase.GetCategoryById(c, categoryId)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()
	data := new(dto.CreateCategoryReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	res, err := uc.CategoryUseCase.CreateCategory(c, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusCreated).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *CategoryControllerImpl) UpdateCategoryById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId := ctx.Params("id")
	if categoryId == "" || categoryId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}

	data := new(dto.UpdateCategoryReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	res, err := uc.CategoryUseCase.UpdateCategoryById(c, categoryId, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *CategoryControllerImpl) DeleteCategoryById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryId := ctx.Params("id")
	if categoryId == "" || categoryId == ":id" {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), errors.New("Bad request")))
	}

	res, err := uc.CategoryUseCase.DeleteCategoryById(c, categoryId)

	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}
