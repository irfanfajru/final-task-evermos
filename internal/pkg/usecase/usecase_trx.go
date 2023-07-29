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
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TrxUseCase interface {
	CreateTrx(ctx context.Context, userId string, data dto.CreateTrxReq) (res uint, err *helper.ErrorStruct)
}

type TrxUseCaseImpl struct {
	ProdukRepository    repository.ProdukRepository
	AlamatRepository    repository.AlamatRepository
	TrxRepository       repository.TrxRepository
	DetailTrxRepository repository.DetailTrxRepository
	LogProdukRepository repository.LogProdukRepository
	db                  *gorm.DB
}

func NewTrxUseCase(
	ProdukRepository repository.ProdukRepository,
	AlamatRepository repository.AlamatRepository,
	TrxRepository repository.TrxRepository,
	DetailTrxRepository repository.DetailTrxRepository,
	LogProdukRepository repository.LogProdukRepository,
	db *gorm.DB,
) TrxUseCase {

	return &TrxUseCaseImpl{
		ProdukRepository:    ProdukRepository,
		AlamatRepository:    AlamatRepository,
		TrxRepository:       TrxRepository,
		DetailTrxRepository: DetailTrxRepository,
		LogProdukRepository: LogProdukRepository,
		db:                  db,
	}
}

func (alc *TrxUseCaseImpl) CreateTrx(ctx context.Context, userId string, data dto.CreateTrxReq) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// prepare daos
	var trxData daos.Trx
	var detailTrxData []daos.DetailTrx

	// begin transaction
	Tx := alc.db.Begin()

	// check alamat
	resAlamat, errAlamat := alc.AlamatRepository.GetMyAlamatById(ctx, userId, fmt.Sprint(data.AlamatKirim))
	if errAlamat != nil {
		Tx.Rollback()
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetAlamatMyAlamatById : %s", errAlamat.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New(fmt.Sprintf("Alamat %v", errAlamat)),
		}
	}
	trxData.AlamatPengiriman = resAlamat.ID
	trxData.IdUser = resAlamat.IdUser
	trxData.MethodBayar = data.MethodBayar
	trxData.KodeInvoice = utils.GenerateInvoiceCode()

	//produk and qty
	for _, produk := range data.DetailTrx {
		// check produk
		resProduk, errProduk := alc.ProdukRepository.GetProdukById(ctx, fmt.Sprint(produk.ProdukId))
		if errProduk != nil {
			Tx.Rollback()
			helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at GetProdukById : %s", errProduk.Error()))
			return res, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  errors.New(fmt.Sprintf("Produk %v", errProduk)),
			}
		}

		// check stok
		if resProduk.Stok < produk.Kuantitas {
			Tx.Rollback()
			helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprint("Error at Check stok : not enough stok"))
			return res, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  errors.New("Stok produk tidak mencukupi"),
			}
		}

		// create log produk
		resLogProduk, errLogProduk := alc.LogProdukRepository.CreateLogProdukWithTx(ctx, daos.LogProduk{
			IdProduk:      resProduk.ID,
			IdToko:        resProduk.IdToko,
			IdCategory:    resProduk.IdCategory,
			NamaProduk:    resProduk.NamaProduk,
			Slug:          resProduk.Slug,
			HargaReseller: resProduk.HargaReseller,
			HargaKonsumen: resProduk.HargaKonsumen,
			Deskripsi:     resProduk.Deskripsi,
		}, Tx)

		if errLogProduk != nil {
			Tx.Rollback()
			helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateLogProduk : %s", errLogProduk.Error()))
			return res, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  errLogProduk,
			}
		}

		// update produk stok
		_, errProdukStok := alc.ProdukRepository.UpdateProdukByIdWithTx(ctx, resProduk.IdToko, fmt.Sprint(resProduk.ID), daos.Produk{
			Stok: resProduk.Stok - produk.Kuantitas,
		}, Tx)

		if errProdukStok != nil {
			Tx.Rollback()
			helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at Update stok produk : %s", errProdukStok.Error()))
			return res, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  errProdukStok,
			}
		}

		harga, _ := strconv.Atoi(resProduk.HargaKonsumen)
		detailTrxData = append(detailTrxData, daos.DetailTrx{
			IdLogProduk: resLogProduk,
			IdToko:      resProduk.IdToko,
			Kuantitas:   produk.Kuantitas,
			HargaTotal:  produk.Kuantitas * uint(harga),
		})

		trxData.HargaTotal += produk.Kuantitas * uint(harga)
	}

	// create trx
	resTrx, errTrx := alc.TrxRepository.CreateTrxWithTx(ctx, trxData, Tx)
	if errTrx != nil {
		Tx.Rollback()
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateTrx : %s", errTrx.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errTrx,
		}
	}

	// create detail trx
	for _, dtlTrx := range detailTrxData {
		dtlTrx.IdTrx = resTrx
		_, errDtl := alc.DetailTrxRepository.CreateDetailTrxWithTx(ctx, dtlTrx, Tx)

		if errDtl != nil {
			Tx.Rollback()
			helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Error at CreateDetailTrx : %s", errDtl.Error()))
			return res, &helper.ErrorStruct{
				Code: fiber.StatusBadRequest,
				Err:  errDtl,
			}
		}
	}

	// commit Transaction
	Tx.Commit()
	return resTrx, nil
}
