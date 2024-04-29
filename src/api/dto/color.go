package dto

import "github.com/naeemaei/golang-clean-web-api/usecase/dto"

type CreateColorRequest struct {
	Name    string `json:"name" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode" binding:"min=7,max=7"`
}

type UpdateColorRequest struct {
	Name    string `json:"name,omitempty" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode,omitempty" binding:"min=7,max=7"`
}

type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	HexCode string `json:"hexCode,omitempty"`
}

func ToColorResponse(from dto.Color) ColorResponse {
	return ColorResponse{
		Id:      from.Id,
		Name:    from.Name,
		HexCode: from.HexCode,
	}
}

func ToCreateColor(from CreateColorRequest) dto.CreateColor {
	return dto.CreateColor{
		Name:    from.Name,
		HexCode: from.HexCode,
	}
}

func ToUpdateColor(from CreateColorRequest) dto.UpdateColor {
	return dto.UpdateColor{
		Name:    from.Name,
		HexCode: from.HexCode,
	}
}
