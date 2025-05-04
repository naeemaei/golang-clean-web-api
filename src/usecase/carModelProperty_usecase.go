package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CarModelPropertyUsecase struct {
	base *BaseUsecase[model.CarModelProperty, dto.CreateCarModelProperty, dto.UpdateCarModelProperty, dto.CarModelProperty]
}

func NewCarModelPropertyUsecase(cfg *config.Config, repository repository.CarModelPropertyRepository) *CarModelPropertyUsecase {
	return &CarModelPropertyUsecase{
		base: NewBaseUsecase[model.CarModelProperty, dto.CreateCarModelProperty, dto.UpdateCarModelProperty, dto.CarModelProperty](cfg, repository),
	}
}

// Create
func (u *CarModelPropertyUsecase) Create(ctx context.Context, req dto.CreateCarModelProperty) (dto.CarModelProperty, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CarModelPropertyUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelProperty) (dto.CarModelProperty, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CarModelPropertyUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CarModelPropertyUsecase) GetById(ctx context.Context, id int) (dto.CarModelProperty, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CarModelPropertyUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelProperty], error) {
	return u.base.GetByFilter(ctx, req)
}
