package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func ProdukRoute(r fiber.Router, containerConf *container.Container) {

	repoProduk := repository.NewProdukRepository(containerConf.Mysqldb)
	repoCategory := repository.NewCategoryRepository(containerConf.Mysqldb)
	repoToko := repository.NewTokoRepository(containerConf.Mysqldb)
	repoFotoProduk := repository.NewFotoProdukRepository(containerConf.Mysqldb)

	useCaseProduk := usecase.NewProdukUseCase(repoProduk, repoFotoProduk, repoToko, repoCategory)
	controller := controller.NewProdukController(useCaseProduk)
	authMidlleware := utils.NewAuthMiddleware(containerConf.Apps.SecretJwt)

	produkAPI := r.Group("/product")
	produkAPI.Post("", authMidlleware, controller.CreateProduk)
	produkAPI.Get("", controller.GetAllProduk)
	produkAPI.Get("/:id", controller.GetProdukById)
	produkAPI.Put("/:id", authMidlleware, controller.UpdateProdukById)
	produkAPI.Delete("/:id", authMidlleware, controller.DeleteProdukById)
}
