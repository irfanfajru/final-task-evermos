package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type DetailTrxRepository interface {
	CreateDetailTrxWithTx(ctx context.Context, data daos.DetailTrx, Tx *gorm.DB) (res uint, err error)
}

type DetailTrxRepositoryImpl struct {
	db *gorm.DB
}

func NewDetailTrxRepository(db *gorm.DB) DetailTrxRepository {
	return &DetailTrxRepositoryImpl{
		db: db,
	}
}

func (alr *DetailTrxRepositoryImpl) CreateDetailTrxWithTx(ctx context.Context, data daos.DetailTrx, Tx *gorm.DB) (res uint, err error) {
	result := Tx.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}
	return data.ID, nil
}
