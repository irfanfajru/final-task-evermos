package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryUseCase interface {
	GetAllCategory(ctx context.Context) (res []dto.Category, err *helper.ErrorStruct)
	GetCategoryById(ctx context.Context, categoryId string) (res dto.Category, err *helper.ErrorStruct)
	CreateCategory(ctx context.Context, data dto.CreateCategoryReq) (res uint, err *helper.ErrorStruct)
	UpdateCategoryById(ctx context.Context, categoryId string, data dto.UpdateCategoryReq) (res string, err *helper.ErrorStruct)
	DeleteCategoryById(ctx context.Context, categoryId string) (res string, err *helper.ErrorStruct)
}

type CategoryUseCaseImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(CategoryRepository repository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		CategoryRepository: CategoryRepository,
	}
}

func (alc *CategoryUseCaseImpl) GetAllCategory(ctx context.Context) (res []dto.Category, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.CategoryRepository.GetAllCategory(ctx)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) || len(resRepo) == 0 {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Category"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllCategory : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, dto.Category{
			ID:           v.ID,
			NamaCategory: v.NamaCategory,
		})
	}

	return res, nil
}

func (alc *CategoryUseCaseImpl) GetCategoryById(ctx context.Context, categoryId string) (res dto.Category, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.CategoryRepository.GetCategoryById(ctx, categoryId)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Alamat"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetMyAlamatById : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.Category{
		ID:           resRepo.ID,
		NamaCategory: resRepo.NamaCategory,
	}
	return res, nil
}

func (alc *CategoryUseCaseImpl) CreateCategory(ctx context.Context, data dto.CreateCategoryReq) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.CategoryRepository.CreateCategory(ctx, daos.Category{
		NamaCategory: data.NamaCategory,
	})

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateCategory : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (alc *CategoryUseCaseImpl) UpdateCategoryById(ctx context.Context, categoryId string, data dto.UpdateCategoryReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.CategoryRepository.UpdateCategoryById(ctx, categoryId, daos.Category{
		NamaCategory: data.NamaCategory,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateCategory : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (alc *CategoryUseCaseImpl) DeleteCategoryById(ctx context.Context, categoryId string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.CategoryRepository.DeleteCategoryById(ctx, categoryId)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at DeleteCategory : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
