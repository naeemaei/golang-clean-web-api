package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CarModelYearUsecase struct {
	base *BaseUsecase[model.CarModelYear, dto.CreateCarModelYear, dto.UpdateCarModelYear, dto.CarModelYear]
}

func NewCarModelYearUsecase(cfg *config.Config, repository repository.CarModelYearRepository) *CarModelYearUsecase {
	return &CarModelYearUsecase{
		base: NewBaseUsecase[model.CarModelYear, dto.CreateCarModelYear, dto.UpdateCarModelYear, dto.CarModelYear](cfg, repository),
	}
}

// Create
func (u *CarModelYearUsecase) Create(ctx context.Context, req dto.CreateCarModelYear) (dto.CarModelYear, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CarModelYearUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelYear) (dto.CarModelYear, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CarModelYearUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CarModelYearUsecase) GetById(ctx context.Context, id int) (dto.CarModelYear, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CarModelYearUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelYear], error) {
	return u.base.GetByFilter(ctx, req)
}
