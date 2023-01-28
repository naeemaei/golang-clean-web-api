package models

type City struct {
	BaseModel
	Name      string `gorm:"size:10;type:string;not null;"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
}
