package daos

import "gorm.io/gorm"

type (
	Trx struct {
		gorm.Model
		IdUser           uint
		AlamatPengiriman uint
		HargaTotal       uint
		KodeInvoice      string `gorm:"size:255"`
		MethodBayar      string `gorm:"size:255"`
	}
)
