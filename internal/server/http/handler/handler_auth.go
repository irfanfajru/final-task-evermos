package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewUsersRepository(containerConf.Mysqldb)
	useCase := usecase.NewUsersUseCase(repo, containerConf.Apps.SecretJwt)
	controller := controller.NewAuthController(useCase)

	authAPI := r.Group("/auth")
	authAPI.Post("/login", controller.Login)
	authAPI.Post("/register", controller.Register)
}
