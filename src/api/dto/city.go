package dto

import "github.com/naeemaei/golang-clean-web-api/usecase/dto"

type CreateCityRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCityRequest struct {
	Name      string `json:"name,omitempty" binding:"alpha,min=3,max=20"`
	CountryId int    `json:"countryId,omitempty"`
}
type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

func ToCityResponse(from dto.City) CityResponse {
	return CityResponse{
		Id:      from.Id,
		Name:    from.Name,
		Country: ToCountryResponse(from.Country),
	}
}

func ToCreateCity(from CreateCityRequest) dto.CreateCity {
	return dto.CreateCity{
		Name:      from.Name,
		CountryId: from.CountryId,
	}
}

func ToUpdateCity(from UpdateCityRequest) dto.UpdateCity {
	return dto.UpdateCity{
		Name:      from.Name,
		CountryId: from.CountryId,
	}
}
