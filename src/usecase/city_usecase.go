package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CityUsecase struct {
	base *BaseUsecase[model.City, dto.CreateCity, dto.UpdateCity, dto.City]
}

func NewCityUsecase(cfg *config.Config, repository repository.CityRepository) *CityUsecase {
	return &CityUsecase{
		base: NewBaseUsecase[model.City, dto.CreateCity, dto.UpdateCity, dto.City](cfg, repository),
	}
}

// Create
func (u *CityUsecase) Create(ctx context.Context, req dto.CreateCity) (dto.City, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CityUsecase) Update(ctx context.Context, id int, req dto.UpdateCity) (dto.City, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CityUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CityUsecase) GetById(ctx context.Context, id int) (dto.City, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CityUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.City], error) {
	return u.base.GetByFilter(ctx, req)
}
