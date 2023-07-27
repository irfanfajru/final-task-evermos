package repository

import (
	"context"
	"fmt"
	"log"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type AlamatRepository interface {
	GetMyAlamat(ctx context.Context, userId string, params daos.FilterAlamat) (res []daos.Alamat, err error)
	GetMyAlamatById(ctx context.Context, userId string, alamatId string) (res daos.Alamat, err error)
	CreateAlamat(ctx context.Context, data daos.Alamat) (res uint, err error)
	UpdateAlamat(ctx context.Context, alamatId string, userId string, data daos.Alamat) (res string, err error)
	DeleteAlamat(ctx context.Context, alamatId string, userId string) (res string, err error)
}

type AlamatRepositoryImpl struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &AlamatRepositoryImpl{
		db: db,
	}
}

func (alr *AlamatRepositoryImpl) GetMyAlamat(ctx context.Context, userId string, params daos.FilterAlamat) (res []daos.Alamat, err error) {
	db := alr.db
	filter := map[string][]any{
		"judul_alamat like ?": []any{fmt.Sprintf("%%%s%%", params.JudulAlamat)},
	}
	log.Println(filter)
	for key, val := range filter {
		db = db.Where(key, val...)
	}

	if err := db.Debug().WithContext(ctx).Where("id_user=?", userId).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *AlamatRepositoryImpl) GetMyAlamatById(ctx context.Context, userId string, alamatId string) (res daos.Alamat, err error) {
	if err := alr.db.Where("id_user = ?", userId).First(&res, alamatId).WithContext(ctx).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (alr *AlamatRepositoryImpl) CreateAlamat(ctx context.Context, data daos.Alamat) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}
	return data.ID, nil
}

func (alr *AlamatRepositoryImpl) UpdateAlamat(ctx context.Context, alamatId string, userId string, data daos.Alamat) (res string, err error) {
	var dataAlamat daos.Alamat
	if err = alr.db.Where("id = ?", alamatId).Where("id_user = ?", userId).First(&dataAlamat).WithContext(ctx).Error; err != nil {
		return "Update alamat failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataAlamat).Updates(&data).Where("id = ?", alamatId).Error; err != nil {
		return "Update alamat failed", err
	}

	return res, nil
}

func (alr *AlamatRepositoryImpl) DeleteAlamat(ctx context.Context, alamatId string, userId string) (res string, err error) {
	var dataAlamat daos.Alamat
	if err = alr.db.Where("id = ?", alamatId).Where("id_user = ?", userId).First(&dataAlamat).WithContext(ctx).Error; err != nil {
		return "Delete alamat failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataAlamat).Delete(&dataAlamat).Error; err != nil {
		return "Delete alamat failed", err
	}

	return res, nil
}
