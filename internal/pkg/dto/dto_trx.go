package dto

type (
	CreateTrxReq struct {
		MethodBayar string `json:"method_bayar" validate:"required"`
		AlamatKirim uint   `json:"alamat_kirim" validate:"required"`
		DetailTrx   []struct {
			ProdukId  uint `json:"product_id" validate:"required"`
			Kuantitas uint `json:"kuantitas" validate:"required"`
		} `json:"detail_trx" validate:"required"`
	}

	FilterTrx struct {
		Limit       int    `query:"limit"`
		Page        int    `query:"page"`
		KodeInvoice string `query:"search"`
	}

	Trx struct {
		ID          uint        `json:"id"`
		HargaTotal  int         `json:"harga_total"`
		KodeInvoice string      `json:"kode_invoice"`
		MethodBayar string      `json:"method_bayar"`
		AlamatKirim Alamat      `json:"alamat_kirim"`
		DetailTrx   []DetailTrx `json:"detail_trx"`
	}

	TrxWithPagination struct {
		Page  int   `json:"page"`
		Limit int   `json:"limit"`
		Data  []Trx `json:"data"`
	}

	DetailTrx struct {
		Produk     Produk `json:"produk"`
		Toko       Toko   `json:"toko"`
		Kuantitas  int    `json:"kuantitas"`
		HargaTotal int    `json:"harga_total"`
	}
)
