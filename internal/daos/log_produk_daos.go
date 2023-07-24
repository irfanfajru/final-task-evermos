package daos

import "gorm.io/gorm"

type (
	LogProduk struct {
		gorm.Model
		IdProduk      uint
		IdToko        uint
		IdCategory    uint
		NamaProduk    string `gorm:"size:255"`
		Slug          string `gorm:"size:255"`
		HargaReseller string `gorm:"size:255"`
		HargaKonsumen string `gorm:"size:255"`
		Deskripsi     string `gorm:"type:text"`
	}
)
