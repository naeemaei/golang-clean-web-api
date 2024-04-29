package dto

import (
	"sync"
	"time"

	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CreateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type UpdateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateGearboxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type UpdateGearboxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required,min=3,max=15"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
}

type UpdateCarModelRequest struct {
	Name      string `json:"name,omitempty"`
	CompanyId int    `json:"companyId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
}

type CarModelResponse struct {
	Id                 int                        `json:"id"`
	Name               string                     `json:"name"`
	CarType            CarTypeResponse            `json:"carType"`
	Company            CompanyResponse            `json:"company"`
	Gearbox            GearboxResponse            `json:"gearbox"`
	CarModelColors     []CarModelColorResponse    `json:"carModelColors,omitempty"`
	CarModelYears      []CarModelYearResponse     `json:"carModelYears,omitempty"`
	CarModelImages     []CarModelImageResponse    `json:"carModelImages,omitempty"`
	CarModelProperties []CarModelPropertyResponse `json:"carModelProperties,omitempty"`
	CarModelComments   []CarModelCommentResponse  `json:"carModelComments,omitempty"`
}

type CreateCarModelColorRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	ColorId    int `json:"colorId" binding:"required"`
}

type UpdateCarModelColorRequest struct {
	CarModelId int `json:"carModelId,omitempty"`
	ColorId    int `json:"colorId,omitempty"`
}

type CarModelColorResponse struct {
	Id    int           `json:"id"`
	Color ColorResponse `json:"color,omitempty"`
}

type CreateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId" binding:"required"`
	PersianYearId int `json:"persianYearId" binding:"required"`
}

type UpdateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId,omitempty"`
	PersianYearId int `json:"persianYearId,omitempty"`
}

type CarModelYearResponse struct {
	Id                     int                            `json:"id"`
	PersianYear            PersianYearWithoutDateResponse `json:"persianYear,omitempty"`
	CarModelId             int                            `json:"carModelId,omitempty"`
	CarModelPriceHistories []CarModelPriceHistoryResponse `json:"carModelPriceHistories,omitempty"`
}

type CreateCarModelPriceHistoryRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	PriceAt        time.Time `json:"priceAt" binding:"required"`
	Price          float64   `json:"price" binding:"required"`
}

type UpdateCarModelPriceHistoryRequest struct {
	PriceAt time.Time `json:"priceAt,omitempty"`
	Price   float64   `json:"price,omitempty"`
}

type CarModelPriceHistoryResponse struct {
	Id             int       `json:"id"`
	CarModelYearId int       `json:"carModelYearId"`
	PriceAt        time.Time `json:"priceAt,omitempty"`
	Price          float64   `json:"price,omitempty"`
}

type CreateCarModelImageRequest struct {
	CarModelId  int  `json:"carModelId" binding:"required"`
	ImageId     int  `json:"imageId" binding:"required"`
	IsMainImage bool `json:"isMainImage"`
}

type UpdateCarModelImageRequest struct {
	IsMainImage bool `json:"isMainImage,omitempty"`
}

type CarModelImageResponse struct {
	Id          int          `json:"id"`
	CarModelId  int          `json:"carModelId,omitempty"`
	Image       FileResponse `json:"image,omitempty"`
	IsMainImage bool         `json:"isMainImage"`
}

type CreateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	PropertyId int    `json:"propertyId" binding:"required"`
	Value      string `json:"value" binding:"required,max=100"`
}

type UpdateCarModelPropertyRequest struct {
	Value string `json:"value" binding:"required,max=100"`
}

type CarModelPropertyResponse struct {
	Id         int              `json:"id"`
	CarModelId int              `json:"carModelId,omitempty"`
	Property   PropertyResponse `json:"property,omitempty"`
	Value      string           `json:"value"`
}

type CreateCarModelCommentRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	UserId     int    `json:"userId"`
	Message    string `json:"message" binding:"required,max=100"`
}

type UpdateCarModelCommentRequest struct {
	Message string `json:"message" binding:"required,max=100"`
}

