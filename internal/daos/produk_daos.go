package daos

import "gorm.io/gorm"

type (
	Produk struct {
		gorm.Model
		IdToko        uint
		IdCategory    uint
		NamaProduk    string `gorm:"size:255"`
		Slug          string `gorm:"size:255"`
		HargaReseller string `gorm:"size:255"`
		HargaKonsumen string `gorm:"size:255"`
		Stok          uint
		Deskripsi     string `grom:"type:text"`
	}
)