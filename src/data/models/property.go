package models

type PropertyCategory struct {
	BaseModel
	Name       string     `gorm:"size=15;type:string;not null;unique"`
	Icon       string     `gorm:"size=1000;type:string;not null;"`
	Properties []Property `gorm:"foreignKey:CategoryId"`
}

type Property struct {
	BaseModel
	Name        string           `gorm:"size=40;type:string;not null;unique"`
	Icon        string           `gorm:"size=1000;type:string;not null;"`
	Category    PropertyCategory `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CategoryId  int
	Description string `gorm:"size=200;type:string;not null;"`
	DataType    string `gorm:"size=15;type:string;not null;"`
	Unit        string `gorm:"size=15;type:string;not null;"`
}
