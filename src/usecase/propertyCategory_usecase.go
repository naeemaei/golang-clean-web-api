package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type PropertyCategoryUsecase struct {
	base *BaseUsecase[model.PropertyCategory, dto.CreatePropertyCategory, dto.UpdatePropertyCategory, dto.PropertyCategory]
}

func NewPropertyCategoryUsecase(cfg *config.Config, repository repository.PropertyCategoryRepository) *PropertyCategoryUsecase {
	return &PropertyCategoryUsecase{
		base: NewBaseUsecase[model.PropertyCategory, dto.CreatePropertyCategory, dto.UpdatePropertyCategory, dto.PropertyCategory](cfg, repository),
	}
}

// Create
func (u *PropertyCategoryUsecase) Create(ctx context.Context, req dto.CreatePropertyCategory) (dto.PropertyCategory, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *PropertyCategoryUsecase) Update(ctx context.Context, id int, req dto.UpdatePropertyCategory) (dto.PropertyCategory, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PropertyCategoryUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PropertyCategoryUsecase) GetById(ctx context.Context, id int) (dto.PropertyCategory, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PropertyCategoryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.PropertyCategory], error) {
	return s.base.GetByFilter(ctx, req)
}
