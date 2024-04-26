package model

type PropertyCategory struct {
	BaseModel
	Name       string     `gorm:"size:50;type:string;not null,unique;"`
	Icon       string     `gorm:"size:1000;type:string;not null,unique;"`
	Properties []Property `gorm:"foreignKey:CategoryId"`
}

type Property struct {
	BaseModel
	Name        string           `gorm:"size:50;type:string;not null,unique;"`
	Icon        string           `gorm:"size:1000;type:string;not null,unique;"`
	Category    PropertyCategory `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CategoryId  int
	Description string `gorm:"size:1000;type:string;not null,unique;"`
	DataType    string `gorm:"size:15;type:string;not null,unique;"`
	Unit        string `gorm:"size:15;type:string;not null,unique;"`
}
