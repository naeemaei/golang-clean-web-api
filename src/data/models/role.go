package models

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;size:10;not null,unique"`
	UserRoles *[]UserRole
}
