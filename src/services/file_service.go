package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type FileService struct {
	base *BaseService[models.File, dto.CreateUpdateFileRequest, dto.CreateUpdateFileRequest, dto.FileResponse]
}

func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		base: &BaseService[models.File, dto.CreateUpdateFileRequest, dto.CreateUpdateFileRequest, dto.FileResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			// Preloads: []preload{
			// 	{Base: "Cities"}, // Inner: []preload{
			// 	// 	{Base: "File",
			// 	// 		Inner: []preload{{Base: "Cities"}},
			// 	// 	},
			// 	// },

			// 	{Base: "Companies"},
			// },
		},
	}
}

// Create
func (s *FileService) Create(ctx context.Context, req *dto.CreateUpdateFileRequest) (*dto.FileResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *FileService) Update(ctx context.Context, id int, req *dto.CreateUpdateFileRequest) (*dto.FileResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *FileService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *FileService) GetById(ctx context.Context, id int) (*dto.FileResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *FileService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.FileResponse], error) {
	return s.base.GetByFilter(ctx, req)

}
