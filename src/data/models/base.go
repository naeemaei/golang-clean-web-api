package models

import "time"

type Country struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null;unique"`
	Cities    []City
	Companies []Company
}

type City struct {
	BaseModel
	Name      string `gorm:"size:10;type:string;not null;unique"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
}

type PersianYear struct {
	BaseModel
	PersianTitle string    `gorm:"type:string;size:10;not null;unique"`
	Year         int       `gorm:"type:int;uniqueIndex;not null;unique"`
	StartAt      time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	EndAt        time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
	CarModelYears []CarModelYear 
}

type Color struct {
	BaseModel
	Name    string `gorm:"size=15;type:string;not null;unique" json:"name"`
	HexCode string `gorm:"size=6;type:string;not null;unique" json:"hexCode"`
	CarModelColors []CarModelColor
}

type File struct {
	BaseModel
	Name        string `gorm:"size=50;type:string;not null;" json:"name"`
	Directory   string `gorm:"size=50;type:string;not null;" json:"directory"`
	Description string `gorm:"size=500;type:string;not null;" json:"description"`
	MimeType    string `gorm:"size=6;type:string;not null;" json:"mimeType"`
}
