package dto

import "mime/multipart"

type CreateFileRequest struct {
	File        *multipart.FileHeader `json:"file"`
	Description string                `json:"description"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
