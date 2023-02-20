package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type CarModelImageService struct {
	base *BaseService[models.CarModelImage, dto.CreateUpdateCarModelImageRequest, dto.CreateUpdateCarModelImageRequest, dto.CarModelImageResponse]
}

func NewCarModelImageService(cfg *config.Config) *CarModelImageService {
	return &CarModelImageService{
		base: &BaseService[models.CarModelImage, dto.CreateUpdateCarModelImageRequest, dto.CreateUpdateCarModelImageRequest, dto.CarModelImageResponse]{
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
					}, {Base: "CarModelYears",
						Inner: []preload{{Base: "PersianYear"} /*{Base: "CarModelPriceHistories"}*/},
					}, {Base: "CarModelColors",
						Inner: []preload{{Base: "Color"}},
					},
					// {Base: "CarModelComments",
					// 	Inner: []preload{{Base: "User"}},
					// },
				},
				},
				{Base: "Image"},
			},
		},
	}
}

// Create
func (s *CarModelImageService) Create(ctx context.Context, req *dto.CreateUpdateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelImageService) Update(ctx context.Context, id int, req *dto.CreateUpdateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelImageService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelImageService) GetById(ctx context.Context, id int) (*dto.CarModelImageResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelImageService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelImageResponse], error) {
	return s.base.GetByFilter(ctx, req)

}