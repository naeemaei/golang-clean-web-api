package dto

type CreateUpdateColorRequest struct {
	Name    string `json:"name" binding:"required,min=0,max=15"`
	HexCode string `json:"hexCode" binding:"required,min=0,max=6"`
}

type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}
