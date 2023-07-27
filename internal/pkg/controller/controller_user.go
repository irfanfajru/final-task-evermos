package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetMyProfile(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	UsersUseCase usecase.UsersUseCase
}

func NewUserController(UsersUseCase usecase.UsersUseCase) UserController {
	return &UserControllerImpl{
		UsersUseCase: UsersUseCase,
	}
}

func (uc *UserControllerImpl) GetMyProfile(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	res, _ := uc.UsersUseCase.GetById(c, userId)

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

func (uc *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := utils.GetUserIdJWT(ctx)
	data := new(dto.UpdateUserReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	res, err := uc.UsersUseCase.Update(c, userId, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}
