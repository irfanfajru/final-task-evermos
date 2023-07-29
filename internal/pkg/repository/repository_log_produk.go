package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type LogProdukRepository interface {
	CreateLogProdukWithTx(ctx context.Context, data daos.LogProduk, Tx *gorm.DB) (res uint, err error)
}

type LogProdukRepositoryImpl struct {
	db *gorm.DB
}

func NewLogProdukRepository(db *gorm.DB) LogProdukRepository {
	return &LogProdukRepositoryImpl{
		db: db,
	}
}

func (alr *LogProdukRepositoryImpl) CreateLogProdukWithTx(ctx context.Context, data daos.LogProduk, Tx *gorm.DB) (res uint, err error) {
	result := Tx.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}
