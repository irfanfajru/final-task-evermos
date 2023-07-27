package repository

import (
	"context"
	"tugas_akhir_example/internal/daos"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory(ctx context.Context) (res []daos.Category, err error)
	GetCategoryById(ctx context.Context, categoryId string) (res daos.Category, err error)
	CreateCategory(ctx context.Context, data daos.Category) (res uint, err error)
	UpdateCategoryById(ctx context.Context, categoryId string, data daos.Category) (res string, err error)
	DeleteCategoryById(ctx context.Context, categoryId string) (res string, err error)
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (alr *CategoryRepositoryImpl) GetAllCategory(ctx context.Context) (res []daos.Category, err error) {
	if err := alr.db.Debug().WithContext(ctx).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *CategoryRepositoryImpl) GetCategoryById(ctx context.Context, categoryId string) (res daos.Category, err error) {
	if err := alr.db.Where("id = ?", categoryId).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (alr *CategoryRepositoryImpl) CreateCategory(ctx context.Context, data daos.Category) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (alr *CategoryRepositoryImpl) UpdateCategoryById(ctx context.Context, categoryId string, data daos.Category) (res string, err error) {
	var dataCategory daos.Category
	if err = alr.db.Where("id = ?", categoryId).First(&dataCategory).WithContext(ctx).Error; err != nil {
		return "Update category failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataCategory).Updates(&data).Where("id = ?", categoryId).Error; err != nil {
		return "Update category failed", err
	}

	return res, nil
}

func (alr *CategoryRepositoryImpl) DeleteCategoryById(ctx context.Context, categoryId string) (res string, err error) {
	var dataCategory daos.Category
	if err = alr.db.Where("id = ?", categoryId).First(&dataCategory).WithContext(ctx).Error; err != nil {
		return "Delete category failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataCategory).Delete(&dataCategory).Error; err != nil {
		return "Delete category failed", err
	}

	return res, nil
}