type CarModelCommentResponse struct {
	Id         int          `json:"id"`
	CarModelId int          `json:"carModelId"`
	User       UserResponse `json:"user"`
	Message    string       `json:"message"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func ToCarTypeResponse(from dto.IdName) CarTypeResponse {
	return CarTypeResponse{
		Id:   from.Id,
		Name: from.Name,
	}
}

func ToCreateCarType(from CreateCarTypeRequest) dto.Name {
	return dto.Name{
		Name: from.Name,
	}
}

func ToUpdateCarType(from UpdateCarTypeRequest) dto.Name {
	return dto.Name{
		Name: from.Name,
	}
}

func ToGearboxResponse(from dto.IdName) GearboxResponse {
	return GearboxResponse{
		Id:   from.Id,
		Name: from.Name,
	}
}

func ToCreateGearbox(from CreateGearboxRequest) dto.Name {
	return dto.Name{
		Name: from.Name,
	}
}

func ToUpdateGearbox(from UpdateGearboxRequest) dto.Name {
	return dto.Name{
		Name: from.Name,
	}
}

func ToCarModelResponse(from dto.CarModel) CarModelResponse {
	colors := []CarModelColorResponse{}
	years := []CarModelYearResponse{}
	images := []CarModelImageResponse{}
	properties := []CarModelPropertyResponse{}
	comments := []CarModelCommentResponse{}

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		for _, item := range from.CarModelColors {
			colors = append(colors, ToCarModelColorResponse(item))
		}
		wg.Done()
	}()

	go func() {
		for _, item := range from.CarModelYears {
			years = append(years, ToCarModelYearResponse(item))
		}
		wg.Done()
	}()

	go func() {
		for _, item := range from.CarModelComments {
			comments = append(comments, ToCarModelCommentResponse(item))
		}
		wg.Done()
	}()

	go func() {
		for _, item := range from.CarModelImages {
			images = append(images, ToCarModelImageResponse(item))
		}
		wg.Done()
	}()

	go func() {
		for _, item := range from.CarModelProperties {
			properties = append(properties, ToCarModelPropertyResponse(item))
		}
		wg.Done()
	}()

	wg.Wait()

	return CarModelResponse{
		Id:                 from.Id,
		Name:               from.Name,
		CarType:            ToCarTypeResponse(from.CarType),
		Company:            ToCompanyResponse(from.Company),
		Gearbox:            ToGearboxResponse(from.Gearbox),
		CarModelColors:     colors,
		CarModelYears:      years,
		CarModelImages:     images,
		CarModelProperties: properties,
		CarModelComments:   comments,
	}
}

func ToCreateCarModel(from CreateCarModelRequest) dto.CreateCarModel {
	return dto.CreateCarModel{
		Name:      from.Name,
		CompanyId: from.CompanyId,
		CarTypeId: from.CarTypeId,
		GearboxId: from.GearboxId,
	}
}

func ToUpdateCarModel(from UpdateCarModelRequest) dto.UpdateCarModel {
	return dto.UpdateCarModel{
		Name:      from.Name,
		CompanyId: from.CompanyId,
		CarTypeId: from.CarTypeId,
		GearboxId: from.GearboxId,
	}
}

func ToCarModelColorResponse(from dto.CarModelColor) CarModelColorResponse {
	return CarModelColorResponse{
		Id:    from.Id,
		Color: ToColorResponse(from.Color),
	}
}

func ToCreateCarModelColor(from CreateCarModelColorRequest) dto.CreateCarModelColor {
	return dto.CreateCarModelColor{
		CarModelId: from.CarModelId,
		ColorId:    from.ColorId,
	}
}

func ToUpdateCarModelColor(from UpdateCarModelColorRequest) dto.UpdateCarModelColor {
	return dto.UpdateCarModelColor{
		CarModelId: from.CarModelId,
		ColorId:    from.ColorId,
	}
}

func ToCarModelYearResponse(from dto.CarModelYear) CarModelYearResponse {
	var histories []CarModelPriceHistoryResponse = []CarModelPriceHistoryResponse{}

	for _, item := range from.CarModelPriceHistories {
		histories = append(histories, ToCarModelPriceHistoryResponse(item))
	}
	return CarModelYearResponse{
		Id:                     from.Id,
		PersianYear:            ToPersianYearWithoutDateResponse(from.PersianYear),
		CarModelId:             from.CarModelId,
		CarModelPriceHistories: histories,
	}
}

func ToCreateCarModelYear(from CreateCarModelYearRequest) dto.CreateCarModelYear {
	return dto.CreateCarModelYear{
		CarModelId:    from.CarModelId,
		PersianYearId: from.PersianYearId,
	}
}

func ToUpdateCarModelYear(from UpdateCarModelYearRequest) dto.UpdateCarModelYear {
	return dto.UpdateCarModelYear{
		CarModelId:    from.CarModelId,
		PersianYearId: from.PersianYearId,
	}
}

func ToCarModelPriceHistoryResponse(from dto.CarModelPriceHistory) CarModelPriceHistoryResponse {
	return CarModelPriceHistoryResponse{
		Id:             from.Id,
		CarModelYearId: from.CarModelYearId,
		PriceAt:        from.PriceAt,
		Price:          from.Price,
	}
}

func ToCreateCarModelPriceHistory(from CreateCarModelPriceHistoryRequest) dto.CreateCarModelPriceHistory {
	return dto.CreateCarModelPriceHistory{
		CarModelYearId: from.CarModelYearId,
		PriceAt:        from.PriceAt,
		Price:          from.Price,
	}
}

func ToUpdateCarModelPriceHistory(from UpdateCarModelPriceHistoryRequest) dto.UpdateCarModelPriceHistory {
	return dto.UpdateCarModelPriceHistory{
		PriceAt: from.PriceAt,
		Price:   from.Price,
	}
}

func ToCarModelImageResponse(from dto.CarModelImage) CarModelImageResponse {
	return CarModelImageResponse{
		Id:          from.Id,
		CarModelId:  from.CarModelId,
		IsMainImage: from.IsMainImage,
		Image:       ToFileResponse(from.Image),
	}
}

func ToCreateCarModelImage(from CreateCarModelImageRequest) dto.CreateCarModelImage {
	return dto.CreateCarModelImage{
		CarModelId:  from.CarModelId,
		ImageId:     from.ImageId,
		IsMainImage: from.IsMainImage,
	}
}

func ToUpdateCarModelImage(from UpdateCarModelImageRequest) dto.UpdateCarModelImage {
	return dto.UpdateCarModelImage{
		IsMainImage: from.IsMainImage,
	}
}

func ToCarModelPropertyResponse(from dto.CarModelProperty) CarModelPropertyResponse {
	return CarModelPropertyResponse{
		Id:         from.Id,
		CarModelId: from.CarModelId,
		Property:   ToPropertyResponse(from.Property),
		Value:      from.Value,
	}
}

func ToCreateCarModelProperty(from CreateCarModelPropertyRequest) dto.CreateCarModelProperty {
	return dto.CreateCarModelProperty{
		CarModelId: from.CarModelId,
		PropertyId: from.PropertyId,
		Value:      from.Value,
	}
}

func ToUpdateCarModelProperty(from UpdateCarModelPropertyRequest) dto.UpdateCarModelProperty {
	return dto.UpdateCarModelProperty{
		Value: from.Value,
	}
}

func ToCarModelCommentResponse(from dto.CarModelComment) CarModelCommentResponse {
	return CarModelCommentResponse{
		Id:         from.Id,
		CarModelId: from.CarModelId,
		Message:    from.Message,
		User:       ToUserResponse(from.User),
	}
}

func ToCreateCarModelComment(from CreateCarModelCommentRequest) dto.CreateCarModelComment {
	return dto.CreateCarModelComment{
		CarModelId: from.CarModelId,
		UserId:     from.UserId,
		Message:    from.Message,
	}
}

func ToUpdateCarModelComment(from UpdateCarModelCommentRequest) dto.UpdateCarModelComment {
	return dto.UpdateCarModelComment{
		Message: from.Message,
	}
}

func ToUserResponse(from dto.User) UserResponse {
	return UserResponse{
		Id:        from.Id,
		Username:  from.Username,
		FirstName: from.FirstName,
		LastName:  from.LastName,
		Email:     from.Email,
	}
}
