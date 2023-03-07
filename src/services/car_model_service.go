package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Company.Country"},
				{string: "CarType"},
				{string: "Gearbox"},
				{string: "CarModelColors.Color"},
				{string: "CarModelYears.PersianYear"},
				{string: "CarModelYears.CarModelPriceHistories"},
				{string: "CarModelProperties.Property.Category"},
				{string: "CarModelImages.Image"},
				{string: "CarModelComments.User"},
			},
		},
	}
}

// Create
func (s *CarModelService) Create(ctx context.Context, req *dto.CreateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelService) Update(ctx context.Context, id int, req *dto.UpdateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelService) GetById(ctx context.Context, id int) (*dto.CarModelResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
