package dto

type CreateUpdateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId" binding:"required"`
	PersianYearId int `json:"persianYearId" binding:"required"`
}

type CarModelYearResponse struct {
	Id                     int                            `json:"id"`
	CarModel               CarModelResponse               `json:"carModel"`
	PersianYear            PersianYearResponse            `json:"persianYear"`
	CarModelPriceHistories []CarModelPriceHistoryResponse `json:"carModelPriceHistories"`
}
