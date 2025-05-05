package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CarModelImageUsecase struct {
	base *BaseUsecase[model.CarModelImage, dto.CreateCarModelImage, dto.UpdateCarModelImage, dto.CarModelImage]
}

func NewCarModelImageUsecase(cfg *config.Config, repository repository.CarModelImageRepository) *CarModelImageUsecase {
	return &CarModelImageUsecase{
		base: NewBaseUsecase[model.CarModelImage, dto.CreateCarModelImage, dto.UpdateCarModelImage, dto.CarModelImage](cfg, repository),
	}
}

// Create
func (u *CarModelImageUsecase) Create(ctx context.Context, req dto.CreateCarModelImage) (dto.CarModelImage, error) {
	return u.base.Create(ctx, req)
}

// Update
func (u *CarModelImageUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelImage) (dto.CarModelImage, error) {
	return u.base.Update(ctx, id, req)
}

// Delete
func (u *CarModelImageUsecase) Delete(ctx context.Context, id int) error {
	return u.base.Delete(ctx, id)
}

// Get By Id
func (u *CarModelImageUsecase) GetById(ctx context.Context, id int) (dto.CarModelImage, error) {
	return u.base.GetById(ctx, id)
}

// Get By Filter
func (u *CarModelImageUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelImage], error) {
	return u.base.GetByFilter(ctx, req)
}
