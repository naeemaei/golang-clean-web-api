package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type PropertyUsecase struct {
	base *BaseUsecase[model.Property, dto.CreateProperty, dto.UpdateProperty, dto.Property]
}

func NewPropertyUsecase(cfg *config.Config, repository repository.PropertyRepository) *PropertyUsecase {
	return &PropertyUsecase{
		base: NewBaseUsecase[model.Property, dto.CreateProperty, dto.UpdateProperty, dto.Property](cfg, repository),
	}
}

// Create
func (u *PropertyUsecase) Create(ctx context.Context, req dto.CreateProperty) (dto.Property, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *PropertyUsecase) Update(ctx context.Context, id int, req dto.UpdateProperty) (dto.Property, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PropertyUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PropertyUsecase) GetById(ctx context.Context, id int) (dto.Property, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PropertyUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Property], error) {
	return s.base.GetByFilter(ctx, req)
}
