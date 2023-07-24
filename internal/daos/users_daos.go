package daos

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Nama         string    `gorm:"size:255"`
		KataSandi    string    `gorm:"size:255"`
		Notelp       string    `gorm:"unique;size255"`
		TanggalLahir time.Time `gorm:"type:date"`
		JenisKelamin string    `gorm:"size:255"`
		Tentang      string    `gorm:"type:text"`
		Pekerjaan    string    `gorm:"size:255"`
		Email        string    `gorm:"unique;size:255"`
		IdProvinsi   string    `gorm:"size:255"`
		IdKota       string    `gorm:"size:255"`
		IsAdmin      bool      `gorm:"default:false"`
	}
)
