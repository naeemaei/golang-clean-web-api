package dto

type CreateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	PropertyId int    `json:"propertyId" binding:"required"`
	Value      string `json:"value" binding:"required"`
}

type CarModelPropertyResponse struct {
	Id       int              `json:"id"`
	CarModel CarModelResponse `json:"carModel"`
	Property PropertyResponse `json:"property"`
	Value    string           `json:"value" binding:"required"`
}
