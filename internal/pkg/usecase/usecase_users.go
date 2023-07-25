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
)

type UsersUseCase interface {
	Login(ctx context.Context, data dto.LoginReq) (res dto.LoginResp, err *helper.ErrorStruct)
	Register(ctx context.Context, data dto.RegisterReq) (res string, err *helper.ErrorStruct)
}

type UsersUseCaseImpl struct {
	UsersRepository repository.UsersRepository
	jwtSecret       string
}

func NewUsersUseCase(UsersRepository repository.UsersRepository, jwtSecret string) UsersUseCase {
	return &UsersUseCaseImpl{
		UsersRepository: UsersRepository,
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

	res = dto.LoginResp{
		Nama:         resRepo.Nama,
		NoTelp:       resRepo.Notelp,
		TanggalLahir: utils.DateResponse(resRepo.TanggalLahir),
		Tentang:      resRepo.Tentang,
		Pekerjaan:    resRepo.Pekerjaan,
		Email:        resRepo.Email,
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

	_, errRepo := alc.UsersRepository.Create(ctx, daos.User{
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

	return "Register Succeed", nil
}
