package dto

import "time"

type CreateUpdateCarModelPriceHistoryRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	Price          float32   `json:"price" binding:"required"`
	PriceAt          time.Time `json:"priceAt" binding:"required"`
}

type CarModelPriceHistoryResponse struct {
	Id                     int                            `json:"id"`
	CarModelYear               CarModelYearResponse               `json:"carModelYear"`
	Price          float32   `json:"price" binding:"required"`
	PriceAt          time.Time `json:"priceAt" binding:"required"`
}
