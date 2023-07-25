package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type UsersRepository interface {
	FindByCredentials(ctx context.Context, telp string) (res daos.User, err error)
	Create(ctx context.Context, data daos.User) (res uint, err error)
}

type UsersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{
		db: db,
	}
}

func (alr *UsersRepositoryImpl) FindByCredentials(ctx context.Context, telp string) (res daos.User, err error) {
	if err := alr.db.Where("notelp", telp).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *UsersRepositoryImpl) Create(ctx context.Context, data daos.User) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}
