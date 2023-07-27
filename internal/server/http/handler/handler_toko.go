package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func TokoRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewTokoRepository(containerConf.Mysqldb)
	usecase := usecase.NewTokoUseCase(repo)
	controller := controller.NewTokoController(usecase)
	authMiddleware := utils.NewAuthMiddleware(containerConf.Apps.SecretJwt)

	tokoAPI := r.Group("/toko")
	tokoAPI.Get("/my", authMiddleware, controller.GetMyToko)
	tokoAPI.Get("/:id_toko", controller.GetTokoById)
	tokoAPI.Get("", controller.GetAllToko)
	tokoAPI.Put("/:id_toko", authMiddleware, controller.UpdateMyTokoById)
}
