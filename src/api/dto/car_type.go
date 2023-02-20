package dto

type CreateUpdateCarTypeRequest struct {
	Name string `json:"name" binding:"required,min=0,max=15"`
}

type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
