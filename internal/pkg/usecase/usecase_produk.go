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
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ProdukUseCase interface {
	CreateProduk(ctx context.Context, userId string, data dto.CreateProdukReq) (res uint, err *helper.ErrorStruct)
	GetAllProduk(ctx context.Context, params dto.FilterProduk) (res dto.ProdukWithPagination, err *helper.ErrorStruct)
	GetProdukById(ctx context.Context, produkId string) (res dto.Produk, err *helper.ErrorStruct)
	UpdateProdukById(ctx context.Context, userId string, produkId string, data dto.UpdateProdukReq) (res string, err *helper.ErrorStruct)
	DeleteProdukById(ctx context.Context, userId string, produkId string) (res string, err *helper.ErrorStruct)
}

type ProdukUseCaseImpl struct {
	TokoRepository       repository.TokoRepository
	CategoryRepository   repository.CategoryRepository
	ProdukRepository     repository.ProdukRepository
	FotoProdukRepository repository.FotoProdukRepository
}

func NewProdukUseCase(ProdukRepository repository.ProdukRepository, FotoProdukRepository repository.FotoProdukRepository, TokoRepository repository.TokoRepository, CategoryRepository repository.CategoryRepository) ProdukUseCase {
	return &ProdukUseCaseImpl{
		ProdukRepository:     ProdukRepository,
		FotoProdukRepository: FotoProdukRepository,
		TokoRepository:       TokoRepository,
		CategoryRepository:   CategoryRepository,
	}
}

func (alc *ProdukUseCaseImpl) CreateProduk(ctx context.Context, userId string, data dto.CreateProdukReq) (res uint, err *helper.ErrorStruct) {
	// validate dto request
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// get toko from db
	resTokoRepo, errTokoRepo := alc.TokoRepository.GetMyToko(ctx, userId)
	if errTokoRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduk get my toko : %s", errTokoRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errTokoRepo,
		}
	}

	// get category from db
	resCategoryRepo, errCategoryRepo := alc.CategoryRepository.GetCategoryById(ctx, data.CategoryId)
	if errCategoryRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduk get category : %s", errCategoryRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New("Category Not found"),
		}
	}

	// create produk
	resRepoProduk, errRepoProduk := alc.ProdukRepository.CreateProduk(ctx, daos.Produk{
		IdToko:        resTokoRepo.ID,
		IdCategory:    resCategoryRepo.ID,
		NamaProduk:    data.NamaProduk,
		Slug:          slug.Make(data.NamaProduk),
		HargaReseller: strconv.Itoa(data.HargaReseller),
		HargaKonsumen: strconv.Itoa(data.HargaKonsumen),
		Stok:          uint(data.Stok),
		Deskripsi:     data.Deskripsi,
	})

	if errRepoProduk != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduk create produk : %s", errRepoProduk.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepoProduk,
		}
	}

	// create foto produk
	for _, v := range data.FotoProduk {
		_, errRepoFotoProduk := alc.FotoProdukRepository.CreateFotoProduk(ctx, daos.FotoProduk{
			IdProduk: resRepoProduk,
			Url:      v,
		})

		if errRepoFotoProduk != nil {
			helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduk create foto produk : %s", errRepoFotoProduk.Error()))
			return res, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  errRepoFotoProduk,
			}
		}
	}

	return resRepoProduk, nil
}

