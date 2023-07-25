package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

type UserController interface {
	GetMyProfile(ctx *fiber.Ctx) error
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
	user := ctx.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	userId := claims["UserId"].(string)
	res, _ := uc.UsersUseCase.GetById(c, userId)

	return ctx.Status(fiber.StatusOK).JSON(helper.SuccessResponse("get", res))
}
