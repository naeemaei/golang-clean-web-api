package dto

import (
	"time"
)

type IdName struct {
	Id   int
	Name string
}

type Country struct {
	Id        int
	Name      string
	Cities    []City
	Companies []Company
}

type CreateCity struct {
	Name      string
	CountryId int
}

type UpdateCity struct {
	Name      string
	CountryId int
}
type City struct {
	Id      int
	Name    string
	Country Country
}

type CreateFile struct {
	Name        string
	Directory   string
	Description string
	MimeType    string
}

type UpdateFile struct {
	Description string
}

type File struct {
	Id          int
	Name        string
	Directory   string
	Description string
	MimeType    string
}

type CreateCompany struct {
	Name      string
	CountryId int
}

type UpdateCompany struct {
	Name      string
	CountryId int
}
type Company struct {
	Id      int
	Name    string
	Country Country
}

type CreateColor struct {
	Name    string `json:"name" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode" binding:"min=7,max=7"`
}

type UpdateColor struct {
	Name    string `json:"name,omitempty" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode,omitempty" binding:"min=7,max=7"`
}

type Color struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	HexCode string `json:"hexCode,omitempty"`
}

type CreatePersianYear struct {
	PersianTitle string
	Year         int
	StartAt      time.Time
	EndAt        time.Time
}

type UpdatePersianYear struct {
	PersianTitle string
	Year         int
	StartAt      time.Time
	EndAt        time.Time
}

type PersianYear struct {
	Id           int
	PersianTitle string
	Year         int
	StartAt      time.Time
	EndAt        time.Time
}

type PersianYearWithoutDate struct {
	Id           int
	PersianTitle string
	Year         int
}
