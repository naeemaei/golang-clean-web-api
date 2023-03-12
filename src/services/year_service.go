package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *PersianYearService) Create(ctx context.Context, req *dto.CreatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *PersianYearService) Update(ctx context.Context, id int, req *dto.UpdatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PersianYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PersianYearService) GetById(ctx context.Context, id int) (*dto.PersianYearResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PersianYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PersianYearResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
