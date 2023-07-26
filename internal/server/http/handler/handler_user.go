package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewUsersRepository(containerConf.Mysqldb)
	useCase := usecase.NewUsersUseCase(repo, containerConf.Apps.SecretJwt)
	controller := controller.NewUserController(useCase)
	authMiddleware := utils.NewAuthMiddleware(containerConf.Apps.SecretJwt)

	userAPI := r.Group("/user")
	userAPI.Get("", authMiddleware, controller.GetMyProfile)
	userAPI.Put("", authMiddleware, controller.Update)
}
