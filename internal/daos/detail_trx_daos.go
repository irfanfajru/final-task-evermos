package daos

import (
	"gorm.io/gorm"
)

type (
	DetailTrx struct {
		gorm.Model
		IdTrx       uint
		IdLogProduk uint
		IdToko      uint
		Kuantitas   uint
		HargaTotal  uint

		// relation belongs to log produk
		LogProduk LogProduk `gorm:"foreignKey:IdLogProduk"`

		// relation belongs to toko
		Toko Toko `gorm:"foreignKey:IdToko"`
	}
)
