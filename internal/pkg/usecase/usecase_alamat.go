package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AlamatUseCase interface {
	GetMyAlamat(ctx context.Context, userId string, params dto.AlamatFilter) (res []dto.Alamat, err *helper.ErrorStruct)
	GetMyAlamatById(ctx context.Context, userId string, alamatId string) (res dto.Alamat, err *helper.ErrorStruct)
	CreateAlamat(ctx context.Context, userId string, data dto.CreateAlamatReq) (res uint, err *helper.ErrorStruct)
	UpdateAlamat(ctx context.Context, alamatId string, userId string, data dto.UpdateAlamatReq) (res string, err *helper.ErrorStruct)
	DeleteAlamat(ctx context.Context, alamatId string, userId string) (res string, err *helper.ErrorStruct)
}

type AlamatUseCaseImpl struct {
	AlamatRepository repository.AlamatRepository
}

func NewAlamatUseCase(AlamatRepository repository.AlamatRepository) AlamatUseCase {
	return &AlamatUseCaseImpl{
		AlamatRepository: AlamatRepository,
	}
}

func (alc *AlamatUseCaseImpl) GetMyAlamat(ctx context.Context, userId string, params dto.AlamatFilter) (res []dto.Alamat, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.AlamatRepository.GetMyAlamat(ctx, userId, daos.FilterAlamat{
		JudulAlamat: params.JudulAlamat,
	})
	if errors.Is(errRepo, gorm.ErrRecordNotFound) || len(resRepo) == 0 {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Alamat"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllBooks : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		res = append(res, dto.Alamat{
			ID:           v.ID,
			JudulAlamat:  v.JudulAlamat,
			NamaPenerima: v.NamaPenerima,
			Notelp:       v.Notelp,
			DetailAlamat: v.DetailAlamat,
		})
	}

	return res, nil
}

func (alc *AlamatUseCaseImpl) GetMyAlamatById(ctx context.Context, userId string, alamatId string) (res dto.Alamat, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.AlamatRepository.GetMyAlamatById(ctx, userId, alamatId)
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

	res = dto.Alamat{
		ID:           resRepo.ID,
		JudulAlamat:  resRepo.JudulAlamat,
		NamaPenerima: resRepo.NamaPenerima,
		Notelp:       resRepo.Notelp,
		DetailAlamat: resRepo.DetailAlamat,
	}
	return res, nil
}

func (alc *AlamatUseCaseImpl) CreateAlamat(ctx context.Context, userId string, data dto.CreateAlamatReq) (res uint, err *helper.ErrorStruct) {
	userIdParsed, _ := strconv.Atoi(userId)
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.AlamatRepository.CreateAlamat(ctx, daos.Alamat{
		IdUser:       uint(userIdParsed),
		JudulAlamat:  data.JudulAlamat,
		NamaPenerima: data.NamaPenerima,
		Notelp:       data.Notelp,
		DetailAlamat: data.DetailAlamat,
	})

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateAlamat : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (alc *AlamatUseCaseImpl) UpdateAlamat(ctx context.Context, alamatId string, userId string, data dto.UpdateAlamatReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.AlamatRepository.UpdateAlamat(ctx, alamatId, userId, daos.Alamat{
		NamaPenerima: data.NamaPenerima,
		Notelp:       data.Notelp,
		DetailAlamat: data.DetailAlamat,
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

func (alc *AlamatUseCaseImpl) DeleteAlamat(ctx context.Context, alamatId string, userId string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.AlamatRepository.DeleteAlamat(ctx, alamatId, userId)

	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at DeleteAlamat : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
