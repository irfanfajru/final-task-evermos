package dto

type (
	Produk struct {
		ID            uint         `json:"id"`
		NamaProduk    string       `json:"nama_produk"`
		Slug          string       `json:"slug"`
		HargaReseller string       `json:"harga_reseller"`
		HargaKonsumen string       `json:"harga_konsumen"`
		Stok          int          `json:"stok"`
		Deskripsi     string       `json:"deskripsi"`
		Toko          Toko         `json:"toko"`
		Category      Category     `json:"category"`
		FotoProduk    []FotoProduk `json:"photos"`
	}

	CreateProdukReq struct {
		NamaProduk    string   `form:"nama_produk" validate:"required"`
		CategoryId    string   `form:"category_id" validate:"required"`
		HargaReseller int      `form:"harga_reseller" validate:"required,min=0"`
		HargaKonsumen int      `form:"harga_konsumen" validate:"required,min=0"`
		Stok          int      `form:"stok" validate:"required,min=0"`
		Deskripsi     string   `form:"deskripsi" validate:"required"`
		FotoProduk    []string `form:"photos" validate:"required"`
	}

	UpdateProdukReq struct {
		NamaProduk string `form:"nama_produk,omitempty"`
	}

	FilterProduk struct {
		Limit      int    `query:"limit"`
		Page       int    `query:"page"`
		NamaProduk string `query:"nama_produk"`
		CategoryId uint   `query:"category_id"`
		TokoId     uint   `query:"toko_id"`
		MaxHarga   int    `query:"max_harga"`
		MinHarga   int    `query:"min_harga"`
	}

	ProdukWithPagination struct {
		Page  int      `json:"page"`
		Limit int      `json:"limit"`
		Data  []Produk `json:"data"`
	}
)
