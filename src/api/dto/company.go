package dto

type CreateUpdateCompanyRequest struct {
	Name      string `json:"name" binding:"required,min=0,max=15"`
	CountryId int    `json:"countryId" binding:"required"`
}

type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}
