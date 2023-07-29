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
)
