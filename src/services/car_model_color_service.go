package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.CreateUpdateCarModelColorRequest, dto.CreateUpdateCarModelColorRequest, dto.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: &BaseService[models.CarModelColor, dto.CreateUpdateCarModelColorRequest, dto.CreateUpdateCarModelColorRequest, dto.CarModelColorResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads : []preload{
				{string :"CarModel.Company"},
				{string :"CarModel.CarType"},
				{string :"CarModel.Gearbox"},
				{string :"CarModel.CarModelProperties.Property.Category"},
				{string :"CarModel.CarModelYears.PersianYear"},
				{string :"CarModel.CarModelImages.Image"},
			},
		},
	}
}

// Create
func (s *CarModelColorService) Create(ctx context.Context, req *dto.CreateUpdateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelColorService) Update(ctx context.Context, id int, req *dto.CreateUpdateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelColorService) GetById(ctx context.Context, id int) (*dto.CarModelColorResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelColorService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelColorResponse], error) {
	return s.base.GetByFilter(ctx, req)

}
