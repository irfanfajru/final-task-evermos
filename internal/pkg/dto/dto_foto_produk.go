package dto

type (
	FotoProduk struct {
		ID       uint   `json:"id"`
		IdProduk uint   `json:"product_id"`
		Url      string `json:"url"`
	}
)
