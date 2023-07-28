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
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UsersUseCase interface {
	Login(ctx context.Context, data dto.LoginReq) (res dto.LoginResp, err *helper.ErrorStruct)
	Register(ctx context.Context, data dto.RegisterReq) (res string, err *helper.ErrorStruct)
	GetById(ctx context.Context, userId string) (res dto.User, err *helper.ErrorStruct)
	Update(ctx context.Context, userId string, data dto.UpdateUserReq) (res string, err *helper.ErrorStruct)
}

type UsersUseCaseImpl struct {
	UsersRepository repository.UsersRepository
	TokoRepository  repository.TokoRepository
	jwtSecret       string
}

func NewUsersUseCase(UsersRepository repository.UsersRepository, TokoRepository repository.TokoRepository, jwtSecret string) UsersUseCase {
	return &UsersUseCaseImpl{
		UsersRepository: UsersRepository,
		TokoRepository:  TokoRepository,
		jwtSecret:       jwtSecret,
	}
}

func (alc *UsersUseCaseImpl) Login(ctx context.Context, data dto.LoginReq) (res dto.LoginResp, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.UsersRepository.FindByCredentials(ctx, data.NoTelp)
	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at FindByCredentials : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Err:  errors.New("No Telp atau kata sandi salah"),
			Code: fiber.StatusNotFound,
		}
	}

	if errRepo == nil {
		if !utils.CheckPasswordHash(data.KataSandi, resRepo.KataSandi) {
			return res, &helper.ErrorStruct{
				Err:  errors.New("No Telp atau kata sandi salah"),
				Code: fiber.StatusNotFound,
			}
		}
	}

	token := utils.CreateToken(resRepo, alc.jwtSecret)
	apiWilayah := repository.NewWilayahRepository()
	provRes, _ := apiWilayah.GetProvinceById(resRepo.IdProvinsi)
	regRes, _ := apiWilayah.GetRegencyById(resRepo.IdProvinsi, resRepo.IdKota)

	res = dto.LoginResp{
		Nama:         resRepo.Nama,
		NoTelp:       resRepo.Notelp,
		TanggalLahir: utils.DateResponse(resRepo.TanggalLahir),
		Tentang:      resRepo.Tentang,
		Pekerjaan:    resRepo.Pekerjaan,
		Email:        resRepo.Email,
		IdProvinsi:   provRes,
		IdKota:       regRes,
		Token:        token,
	}
	return res, nil
}

func (alc *UsersUseCaseImpl) Register(ctx context.Context, data dto.RegisterReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	userId, errRepo := alc.UsersRepository.Create(ctx, daos.User{
		Nama:         data.Nama,
		KataSandi:    utils.HashPassword(data.KataSandi),
		Notelp:       data.NoTelp,
		TanggalLahir: utils.ParseDate(data.TanggalLahir),
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
		IdProvinsi:   data.IdProvinsi,
		IdKota:       data.IdKota,
	})

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at Register : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	// create toko
	_, errRepoToko := alc.TokoRepository.CreateToko(ctx, daos.Toko{
		IdUser:   userId,
		NamaToko: fmt.Sprintf("Toko %s", data.Nama),
	})

	if errRepoToko != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at Register : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepoToko,
		}
	}

	return "Register Succeed", nil
}

func (alc *UsersUseCaseImpl) GetById(ctx context.Context, userId string) (res dto.User, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.UsersRepository.GetById(ctx, userId)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Found"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetById : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	apiWilayah := repository.NewWilayahRepository()
	provRes, _ := apiWilayah.GetProvinceById(resRepo.IdProvinsi)
	regRes, _ := apiWilayah.GetRegencyById(resRepo.IdProvinsi, resRepo.IdKota)

	res = dto.User{
		Nama:         resRepo.Nama,
		NoTelp:       resRepo.Notelp,
		TanggalLahir: utils.DateResponse(resRepo.TanggalLahir),
		Tentang:      resRepo.Tentang,
		Pekerjaan:    resRepo.Pekerjaan,
		Email:        resRepo.Email,
		IdProvinsi:   provRes,
		IdKota:       regRes,
	}
	return res, nil
}

func (alc *UsersUseCaseImpl) Update(ctx context.Context, userId string, data dto.UpdateUserReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.UsersRepository.Update(ctx, userId, daos.User{
		Nama:         data.Nama,
		KataSandi:    utils.HashPassword(data.KataSandi),
		Notelp:       data.NoTelp,
		TanggalLahir: utils.ParseDate(data.TanggalLahir),
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
		IdProvinsi:   data.IdProvinsi,
		IdKota:       data.IdKota,
	})

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at Update : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
