package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dto.CreateUpdateCarModelYearRequest, dto.CreateUpdateCarModelYearRequest, dto.CarModelYearResponse]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: &BaseService[models.CarModelYear, dto.CreateUpdateCarModelYearRequest, dto.CreateUpdateCarModelYearRequest, dto.CarModelYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{Base: "CarModel", Inner: []preload{
					{Base: "Company"},
					{Base: "CarType"},
					{Base: "Gearbox"},
					{Base: "CarModelProperties",
						Inner: []preload{{Base: "Property",
							Inner: []preload{{Base: "Category"}},
						}},
					}, {Base: "CarModelColors",
						Inner: []preload{{Base: "Color"}},
					}, {Base: "CarModelImages",
						Inner: []preload{{Base: "Image"}},
					},
					//  {Base: "CarModelComments",
					// 	Inner: []preload{{Base: "User"}},
					// },
				},
				},
				{Base: "PersianYear"},
				/*{Base: "CarModelPriceHistories"},*/
			},
		},
	}
}

// Create
func (s *CarModelYearService) Create(ctx context.Context, req *dto.CreateUpdateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelYearService) Update(ctx context.Context, id int, req *dto.CreateUpdateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelYearService) GetById(ctx context.Context, id int) (*dto.CarModelYearResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelYearResponse], error) {
	return s.base.GetByFilter(ctx, req)

}
