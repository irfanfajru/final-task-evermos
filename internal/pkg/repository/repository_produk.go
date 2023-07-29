package repository

import (
	"context"
	"fmt"
	"reflect"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	CreateProduk(ctx context.Context, data daos.Produk) (res uint, err error)
	GetAllProduk(ctx context.Context, params daos.FilterProduk) (res []daos.Produk, err error)
	GetProdukById(ctx context.Context, produkId string) (res daos.Produk, err error)
	UpdateProdukById(ctx context.Context, tokoId uint, produkId string, data daos.Produk) (res string, err error)
	UpdateProdukByIdWithTx(ctx context.Context, tokoId uint, produkId string, data daos.Produk, Tx *gorm.DB) (res string, err error)
	DeleteProdukById(ctx context.Context, tokoId uint, produkId string) (res string, err error)
}

type ProdukRepositoryImpl struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &ProdukRepositoryImpl{
		db: db,
	}
}

func (alr *ProdukRepositoryImpl) CreateProduk(ctx context.Context, data daos.Produk) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (alr *ProdukRepositoryImpl) GetAllProduk(ctx context.Context, params daos.FilterProduk) (res []daos.Produk, err error) {
	db := alr.db
	filter := map[string][]any{
		"nama_produk like ? ": []any{fmt.Sprintf("%%%s%%", params.NamaProduk)},
		"id_category = ?":     []any{params.CategoryId},
		"id_toko = ?":         []any{params.TokoId},
		"harga_konsumen >= ? or harga_reseller >= ?": []any{params.MinHarga, params.MinHarga},
		"harga_konsumen <= ? or harga_reseller <= ?": []any{params.MaxHarga, params.MaxHarga},
	}

	for key, val := range filter {
		if reflect.ValueOf(val[0]).IsZero() {
			continue
		}
		db = db.Where(key, val...)
	}

	if err := db.Debug().WithContext(ctx).Preload("Toko").Preload("Category").Preload("FotoProduk").Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *ProdukRepositoryImpl) GetProdukById(ctx context.Context, produkId string) (res daos.Produk, err error) {
	if err := alr.db.Preload("Toko").Preload("Category").Preload("FotoProduk").First(&res, produkId).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *ProdukRepositoryImpl) UpdateProdukById(ctx context.Context, tokoId uint, produkId string, data daos.Produk) (res string, err error) {
	var dataProduk daos.Produk
	if err = alr.db.Where("id_toko = ?", tokoId).First(&dataProduk, produkId).WithContext(ctx).Error; err != nil {
		return "Update produk failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataProduk).Updates(&data).Where("id_toko = ?", tokoId).Where("id = ?", produkId).Error; err != nil {
		return "Update produk failed", err
	}

	return res, nil
}

func (alr *ProdukRepositoryImpl) DeleteProdukById(ctx context.Context, tokoId uint, produkId string) (res string, err error) {
	var dataProduk daos.Produk
	if err = alr.db.Where("id_toko = ?", tokoId).First(&dataProduk, produkId).WithContext(ctx).Error; err != nil {
		return "Delete produk failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataProduk).Delete(&dataProduk).Error; err != nil {
		return "Delete produk failed", err
	}
	return res, nil
}

func (alr *ProdukRepositoryImpl) UpdateProdukByIdWithTx(ctx context.Context, tokoId uint, produkId string, data daos.Produk, Tx *gorm.DB) (res string, err error) {
	var dataProduk daos.Produk
	if err = Tx.Where("id_toko = ?", tokoId).First(&dataProduk, produkId).WithContext(ctx).Error; err != nil {
		return "Update produk failed", gorm.ErrRecordNotFound
	}

	if err := Tx.Model(dataProduk).Updates(&data).Where("id_toko = ?", tokoId).Where("id = ?", produkId).Error; err != nil {
		return "Update produk failed", err
	}

	return res, nil
}
