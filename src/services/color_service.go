package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type ColorService struct {
	base *BaseService[models.Color, dto.CreateUpdateColorRequest, dto.CreateUpdateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: &BaseService[models.Color, dto.CreateUpdateColorRequest, dto.CreateUpdateColorRequest, dto.ColorResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{string: "CarModelColors"}},
		},
	}
}

// Create
func (s *ColorService) Create(ctx context.Context, req *dto.CreateUpdateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *ColorService) Update(ctx context.Context, id int, req *dto.CreateUpdateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *ColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *ColorService) GetById(ctx context.Context, id int) (*dto.ColorResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *ColorService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.ColorResponse], error) {
	return s.base.GetByFilter(ctx, req)

}
