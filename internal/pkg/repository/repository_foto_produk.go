package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type FotoProdukRepository interface {
	CreateFotoProduk(ctx context.Context, data daos.FotoProduk) (res uint, err error)
}

type FotoProdukRepositoryImpl struct {
	db *gorm.DB
}

func NewFotoProdukRepository(db *gorm.DB) FotoProdukRepository {
	return &FotoProdukRepositoryImpl{
		db: db,
	}
}

func (alr *FotoProdukRepositoryImpl) CreateFotoProduk(ctx context.Context, data daos.FotoProduk) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}
