package dto

type CreateUpdateCarModelColorRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	ColorId    int `json:"colorId" binding:"required"`
}

type CarModelColorResponse struct {
	Id       int              `json:"id"`
	Name     string           `json:"name"`
	CarModel CarModelResponse `json:"carModel"`
	Color    ColorResponse    `json:"color"`
}
