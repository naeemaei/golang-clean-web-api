package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type FileUsecase struct {
	base *BaseUsecase[model.File, dto.CreateFile, dto.UpdateFile, dto.File]
}

func NewFileUsecase(cfg *config.Config, repository repository.FileRepository) *FileUsecase {
	return &FileUsecase{
		base: NewBaseUsecase[model.File, dto.CreateFile, dto.UpdateFile, dto.File](cfg, repository),
	}
}

// Create
func (u *FileUsecase) Create(ctx context.Context, req dto.CreateFile) (dto.File, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *FileUsecase) Update(ctx context.Context, id int, req dto.UpdateFile) (dto.File, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *FileUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *FileUsecase) GetById(ctx context.Context, id int) (dto.File, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *FileUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.File], error) {
	return s.base.GetByFilter(ctx, req)
}
