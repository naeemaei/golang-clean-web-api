package dto

type CreateUpdateGearboxRequest struct {
	Name string `json:"name" binding:"required,min=0,max=15"`
}

type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
