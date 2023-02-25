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
	base *BaseService[models.CarModelProperty, dto.CreateUpdateCarModelPropertyRequest, dto.CreateUpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]
}

func NewCarModelPropertyService(cfg *config.Config) *CarModelPropertyService {
	return &CarModelPropertyService{
		base: &BaseService[models.CarModelProperty, dto.CreateUpdateCarModelPropertyRequest, dto.CreateUpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "CarModel.Company"},
				{string: "CarModel.CarType"},
				{string: "CarModel.Gearbox"},
				{string: "CarModel.CarModelColors.Color"},
				{string: "CarModel.CarModelYears.PersianYear"},
				{string: "CarModel.CarModelImages.Image"},
			},
		},
	}
}

// Create
func (s *CarModelPropertyService) Create(ctx context.Context, req *dto.CreateUpdateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelPropertyService) Update(ctx context.Context, id int, req *dto.CreateUpdateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
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

func (s *CarModelPropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)

}
