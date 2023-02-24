package dto

import "time"

type CreateUpdateCarTypeRequest struct {
	Name string `json:"name" binding:"required,min=0,max=15"`
}

type CarTypeResponse struct {
	Id        int                `json:"id"`
	Name      string             `json:"name"`
	//CarModels []CarModelResponse `json:"carModels,omitempty"`
}

type CreateUpdateGearboxRequest struct {
	Name string `json:"name" binding:"required,min=0,max=15"`
}

type GearboxResponse struct {
	Id        int                `json:"id"`
	Name      string             `json:"name"`
	//CarModels []CarModelResponse `json:"carModels,omitempty"`
}

type CreateUpdateCarModelRequest struct {
	Name      string `json:"name" binding:"required,min=0,max=15"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
}

type CarModelResponse struct {
	Id                 int                        `json:"id"`
	Name               string                     `json:"name"`
	Company            CompanyResponse            `json:"company,omitempty"`
	CarType            CarTypeResponse            `json:"carType,omitempty"`
	Gearbox            GearboxResponse            `json:"gearbox,omitempty"`
	CarModelProperties []CarModelPropertyResponse `json:"carModelProperties,omitempty"`
	CarModelColors     []CarModelColorResponse    `json:"carModelColors,omitempty"`
	CarModelYears      []CarModelYearResponse     `json:"carModelYears,omitempty"`
	CarModelImages     []CarModelImageResponse    `json:"carModelImages,omitempty"`
	CarModelComments   []CarModelCommentResponse  `json:"carModelComments,omitempty"`
}

type CreateUpdateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId" binding:"required"`
	PersianYearId int `json:"persianYearId" binding:"required"`
}

type CarModelYearResponse struct {
	Id                     int                            `json:"id"`
	CarModel               CarModelResponse               `json:"carModel,omitempty"`
	PersianYear            PersianYearResponse            `json:"persianYear,omitempty"`
	CarModelPriceHistories []CarModelPriceHistoryResponse `json:"carModelPriceHistories,omitempty"`
}

type CreateUpdateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	PropertyId int    `json:"propertyId" binding:"required"`
	Value      string `json:"value" binding:"required"`
}

type CarModelPropertyResponse struct {
	Id       int              `json:"id"`
	//CarModel CarModelResponse `json:"carModel,omitempty"`
	Property PropertyResponse `json:"property,omitempty"`
	Value    string           `json:"value"`
}

type CreateUpdateCarModelPriceHistoryRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	Price          float32   `json:"price" binding:"required"`
	PriceAt        time.Time `json:"priceAt" binding:"required"`
}

type CarModelPriceHistoryResponse struct {
	Id           int                  `json:"id"`
	//CarModelYear CarModelYearResponse `json:"carModelYear,omitempty"`
	Price        float32              `json:"price" binding:"required"`
	PriceAt      time.Time            `json:"priceAt" binding:"required"`
}

type CreateUpdateCarModelImageRequest struct {
	CarModelId  int  `json:"carModelId" binding:"required"`
	ImageId     int  `json:"imageId" binding:"required"`
	IsMainImage bool `json:"isMainImage"`
}

type CarModelImageResponse struct {
	Id          int              `json:"id"`
	//CarModel    CarModelResponse `json:"carModel,omitempty"`
	Image       FileResponse     `json:"image,omitempty"`
	Description string           `json:"description"`
	IsMainImage bool             `json:"isMainImage"`
}

type CreateUpdateCarModelCommentRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	UserId     int    `json:"userId" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

type CarModelCommentResponse struct {
	Id       int              `json:"id"`
	//CarModel CarModelResponse `json:"carModel,omitempty"`
	User     UserResponse     `json:"user"`
	Message  string           `json:"message" binding:"required"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"message"`
}

type CreateUpdateCarModelColorRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	ColorId    int `json:"colorId" binding:"required"`
}

type CarModelColorResponse struct {
	Id       int              `json:"id"`
	Name     string           `json:"name"`
	//CarModel CarModelResponse `json:"carModel,omitempty"`
	Color    ColorResponse    `json:"color,omitempty"`
}