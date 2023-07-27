package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type TokoRepository interface {
	GetMyToko(ctx context.Context, userId string) (res daos.Toko, err error)
	GetAllToko(ctx context.Context, params daos.FilterToko) (res []daos.Toko, err error)
	GetTokoById(ctx context.Context, tokoId string) (res daos.Toko, err error)
	UpdateMyTokoById(ctx context.Context, userId string, tokoId string, data daos.Toko) (res string, err error)
	CreateToko(ctx context.Context, data daos.Toko) (res uint, err error)
}

type TokoRepositoryImpl struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &TokoRepositoryImpl{
		db: db,
	}
}

func (alr *TokoRepositoryImpl) GetMyToko(ctx context.Context, userId string) (res daos.Toko, err error) {
	if err := alr.db.Where("id_user = ?", userId).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (alr *TokoRepositoryImpl) GetAllToko(ctx context.Context, params daos.FilterToko) (res []daos.Toko, err error) {
	db := alr.db
	filter := map[string][]any{
		"nama_toko like ?": []any{fmt.Sprintf("%%%s%%", params.NamaToko)},
	}

	for key, val := range filter {
		db = db.Where(key, val...)
	}

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (alr *TokoRepositoryImpl) GetTokoById(ctx context.Context, tokoId string) (res daos.Toko, err error) {
	if err := alr.db.First(&res, tokoId).WithContext(ctx).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (alr *TokoRepositoryImpl) UpdateMyTokoById(ctx context.Context, userId string, tokoId string, data daos.Toko) (res string, err error) {
	var dataToko daos.Toko
	if err = alr.db.Where("id_user = ?", userId).First(&dataToko, tokoId).WithContext(ctx).Error; err != nil {
		return "Update toko failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataToko).Updates(&data).Where("id_user = ?", userId).Where("id = ?", tokoId).Error; err != nil {
		return "Update toko failed", err
	}

	return "Update toko succeed", nil
}

func (alr *TokoRepositoryImpl) CreateToko(ctx context.Context, data daos.Toko) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}
