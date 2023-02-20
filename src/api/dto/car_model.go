package dto

type CreateUpdateCarModelRequest struct {
	Name      string `json:"name" binding:"required,min=0,max=15"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
}

type CarModelResponse struct {
	Id                 int                        `json:"id"`
	Name               string                     `json:"name"`
	Company            CompanyResponse            `json:"company"`
	CarType            CarTypeResponse            `json:"carType"`
	Gearbox            GearboxResponse            `json:"gearbox"`
	CarModelColors     []CarModelColorResponse    `json:"carModelColors"`
	CarModelYears      []CarModelYearResponse     `json:"carModelYears"`
	CarModelImages     []CarModelImageResponse    `json:"carModelImages"`
	CarModelComments   []CarModelCommentResponse  `json:"carModelComments"`
	CarModelProperties []CarModelPropertyResponse `json:"carModelProperties"`
}
