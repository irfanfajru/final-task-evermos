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

		// relation belongs to user
		User User `gorm:"foreignKey:IdUser"`

		// relation belongs to alamat pengiriman
		Alamat Alamat `gorm:"foreignKey:AlamatPengiriman"`

		// relation one to many detail trx
		DetailTrx []DetailTrx `gorm:"foreignKey:IdTrx"`
	}

	FilterTrx struct {
		KodeInvoice   string
		Limit, Offset int
	}
)
