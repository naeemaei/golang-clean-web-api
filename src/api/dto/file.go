package dto

import (
	"mime/multipart"

	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type FileFormRequest struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required" swaggerignore:"true"`
}

type UploadFileRequest struct {
	FileFormRequest
	Description string `json:"description" form:"description" binding:"required"`
}

type CreateFileRequest struct {
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}

type UpdateFileRequest struct {
	Description string `json:"description"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}

func ToFileResponse(from dto.File) FileResponse {
	return FileResponse{
		Id:          from.Id,
		Name:        from.Name,
		Directory:   from.Directory,
		Description: from.Description,
		MimeType:    from.MimeType,
	}
}

func ToCreateFile(from CreateFileRequest) dto.CreateFile {
	return dto.CreateFile{
		Name:        from.Name,
		Directory:   from.Directory,
		Description: from.Description,
		MimeType:    from.MimeType,
	}
}

func ToUpdateFile(from UpdateFileRequest) dto.UpdateFile {
	return dto.UpdateFile{
		Description: from.Description,
	}
}
