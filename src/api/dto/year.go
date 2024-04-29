package dto

import (
	"time"

	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CreatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle" binding:"min=4,max=4"`
	Year         int       `json:"year"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

type UpdatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle,omitempty" binding:"min=4,max=4"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type PersianYearResponse struct {
	Id           int       `json:"id"`
	PersianTitle string    `json:"persianTitle,omitempty"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type PersianYearWithoutDateResponse struct {
	Id           int    `json:"id"`
	PersianTitle string `json:"persianTitle,omitempty"`
	Year         int    `json:"year,omitempty"`
}

func ToPersianYearResponse(from dto.PersianYear) PersianYearResponse {
	return PersianYearResponse{
		Id:           from.Id,
		PersianTitle: from.PersianTitle,
		Year:         from.Year,
		StartAt:      from.StartAt,
		EndAt:        from.EndAt,
	}
}

func ToPersianYearWithoutDateResponse(from dto.PersianYearWithoutDate) PersianYearWithoutDateResponse {
	return PersianYearWithoutDateResponse{
		Id:           from.Id,
		PersianTitle: from.PersianTitle,
		Year:         from.Year,
	}
}

func ToCreatePersianYear(from CreatePersianYearRequest) dto.CreatePersianYear {
	return dto.CreatePersianYear{
		PersianTitle: from.PersianTitle,
		Year:         from.Year,
		StartAt:      from.StartAt,
		EndAt:        from.EndAt,
	}
}

func ToUpdatePersianYear(from UpdatePersianYearRequest) dto.UpdatePersianYear {
	return dto.UpdatePersianYear{
		PersianTitle: from.PersianTitle,
		Year:         from.Year,
		StartAt:      from.StartAt,
		EndAt:        from.EndAt,
	}
}