func (alc *ProdukUseCaseImpl) GetAllProduk(ctx context.Context, params dto.FilterProduk) (res dto.ProdukWithPagination, err *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := alc.ProdukRepository.GetAllProduk(ctx, daos.FilterProduk{
		Limit:      params.Limit,
		Offset:     params.Page,
		NamaProduk: params.NamaProduk,
		CategoryId: params.CategoryId,
		TokoId:     params.TokoId,
		MaxHarga:   params.MaxHarga,
		MinHarga:   params.MinHarga,
	})

	if errors.Is(errRepo, gorm.ErrRecordNotFound) || len(resRepo) == 0 {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Produk"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAllProduk : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	var dataProduk []dto.Produk
	for _, v := range resRepo {

		var fotoProduk []dto.FotoProduk
		for _, foto := range v.FotoProduk {
			fotoProduk = append(fotoProduk, dto.FotoProduk{
				ID:       foto.ID,
				IdProduk: foto.IdProduk,
				Url:      foto.Url,
			})
		}

		dataProduk = append(dataProduk, dto.Produk{
			ID:            v.ID,
			NamaProduk:    v.NamaProduk,
			Slug:          v.Slug,
			HargaReseller: v.HargaReseller,
			HargaKonsumen: v.HargaKonsumen,
			Stok:          int(v.Stok),
			Deskripsi:     v.Deskripsi,
			Toko: dto.Toko{
				ID:       v.Toko.ID,
				NamaToko: v.Toko.NamaToko,
				UrlFoto:  v.Toko.UrlFoto,
			},
			Category: dto.Category{
				ID:           v.Category.ID,
				NamaCategory: v.Category.NamaCategory,
			},
			FotoProduk: fotoProduk,
		})
	}
	res.Data = dataProduk
	res.Limit = params.Limit
	res.Page = params.Page + 1
	return res, nil
}

func (alc *ProdukUseCaseImpl) GetProdukById(ctx context.Context, produkId string) (res dto.Produk, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.ProdukRepository.GetProdukById(ctx, produkId)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Product"),
		}
	}

	if errRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetTokoById : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	var fotoProduk []dto.FotoProduk
	for _, val := range resRepo.FotoProduk {
		fotoProduk = append(fotoProduk, dto.FotoProduk{
			ID:       val.ID,
			IdProduk: val.IdProduk,
			Url:      val.Url,
		})
	}

	res = dto.Produk{
		ID:            resRepo.ID,
		NamaProduk:    resRepo.NamaProduk,
		Slug:          resRepo.Slug,
		HargaReseller: resRepo.HargaReseller,
		HargaKonsumen: resRepo.HargaKonsumen,
		Stok:          int(resRepo.Stok),
		Deskripsi:     resRepo.Deskripsi,
		Toko: dto.Toko{
			ID:       resRepo.Toko.ID,
			NamaToko: resRepo.Toko.NamaToko,
			UrlFoto:  resRepo.Toko.UrlFoto,
		},
		Category: dto.Category{
			ID:           resRepo.Category.ID,
			NamaCategory: resRepo.Category.NamaCategory,
		},
		FotoProduk: fotoProduk,
	}
	return res, nil
}

func (alc *ProdukUseCaseImpl) UpdateProdukById(ctx context.Context, userId string, produkId string, data dto.UpdateProdukReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// get toko from db
	resTokoRepo, errTokoRepo := alc.TokoRepository.GetMyToko(ctx, userId)
	if errTokoRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduk get my toko : %s", errTokoRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errTokoRepo,
		}
	}

	resProdukRepo, errProdukRepo := alc.ProdukRepository.UpdateProdukById(ctx, resTokoRepo.ID, produkId, daos.Produk{
		NamaProduk: data.NamaProduk,
		Slug:       slug.Make(data.NamaProduk),
	})

	if errors.Is(errProdukRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errProdukRepo,
		}
	}

	if errProdukRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at UpdateAlamat : %s", errProdukRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errProdukRepo,
		}
	}

	return resProdukRepo, nil
}

func (alc *ProdukUseCaseImpl) DeleteProdukById(ctx context.Context, userId string, produkId string) (res string, err *helper.ErrorStruct) {
	// get toko from db
	resTokoRepo, errTokoRepo := alc.TokoRepository.GetMyToko(ctx, userId)
	if errTokoRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateProduk get my toko : %s", errTokoRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errTokoRepo,
		}
	}

	resProdukRepo, errProdukRepo := alc.ProdukRepository.DeleteProdukById(ctx, resTokoRepo.ID, produkId)
	if errors.Is(errProdukRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errProdukRepo,
		}
	}

	if errProdukRepo != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at DeleteCategory : %s", errProdukRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errProdukRepo,
		}
	}

	return resProdukRepo, nil
}
