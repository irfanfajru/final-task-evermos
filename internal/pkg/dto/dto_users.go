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
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	Email        string `json:"email" validate:"required"`
	IdProvinsi   string `json:"id_provinsi" validate:"required"`
	IdKota       string `json:"id_kota" validate:"required"`
}

type User struct {
	Nama         string `json:"nama" `
	NoTelp       string `json:"no_telp" `
	TanggalLahir string `json:"tanggal_lahir" `
	Pekerjaan    string `json:"pekerjaan" `
	Tentang      string `json:"tentang"`
	Email        string `json:"email" `
	IdProvinsi   string `json:"id_provinsi" `
	IdKota       string `json:"id_kota" `
}
