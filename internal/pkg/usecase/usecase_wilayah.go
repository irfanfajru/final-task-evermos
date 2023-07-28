package usecase

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

type WilayahUseCase interface {
	GetAllProvince(ctx context.Context) (res []dto.Province, err *helper.ErrorStruct)
	GetAllRegency(ctx context.Context, provinceId string) (res []dto.Regency, err *helper.ErrorStruct)
	GetProvinceById(ctx context.Context, provinceId string) (res dto.Province, err *helper.ErrorStruct)
	GetRegencyById(ctx context.Context, regencyId string) (res dto.Regency, err *helper.ErrorStruct)
}

type WilayahUseCaseImpl struct {
	WilayahRepository repository.WilayahRepository
}

func NewWilayahUseCase(WilayahRepository repository.WilayahRepository) WilayahUseCase {
	return &WilayahUseCaseImpl{
		WilayahRepository: WilayahRepository,
	}
}

func (alc *WilayahUseCaseImpl) GetAllProvince(ctx context.Context) (res []dto.Province, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.WilayahRepository.GetAllProvince()
	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllProvince : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	return resRepo, nil
}

func (alc *WilayahUseCaseImpl) GetAllRegency(ctx context.Context, provinceId string) (res []dto.Regency, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.WilayahRepository.GetAllRegency(provinceId)
	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllRegency : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	return resRepo, nil
}

func (alc *WilayahUseCaseImpl) GetProvinceById(ctx context.Context, provinceId string) (res dto.Province, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.WilayahRepository.GetProvinceById(provinceId)
	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllRegency : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	return resRepo, nil
}

func (alc *WilayahUseCaseImpl) GetRegencyById(ctx context.Context, regencyId string) (res dto.Regency, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.WilayahRepository.GetRegencyById(regencyId)
	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllRegency : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}
	return resRepo, nil
}
