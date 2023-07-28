package dto

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Regency struct {
	ID         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}

type District struct {
	ID        string `json:"id"`
	RegencyId string `json:"regency_id"`
	Name      string `json:"name"`
}

type Village struct {
	ID         string `json:"id"`
	DistrictId string `json:"district_id"`
	Name       string `json:"name"`
}
