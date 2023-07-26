package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type UsersRepository interface {
	FindByCredentials(ctx context.Context, telp string) (res daos.User, err error)
	GetById(ctx context.Context, userId string) (res daos.User, err error)
	Create(ctx context.Context, data daos.User) (res uint, err error)
	Update(ctx context.Context, userId string, data daos.User) (res string, err error)
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
	if err := alr.db.Where("notelp = ?", telp).First(&res).WithContext(ctx).Error; err != nil {
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

func (alr *UsersRepositoryImpl) GetById(ctx context.Context, userId string) (res daos.User, err error) {
	if err := alr.db.First(&res, userId).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *UsersRepositoryImpl) Update(ctx context.Context, userId string, data daos.User) (res string, err error) {
	var dataUser daos.User
	if err = alr.db.Where("id=?", userId).First((&dataUser)).WithContext(ctx).Error; err != nil {
		return "Update user failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataUser).Updates(&data).Where("id=?", userId).Error; err != nil {
		return "Update user failed", err
	}
	return res, nil
}
