package dto

import (
	"mime/multipart"
	"time"
)

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type CountryResponse struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	Cities    []CityResponse    `json:"cities,omitempty"`
	Companies []CompanyResponse `json:"companies,omitempty"`
}

type CreateUpdateCityRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

type CreateUpdateFileRequest struct {
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	MimeType    string `json:"mimeType"`
	Description string `json:"description"`
}

type UploadFileRequest struct {
	FileFormRequest
	Description string `json:"description" form:"description"`
}

type FileFormRequest struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required" swaggerignore:"true"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Directory   string `json:"directory"`
	MimeType    string `json:"mimeType"`
}

type CreateUpdatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle" binding:"required,min=4,max=4"`
	Year         int       `json:"year"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

type PersianYearResponse struct {
	Id           int       `json:"id"`
	PersianTitle string    `json:"persianTitle" binding:"required,min=4,max=4"`
	Year         int       `json:"year"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

type CreateUpdateColorRequest struct {
	Name    string `json:"name" binding:"required,min=0,max=15"`
	HexCode string `json:"hexCode" binding:"required,min=0,max=6"`
}

type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}

type CreateUpdateCompanyRequest struct {
	Name      string `json:"name" binding:"required,min=0,max=15"`
	CountryId int    `json:"countryId" binding:"required"`
}

type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}