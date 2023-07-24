package daos

import "gorm.io/gorm"

type (
	Toko struct {
		gorm.Model
		IdUser   uint
		NamaToko string `gorm:"size:255;unique"`
		UrlFoto  string `gorm:"size:255"`
	}
)
