package dto

type (
	Toko struct {
		ID       uint   `json:"id"`
		NamaToko string `json:"nama_toko"`
		UrlFoto  string `json:"url_foto"`
	}

	FilterToko struct {
		Page     int    `query:"page"`
		Limit    int    `query:"limit"`
		NamaToko string `query:"nama"`
	}

	UpdateTokoReq struct {
		NamaToko string `form:"nama_toko, omitempty"`
		Photo    string `form:"photo, omitempty"`
	}

	MyToko struct {
		ID       uint   `json:"id"`
		NamaToko string `json:"nama_toko"`
		UrlFoto  string `json:"url_foto"`
		UserId   string `json:"user_id"`
	}

	TokoWithPagination struct {
		Page  int    `json:"page"`
		Limit int    `json:"limit"`
		Data  []Toko `json:"data"`
	}
)
