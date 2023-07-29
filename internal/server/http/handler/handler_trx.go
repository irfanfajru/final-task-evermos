package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func TrxRouter(r fiber.Router, containerConf *container.Container) {
	produkRepo := repository.NewProdukRepository(containerConf.Mysqldb)
	alamatRepo := repository.NewAlamatRepository(containerConf.Mysqldb)
	trxRepo := repository.NewTrxRepository(containerConf.Mysqldb)
	detailTrxRepo := repository.NewDetailTrxRepository(containerConf.Mysqldb)
	logProdukRepo := repository.NewLogProdukRepository(containerConf.Mysqldb)

	usecase := usecase.NewTrxUseCase(produkRepo, alamatRepo, trxRepo, detailTrxRepo, logProdukRepo, containerConf.Mysqldb)
	controller := controller.NewTrxController(usecase)
	authMiddleware := utils.NewAuthMiddleware(containerConf.Apps.SecretJwt)

	trxAPI := r.Group("/trx")
	trxAPI.Post("", authMiddleware, controller.CreateTrx)

}
