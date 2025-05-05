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
func (u *CarModelUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModel) (dto.CarModel, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CarModelUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CarModelUsecase) GetById(ctx context.Context, id int) (dto.CarModel, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CarModelUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModel], error) {
	return u.base.GetByFilter(ctx, req)
}
