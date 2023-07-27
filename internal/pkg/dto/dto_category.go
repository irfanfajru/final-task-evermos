package dto

type (
	Category struct {
		ID           uint   `json:"id"`
		NamaCategory string `json:"nama_category"`
	}

	CreateCategoryReq struct {
		NamaCategory string `json:"nama_category" validate:"required"`
	}

	UpdateCategoryReq struct {
		NamaCategory string `json:"nama_category" validate:"required"`
	}
)
