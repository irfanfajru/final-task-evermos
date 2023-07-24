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
	}
)
