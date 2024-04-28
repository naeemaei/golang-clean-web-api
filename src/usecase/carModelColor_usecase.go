package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CarModelColorUsecase struct {
	base *BaseUsecase[model.CarModelColor, dto.CreateCarModelColor, dto.UpdateCarModelColor, dto.CarModelColor]
}

func NewCarModelColorUsecase(cfg *config.Config, repository repository.CarModelColorRepository) *CarModelColorUsecase {
	return &CarModelColorUsecase{
		base: NewBaseUsecase[model.CarModelColor, dto.CreateCarModelColor, dto.UpdateCarModelColor, dto.CarModelColor](cfg, repository),
	}
}

// Create
func (u *CarModelColorUsecase) Create(ctx context.Context, req dto.CreateCarModelColor) (dto.CarModelColor, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CarModelColorUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelColor) (dto.CarModelColor, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelColorUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelColorUsecase) GetById(ctx context.Context, id int) (dto.CarModelColor, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelColorUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelColor], error) {
	return s.base.GetByFilter(ctx, req)
}
