package repository

import (
	"encoding/json"
	"tugas_akhir_example/internal/pkg/dto"

	"github.com/gofiber/fiber/v2"
)

type WilayahRepository interface {
	GetAllProvince() (res []dto.Province, err error)
	GetAllRegency(provinceId string) (res []dto.Regency, err error)
	GetProvinceById(provinceId string) (res dto.Province, err error)
	GetRegencyById(provinceId string, regencyId string) (res dto.Regency, err error)
}

type WilayahRepositoryImpl struct {
	API string
}

func NewWilayahRepository() WilayahRepository {
	return &WilayahRepositoryImpl{
		API: "https://www.emsifa.com/api-wilayah-indonesia/api/",
	}
}

func (alr *WilayahRepositoryImpl) GetAllProvince() (res []dto.Province, err error) {
	agent := fiber.AcquireAgent()
	agent.Request().Header.SetMethod("GET")
	agent.Request().SetRequestURI(alr.API + "/provinces.json")
	err = agent.Parse()
	if err != nil {
		return res, err
	}

	statusCode, body, errs := agent.Bytes()
	if statusCode != 200 {
		return res, errs[0]
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (alr *WilayahRepositoryImpl) GetAllRegency(provinceId string) (res []dto.Regency, err error) {
	agent := fiber.AcquireAgent()
	agent.Request().Header.SetMethod("GET")
	agent.Request().SetRequestURI(alr.API + "/regencies/" + provinceId + ".json")
	err = agent.Parse()
	if err != nil {
		return res, err
	}

	statusCode, body, errs := agent.Bytes()
	if statusCode != 200 {
		return res, errs[0]
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (alr *WilayahRepositoryImpl) GetProvinceById(provinceId string) (res dto.Province, err error) {
	provinces, err := alr.GetAllProvince()
	if err != nil {
		return res, err
	}

	for _, v := range provinces {
		if v.ID == provinceId {
			res = v
			break
		}
	}
	return res, nil
}

func (alr *WilayahRepositoryImpl) GetRegencyById(provinceId string, regencyId string) (res dto.Regency, err error) {
	regencies, err := alr.GetAllRegency(provinceId)
	if err != nil {
		return res, err
	}

	for _, v := range regencies {
		if v.ID == regencyId {
			res = v
			break
		}
	}

	return res, nil
}
