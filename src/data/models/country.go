package models

type Country struct {
	BaseModel
	Name   string `gorm:"size:15;type:string;not null;"`
	Cities *[]City
}
