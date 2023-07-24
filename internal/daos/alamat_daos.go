package daos

import "gorm.io/gorm"

type (
	Alamat struct {
		gorm.Model
		IdUser       uint
		JudulAlamat  string `gorm:"size:255"`
		NamaPenerima string `gorm:"size:255"`
		Notelp       string `gorm:"size:255"`
		DetailAlamat string `gorm:"size:255"`

		// relation belongsTo user
		User User `gorm:"foreignKey:IdUser"`
	}
)
