package dto

type LoginReq struct {
	NoTelp    string `json:"no_telp" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type LoginResp struct {
	Nama         string   `json:"nama"`
	NoTelp       string   `json:"no_telp"`
	TanggalLahir string   `json:"tanggal_lahir"`
	Tentang      string   `json:"tentang"`
	Pekerjaan    string   `json:"pekerjaan"`
	Email        string   `json:"email"`
	IdProvinsi   Province `json:"id_provinsi"`
	IdKota       Regency  `json:"id_kota"`
	Token        string   `json:"token"`
}

type RegisterReq struct {
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	Email        string `json:"email" validate:"required"`
	IdProvinsi   string `json:"id_provinsi" validate:"required"`
	IdKota       string `json:"id_kota" validate:"required"`
}

type UpdateUserReq struct {
	Nama         string `json:"nama,omitempty" `
	KataSandi    string `json:"kata_sandi,omitempty" `
	NoTelp       string `json:"no_telp,omitempty" `
	TanggalLahir string `json:"tanggal_lahir,omitempty" `
	Pekerjaan    string `json:"pekerjaan,omitempty" `
	Email        string `json:"email,omitempty" `
	IdProvinsi   string `json:"id_provinsi,omitempty" `
	IdKota       string `json:"id_kota,omitempty" `
}

type User struct {
	Nama         string   `json:"nama" `
	NoTelp       string   `json:"no_telp" `
	TanggalLahir string   `json:"tanggal_lahir" `
	Pekerjaan    string   `json:"pekerjaan" `
	Tentang      string   `json:"tentang"`
	Email        string   `json:"email" `
	IdProvinsi   Province `json:"id_provinsi" `
	IdKota       Regency  `json:"id_kota" `
}
