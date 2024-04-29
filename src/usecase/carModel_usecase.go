package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CarModelUsecase struct {
	base *BaseUsecase[model.CarModel, dto.CreateCarModel, dto.UpdateCarModel, dto.CarModel]
}

func NewCarModelUsecase(cfg *config.Config, repository repository.CarModelRepository) *CarModelUsecase {
	return &CarModelUsecase{
		base: NewBaseUsecase[model.CarModel, dto.CreateCarModel, dto.UpdateCarModel, dto.CarModel](cfg, repository),
	}
}

// Create
func (u *CarModelUsecase) Create(ctx context.Context, req dto.CreateCarModel) (dto.CarModel, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CarModelUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModel) (dto.CarModel, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelUsecase) GetById(ctx context.Context, id int) (dto.CarModel, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModel], error) {
	return s.base.GetByFilter(ctx, req)
}
