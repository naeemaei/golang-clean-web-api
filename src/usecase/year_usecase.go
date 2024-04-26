package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type PersianYearUsecase struct {
	base *BaseUsecase[model.PersianYear, dto.CreatePersianYear, dto.UpdatePersianYear, dto.PersianYear]
}

func NewPersianYearUsecase(cfg *config.Config, repository repository.PersianYearRepository) *PersianYearUsecase {
	return &PersianYearUsecase{
		base: NewBaseUsecase[model.PersianYear, dto.CreatePersianYear, dto.UpdatePersianYear, dto.PersianYear](cfg, repository),
	}
}

// Create
func (u *PersianYearUsecase) Create(ctx context.Context, req dto.CreatePersianYear) (dto.PersianYear, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *PersianYearUsecase) Update(ctx context.Context, id int, req dto.UpdatePersianYear) (dto.PersianYear, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PersianYearUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PersianYearUsecase) GetById(ctx context.Context, id int) (dto.PersianYear, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PersianYearUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.PersianYear], error) {
	return s.base.GetByFilter(ctx, req)
}
