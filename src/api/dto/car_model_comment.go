package dto

type CreateCarModelCommentRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	UserId     int    `json:"userId" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

type CarModelCommentResponse struct {
	Id       int              `json:"id"`
	CarModel CarModelResponse `json:"carModel"`
	User     CarModelResponse `json:"user"`
	Property PropertyResponse `json:"property"`
	Message  string           `json:"message" binding:"required"`
}
