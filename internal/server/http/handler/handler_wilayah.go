package handler

import (
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

func WilayahRoute(r fiber.Router) {
	repo := repository.NewWilayahRepository()
	usecase := usecase.NewWilayahUseCase(repo)
	controller := controller.NewWilayahController(usecase)

	wilayahAPI := r.Group("/provcity")
	wilayahAPI.Get("/listprovincies", controller.GetAllProvince)
	wilayahAPI.Get("/listcities/:prov_id", controller.GetAllRegency)
	wilayahAPI.Get("/detailprovince/:prov_id", controller.GetProvinceById)
	wilayahAPI.Get("/detailcity/:city_id", controller.GetRegencyById)
}
