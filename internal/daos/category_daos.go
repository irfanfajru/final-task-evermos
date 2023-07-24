package daos

import "gorm.io/gorm"

type (
	Category struct {
		gorm.Model
		NamaCategory string `gorm:"size:255"`

		// relation one to many to produk
		Produk []Produk `gorm:"foreignKey:IdCategory"`
	}
)
