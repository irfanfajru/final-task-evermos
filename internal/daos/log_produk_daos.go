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

		// relation belongs to produk
		Produk Produk `gorm:"foreignKey:IdProduk"`

		// relation belongs to toko
		Toko Toko `gorm:"foreignKey:IdToko"`

		// relation belongs to category
		Category Category `gorm:"foreignKey:IdCategory"`
	}
)
