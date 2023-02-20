package dto

import "time"

type CreatePersianYearRequest struct {
	PersianTitle string `json:"persianTitle" binding:"required,min=4,max=4"`
	Year int `json:"year"`
	StartAt time.Time `json:"startAt"`
	EndAt time.Time `json:"endAt"`
}

type PersianYearResponse struct {
	Id int `json:"id"`
	PersianTitle string `json:"persianTitle" binding:"required,min=4,max=4"`
	Year int `json:"year"`
	StartAt time.Time `json:"startAt"`
	EndAt time.Time `json:"endAt"`
}