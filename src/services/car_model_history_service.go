package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type CarModelPriceHistoryService struct {
	base *BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryRequest,
	 dto.UpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]
}

func NewCarModelPriceHistoryService(cfg *config.Config) *CarModelPriceHistoryService {
	return &CarModelPriceHistoryService{
		base: &BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryRequest, 
		dto.UpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "CarModelYear.CarModel.Company"},
				{string: "CarModelYear.CarModel.CarType"},
				{string: "CarModelYear.CarModel.Gearbox"},
				{string: "CarModelYear.CarModel.CarModelProperties.Property.Category"},
				{string: "CarModelYear.CarModel.CarModelYears.PersianYear"},
				{string: "CarModelYear.CarModel.CarModelImages.Image"},
				{string: "CarModelYear.CarModel.CarModelColors.Color"},
			},
		},
	}
}

// Create
func (s *CarModelPriceHistoryService) Create(ctx context.Context, req *dto.CreateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelPriceHistoryService) Update(ctx context.Context, id int, req *dto.UpdateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelPriceHistoryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelPriceHistoryService) GetById(ctx context.Context, id int) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelPriceHistoryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPriceHistoryResponse], error) {
	return s.base.GetByFilter(ctx, req)

}
