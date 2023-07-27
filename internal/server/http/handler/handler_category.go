package handler

import (
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(r fiber.Router, containerConf *container.Container) {
	repo := repository.NewCategoryRepository(containerConf.Mysqldb)
	usecase := usecase.NewCategoryUseCase(repo)
	controller := controller.NewCategoryController(usecase)
	authMiddleware := utils.NewAuthMiddleware(containerConf.Apps.SecretJwt)

	categoryAPI := r.Group("/category")
	categoryAPI.Get("", controller.GetAllCategory)
	categoryAPI.Get("/:id", controller.GetCategoryById)
	categoryAPI.Post("", authMiddleware, utils.AdminRoleMiddleware, controller.CreateCategory)
	categoryAPI.Put("/:id", authMiddleware, utils.AdminRoleMiddleware, controller.UpdateCategoryById)
	categoryAPI.Delete("/:id", authMiddleware, utils.AdminRoleMiddleware, controller.DeleteCategoryById)
}
