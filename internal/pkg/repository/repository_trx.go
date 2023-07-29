package repository

import (
	"context"
	"fmt"
	"reflect"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type TrxRepository interface {
	CreateTrxWithTx(ctx context.Context, data daos.Trx, Tx *gorm.DB) (res uint, err error)
	GetAllTrx(ctx context.Context, userId string, params daos.FilterTrx) (res []daos.Trx, err error)
}

type TrxRepositoryImpl struct {
	db *gorm.DB
}

func NewTrxRepository(db *gorm.DB) TrxRepository {
	return &TrxRepositoryImpl{
		db: db,
	}
}

func (alr *TrxRepositoryImpl) CreateTrxWithTx(ctx context.Context, data daos.Trx, Tx *gorm.DB) (res uint, err error) {
	result := Tx.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (alr *TrxRepositoryImpl) GetAllTrx(ctx context.Context, userId string, params daos.FilterTrx) (res []daos.Trx, err error) {
	db := alr.db
	filter := map[string][]any{
		"kode_invoice like ?": []any{fmt.Sprintf("%%%s%%", params.KodeInvoice)},
	}

	for key, val := range filter {
		if reflect.ValueOf(val[0]).IsZero() {
			continue
		}

		db = db.Where(key, val...)
	}

	// preload
	db = db.Debug().WithContext(ctx)
	db = db.Preload("Alamat")
	db = db.Preload("DetailTrx")
	db = db.Preload("DetailTrx.LogProduk")
	db = db.Preload("DetailTrx.LogProduk.Category")
	db = db.Preload("DetailTrx.LogProduk.Produk.FotoProduk")
	db = db.Preload("DetailTrx.Toko")
	if err := db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
