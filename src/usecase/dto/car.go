package dto

import "time"

type CreateCarModel struct {
	Name      string
	CompanyId int
	CarTypeId int
	GearboxId int
}

type UpdateCarModel struct {
	Name      string
	CompanyId int
	CarTypeId int
	GearboxId int
}

type CarModel struct {
	IdName
	CarType            IdName
	Company            Company
	Gearbox            IdName
	CarModelColors     []CarModelColor
	CarModelYears      []CarModelYear
	CarModelImages     []CarModelImage
	CarModelProperties []CarModelProperty
	CarModelComments   []CarModelComment
}

type CreateCarModelColor struct {
	CarModelId int
	ColorId    int
}

type UpdateCarModelColor struct {
	CarModelId int
	ColorId    int
}

type CarModelColor struct {
	Id    int
	Color Color
}

type CreateCarModelYear struct {
	CarModelId    int
	PersianYearId int
}

type UpdateCarModelYear struct {
	CarModelId    int
	PersianYearId int
}

type CarModelYear struct {
	Id                     int
	PersianYear            PersianYearWithoutDate
	CarModelId             int
	CarModelPriceHistories []CarModelPriceHistory
}

type CreateCarModelPriceHistory struct {
	CarModelYearId int
	PriceAt        time.Time
	Price          float64
}

type UpdateCarModelPriceHistory struct {
	PriceAt time.Time
	Price   float64
}

type CarModelPriceHistory struct {
	Id             int
	CarModelYearId int
	PriceAt        time.Time
	Price          float64
}

type CreateCarModelImage struct {
	CarModelId  int
	ImageId     int
	IsMainImage bool
}

type UpdateCarModelImage struct {
	IsMainImage bool
}

type CarModelImage struct {
	Id          int
	CarModelId  int
	Image       File
	IsMainImage bool
}

type CreateCarModelProperty struct {
	CarModelId int
	PropertyId int
	Value      string
}

type UpdateCarModelProperty struct {
	Value string
}

type CarModelProperty struct {
	Id         int
	CarModelId int
	Property   Property
	Value      string
}

type CreateCarModelComment struct {
	CarModelId int
	UserId     int
	Message    string
}

type UpdateCarModelComment struct {
	Message string
}

type CarModelComment struct {
	Id         int
	CarModelId int
	User       User
	Message    string
}

type User struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	Email     string
}
