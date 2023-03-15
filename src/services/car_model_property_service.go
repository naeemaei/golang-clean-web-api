package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type CarModelPropertyService struct {
	base *BaseService[models.CarModelProperty, dto.CreateCarModelPropertyRequest, dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]
}

func NewCarModelPropertyService(cfg *config.Config) *CarModelPropertyService {
	return &CarModelPropertyService{
		base: &BaseService[models.CarModelProperty, dto.CreateCarModelPropertyRequest, dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Property.Category"},
			},
		},
	}
}

// Create
func (s *CarModelPropertyService) Create(ctx context.Context, req *dto.CreateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelPropertyService) Update(ctx context.Context, id int, req *dto.UpdateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelPropertyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelPropertyService) GetById(ctx context.Context, id int) (*dto.CarModelPropertyResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelPropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
