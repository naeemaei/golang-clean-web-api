package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	model "github.com/naeemaei/golang-clean-web-api/domain/model"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/usecase/dto"
)

type CompanyUsecase struct {
	base *BaseUsecase[model.Company, dto.CreateCompany, dto.UpdateCompany, dto.Company]
}

func NewCompanyUsecase(cfg *config.Config, repository repository.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{
		base: NewBaseUsecase[model.Company, dto.CreateCompany, dto.UpdateCompany, dto.Company](cfg, repository),
	}
}

// Create
func (u *CompanyUsecase) Create(ctx context.Context, req dto.CreateCompany) (dto.Company, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CompanyUsecase) Update(ctx context.Context, id int, req dto.UpdateCompany) (dto.Company, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CompanyUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CompanyUsecase) GetById(ctx context.Context, id int) (dto.Company, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CompanyUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Company], error) {
	return s.base.GetByFilter(ctx, req)
}
