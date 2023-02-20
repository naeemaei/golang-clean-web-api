package services

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/data/models"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.CreateCarModelCommentRequest, dto.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: &BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.CreateCarModelCommentRequest, dto.CarModelCommentResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{Base: "CarModel"}, {Base: "User"},
			},
		}}
}

// Create
func (s *CarModelCommentService) Create(ctx context.Context, req *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelCommentService) Update(ctx context.Context, id int, req *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelCommentService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelCommentService) GetById(ctx context.Context, id int) (*dto.CarModelCommentResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelCommentService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelCommentResponse], error) {
	return s.base.GetByFilter(ctx, req)

}
