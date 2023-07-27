package dto

type Alamat struct {
	ID           uint   `json:"id"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	Notelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}

type AlamatFilter struct {
	JudulAlamat string `query:"judul_alamat"`
}

type CreateAlamatReq struct {
	JudulAlamat  string `json:"judul_alamat" validate:"required"`
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	Notelp       string `json:"no_telp" validate:"required"`
	DetailAlamat string `json:"detail_alamat" validate:"required"`
}

type UpdateAlamatReq struct {
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	Notelp       string `json:"no_telp" validate:"required"`
	DetailAlamat string `json:"detail_alamat" validate:"required"`
}
