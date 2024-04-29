package usecase

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/common"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	"github.com/naeemaei/golang-clean-web-api/domain/repository"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
)

type BaseUsecase[TEntity any, TCreate any, TUpdate any, TResponse any] struct {
	logger     logging.Logger
	repository repository.BaseRepository[TEntity]
}

func NewBaseUsecase[TEntity any, TCreate any, TUpdate any, TResponse any](cfg *config.Config, repository repository.BaseRepository[TEntity]) *BaseUsecase[TEntity, TCreate, TUpdate, TResponse] {
	logger := logging.NewLogger(cfg)
	return &BaseUsecase[TEntity, TCreate, TUpdate, TResponse]{
		repository: repository,
		logger:     logger,
	}
}

func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) Create(ctx context.Context, req TCreate) (TResponse, error) {
	var response TResponse
	entity, _ := common.TypeConverter[TEntity](req)

	entity, err := u.repository.Create(ctx, entity)
	if err != nil {
		return response, err
	}

	response, _ = common.TypeConverter[TResponse](entity)
	return response, nil
}

func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) Update(ctx context.Context, id int, req TUpdate) (TResponse, error) {
	var response TResponse
	updateMap, _ := common.TypeConverter[map[string]interface{}](req)

	entity, err := u.repository.Update(ctx, id, updateMap)
	if err != nil {
		return response, err
	}
	response, _ = common.TypeConverter[TResponse](entity)

	return response, nil
}

func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) Delete(ctx context.Context, id int) error {

	return u.repository.Delete(ctx, id)
}

func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) GetById(ctx context.Context, id int) (TResponse, error) {
	var response TResponse
	entity, err := u.repository.GetById(ctx, id)
	if err != nil {
		return response, err
	}
	return common.TypeConverter[TResponse](entity)
}

func (u *BaseUsecase[TEntity, TCreate, TUpdate, TResponse]) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[TResponse], error) {
	var response *filter.PagedList[TResponse]
	count, entities, err := u.repository.GetByFilter(ctx, req)
	if err != nil {
		return response, err
	}

	return filter.Paginate[TEntity, TResponse](count, entities, req.PageNumber, int64(req.PageSize))
}
