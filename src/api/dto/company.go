package dto

import "github.com/naeemaei/golang-clean-web-api/usecase/dto"

type CreateCompanyRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCompanyRequest struct {
	Name      string `json:"name,omitempty" binding:"alpha,min=3,max=20"`
	CountryId int    `json:"countryId,omitempty"`
}
type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

func ToCompanyResponse(from dto.Company) CompanyResponse {
	return CompanyResponse{
		Id:      from.Id,
		Name:    from.Name,
		Country: ToCountryResponse(from.Country),
	}
}

func ToCreateCompany(from CreateCompanyRequest) dto.CreateCompany {
	return dto.CreateCompany{
		Name:      from.Name,
		CountryId: from.CountryId,
	}
}

func ToUpdateCompany(from UpdateCompanyRequest) dto.UpdateCompany {
	return dto.UpdateCompany{
		Name:      from.Name,
		CountryId: from.CountryId,
	}
}
