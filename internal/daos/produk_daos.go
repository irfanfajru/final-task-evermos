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

		// relation belongs to toko
		Toko Toko `gorm:"foreignKey:IdToko"`

		// relation belongs to category
		Category Category `gorm:"foreignKey:IdCategory"`

		// relation one to many to foto produk
		FotoProduk []FotoProduk `gorm:"foreignKey:IdProduk"`
	}

	FilterProduk struct {
		Limit, Offset int
		NamaProduk    string
		CategoryId    uint
		TokoId        uint
		MaxHarga      int
		MinHarga      int
	}
)
