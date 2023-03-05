package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *GearboxService) Create(ctx context.Context, req *dto.CreateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *GearboxService) Update(ctx context.Context, id int, req *dto.UpdateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *GearboxService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *GearboxService) GetById(ctx context.Context, id int) (*dto.GearboxResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *GearboxService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
