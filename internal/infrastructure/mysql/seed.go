package mysql

import (
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/utils"
)

var usersSeed = []daos.User{
	{
		Nama:         "Admin",
		KataSandi:    utils.HashPassword("123456"),
		Notelp:       "08961231235",
		TanggalLahir: utils.ParseDate("02/01/2006"),
		JenisKelamin: "Laki-Laki",
		Tentang:      "Admin",
		Pekerjaan:    "Admin",
		Email:        "admin@mail.com",
		IdProvinsi:   "35",
		IdKota:       "3524",
		IsAdmin:      true,
	},
}
