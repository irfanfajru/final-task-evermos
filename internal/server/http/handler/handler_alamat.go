package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func AlamatRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewAlamatRepository(containerConf.Mysqldb)
	useCase := usecase.NewAlamatUseCase(repo)
	controller := controller.NewAlamatController(useCase)
	authMiddleware := utils.NewAuthMiddleware(containerConf.Apps.SecretJwt)

	alamatAPI := r.Group("/user/alamat")
	alamatAPI.Get("", authMiddleware, controller.GetMyAlamat)
	alamatAPI.Get("/:id", authMiddleware, controller.GetMyAlamatById)
	alamatAPI.Post("", authMiddleware, controller.CreateAlamat)
	alamatAPI.Put("/:id", authMiddleware, controller.UpdateAlamat)
	alamatAPI.Delete("/:id", authMiddleware, controller.DeleteAlamat)
}
