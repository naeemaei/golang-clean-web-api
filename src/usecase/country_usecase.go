package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CountryUsecase struct {
	base *BaseUsecase[model.Country, dto.Name, dto.Name, dto.Country]
}

func NewCountryUsecase(cfg *config.Config, repository repository.CountryRepository) *CountryUsecase {
	return &CountryUsecase{
		base: NewBaseUsecase[model.Country, dto.Name, dto.Name, dto.Country](cfg, repository),
	}
}

// Create
func (u *CountryUsecase) Create(ctx context.Context, req dto.Name) (dto.Country, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CountryUsecase) Update(ctx context.Context, id int, req dto.Name) (dto.Country, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CountryUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CountryUsecase) GetById(ctx context.Context, id int) (dto.Country, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CountryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Country], error) {
	return u.base.GetByFilter(ctx, req)
}
