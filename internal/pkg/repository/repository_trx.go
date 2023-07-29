package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type TrxRepository interface {
	CreateTrxWithTx(ctx context.Context, data daos.Trx, Tx *gorm.DB) (res uint, err error)
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
