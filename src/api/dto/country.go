package dto

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type CountryResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Cities []CityResponse `json:"cities,omitempty"`
}

type CreateUpdateCityRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}
