package dto

type CreateCarModelImageRequest struct {
	CarModelId  int  `json:"carModelId" binding:"required"`
	ImageId     int  `json:"imageId" binding:"required"`
	IsMainImage bool `json:"isMainImage"`
}

type CarModelImageResponse struct {
	Id          int              `json:"id"`
	CarModel    CarModelResponse `json:"carModel"`
	Image       FileResponse     `json:"image"`
	Description string           `json:"description"`
	IsMainImage bool             `json:"isMainImage"`
}
