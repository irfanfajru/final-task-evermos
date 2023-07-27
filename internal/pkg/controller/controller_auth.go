package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	UsersUseCase usecase.UsersUseCase
}

func NewAuthController(UsersUseCase usecase.UsersUseCase) AuthController {
	return &AuthControllerImpl{
		UsersUseCase: UsersUseCase,
	}
}

// login
func (uc *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	c := ctx.Context()
	data := new(dto.LoginReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	res, err := uc.UsersUseCase.Login(c, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}

// register
func (uc *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	c := ctx.Context()
	data := new(dto.RegisterReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(string(c.Method()), err))
	}

	res, err := uc.UsersUseCase.Register(c, *data)
	if err != nil {
		return ctx.Status(err.Code).JSON(helper.ErrorResponse(string(c.Method()), err.Err))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse(string(c.Method()), res))
}
