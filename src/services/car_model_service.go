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
	base *BaseService[models.CarModel, dto.CreateUpdateCarModelRequest, dto.CreateUpdateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dto.CreateUpdateCarModelRequest, dto.CreateUpdateCarModelRequest, dto.CarModelResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads : []string{
				"Company",
				"CarType",
				"Gearbox",
				"CarModelProperties.Property.Category",
				"CarModelColors.Color",
				"CarModelYears.PersianYear.CarModelPriceHistories",
				"CarModelImages.Image",
				"CarModelComments.User",
			},  
		},
	}
}

// Create
func (s *CarModelService) Create(ctx context.Context, req *dto.CreateUpdateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelService) Update(ctx context.Context, id int, req *dto.CreateUpdateCarModelRequest) (*dto.CarModelResponse, error) {
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

func (s *CarModelService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelResponse], error) {
	return s.base.GetByFilter(ctx, req)

}
