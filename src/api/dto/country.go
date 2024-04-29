package dto

import "github.com/naeemaei/golang-clean-web-api/usecase/dto"

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type CountryResponse struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	Cities    []CityResponse    `json:"cities,omitempty"`
	Companies []CompanyResponse `json:"companies,omitempty"`
}

func ToCountryResponse(from dto.Country) CountryResponse {
	return CountryResponse{
		Id:   from.Id,
		Name: from.Name,
	}
}

func ToCreateUpdateCountry(from CreateUpdateCountryRequest) dto.Name {
	return dto.Name{
		Name: from.Name,
	}
}
