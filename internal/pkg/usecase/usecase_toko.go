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

type TokoUseCase interface {
	GetMyToko(ctx context.Context, userId string) (res dto.MyToko, err *helper.ErrorStruct)
	GetAllToko(ctx context.Context, params dto.FilterToko) (res dto.TokoWithPagination, err *helper.ErrorStruct)
	GetTokoById(ctx context.Context, tokoId string) (res dto.Toko, err *helper.ErrorStruct)
	UpdateMyTokoById(ctx context.Context, userId string, tokoId string, data dto.UpdateTokoReq) (res string, err *helper.ErrorStruct)
}

type TokoUseCaseImpl struct {
	TokoRepository repository.TokoRepository
}

func NewTokoUseCase(TokoRepository repository.TokoRepository) TokoUseCase {
	return &TokoUseCaseImpl{
		TokoRepository: TokoRepository,
	}
}

func (alc *TokoUseCaseImpl) GetMyToko(ctx context.Context, userId string) (res dto.MyToko, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.TokoRepository.GetMyToko(ctx, userId)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Toko"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetMyToko : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.MyToko{
		ID:       resRepo.ID,
		NamaToko: resRepo.NamaToko,
		UrlFoto:  resRepo.UrlFoto,
		UserId:   userId,
	}
	return res, nil
}

func (alc *TokoUseCaseImpl) GetAllToko(ctx context.Context, params dto.FilterToko) (res dto.TokoWithPagination, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := alc.TokoRepository.GetAllToko(ctx, daos.FilterToko{
		Limit:    params.Limit,
		Offset:   params.Page,
		NamaToko: params.NamaToko,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) || len(resRepo) == 0 {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Toko"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllToko : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	var dataToko []dto.Toko
	for _, v := range resRepo {
		dataToko = append(dataToko, dto.Toko{
			ID:       v.ID,
			NamaToko: v.NamaToko,
			UrlFoto:  v.UrlFoto,
		})
	}
	res.Data = dataToko
	res.Limit = params.Limit
	res.Page = params.Page + 1
	return res, nil
}

func (alc *TokoUseCaseImpl) GetTokoById(ctx context.Context, tokoId string) (res dto.Toko, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.TokoRepository.GetTokoById(ctx, tokoId)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("Toko tidak ditemukan"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetTokoById : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = dto.Toko{
		ID:       resRepo.ID,
		NamaToko: resRepo.NamaToko,
		UrlFoto:  resRepo.UrlFoto,
	}

	return res, nil
}

func (alc *TokoUseCaseImpl) UpdateMyTokoById(ctx context.Context, userId string, tokoId string, data dto.UpdateTokoReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.TokoRepository.UpdateMyTokoById(ctx, userId, tokoId, daos.Toko{
		NamaToko: data.NamaToko,
		UrlFoto:  data.Photo,
	})
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateAlamat : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
