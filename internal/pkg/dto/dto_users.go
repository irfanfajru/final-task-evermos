package dto

type LoginReq struct {
	NoTelp    string `json:"no_telp" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type LoginResp struct {
	Nama         string `json:"nama"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_lahir"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	Token        string `json:"token"`
}
