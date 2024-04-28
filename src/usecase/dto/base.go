package dto

import (
	"time"
)

type IdName struct {
	Id   int
	Name string
}
type Name struct {
	Name string
}

type Country struct {
	IdName
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
	IdName
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
	IdName
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
	IdName
	Country Country
}

type CreateColor struct {
	Name    string 
	HexCode string 
}

type UpdateColor struct {
	Name    string  
	HexCode string  
}

type Color struct {
	IdName
	HexCode string 
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
