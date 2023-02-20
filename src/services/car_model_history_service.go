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
	base *BaseService[models.CarModelPriceHistory, dto.CreateUpdateCarModelPriceHistoryRequest, dto.CreateUpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]
}

func NewCarModelPriceHistoryService(cfg *config.Config) *CarModelPriceHistoryService {
	return &CarModelPriceHistoryService{
		base: &BaseService[models.CarModelPriceHistory, dto.CreateUpdateCarModelPriceHistoryRequest, dto.CreateUpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{Base: "CarModelYear",
				Inner: []preload{{Base: "CarModel", Inner: []preload{
					{Base: "Company"},
					{Base: "CarType"},
					{Base: "Gearbox"},
					{Base: "CarModelProperties",
						Inner: []preload{{Base: "Property",
							Inner: []preload{{Base: "Category"}},
						}},
					}, {Base: "CarModelYears",
						Inner: []preload{{Base: "PersianYear"} /*{Base: "CarModelPriceHistories"}*/},
					}, {Base: "CarModelColors",
						Inner: []preload{{Base: "Color"}},
					},
				},
				},
				}},
			},
		},
	}
}

// Create
func (s *CarModelPriceHistoryService) Create(ctx context.Context, req *dto.CreateUpdateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelPriceHistoryService) Update(ctx context.Context, id int, req *dto.CreateUpdateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
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
