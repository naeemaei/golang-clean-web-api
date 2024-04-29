package repository

import (
	"context"
	"database/sql"
	"reflect"
	"time"

	"github.com/naeemaei/golang-clean-web-api/common"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/constant"
	filter "github.com/naeemaei/golang-clean-web-api/domain/filter"
	database "github.com/naeemaei/golang-clean-web-api/infra/persistence/database"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
	"github.com/naeemaei/golang-clean-web-api/pkg/metrics"
	"github.com/naeemaei/golang-clean-web-api/pkg/service_errors"
	"gorm.io/gorm"
)

const softDeleteExp string = "id = ? and deleted_by is null"

type BaseRepository[TEntity any] struct {
	database *gorm.DB
	logger   logging.Logger
	preloads []database.PreloadEntity
}

func NewBaseRepository[TEntity any](cfg *config.Config, preloads []database.PreloadEntity) *BaseRepository[TEntity] {
	return &BaseRepository[TEntity]{
		database: database.GetDb(),
		logger:   logging.NewLogger(cfg),
		preloads: preloads,
	}
}

func (r BaseRepository[TEntity]) Create(ctx context.Context, entity TEntity) (TEntity, error) {
	tx := r.database.WithContext(ctx).Begin()
	err := tx.
		Create(&entity).
		Error
	if err != nil {
		tx.Rollback()
		r.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		metrics.DbCall.WithLabelValues(reflect.TypeOf(entity).String(), "Create", "Failed").Inc()
		return entity, err
	}
	tx.Commit()

	metrics.DbCall.WithLabelValues(reflect.TypeOf(entity).String(), "Create", "Success").Inc()
	return entity, nil
}

func (r BaseRepository[TEntity]) Update(ctx context.Context, id int, entity map[string]interface{}) (TEntity, error) {
	snakeMap := map[string]interface{}{}
	for k, v := range entity {
		snakeMap[common.ToSnakeCase(k)] = v
	}
	snakeMap["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constant.UserIdKey).(float64)), Valid: true}
	snakeMap["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	model := new(TEntity)
	tx := r.database.WithContext(ctx).Begin()
	if err := tx.Model(model).
		Where(softDeleteExp, id).
		Updates(snakeMap).
		Error; err != nil {
		tx.Rollback()
		r.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Update", "Failed").Inc()
		return *model, err
	}
	tx.Commit()
	metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Update", "Success").Inc()
	return *model, nil
}

func (r BaseRepository[TEntity]) Delete(ctx context.Context, id int) error {
	tx := r.database.WithContext(ctx).Begin()

	model := new(TEntity)

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constant.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}

	if ctx.Value(constant.UserIdKey) == nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	if cnt := tx.
		Model(model).
		Where(softDeleteExp, id).
		Updates(deleteMap).
		RowsAffected; cnt == 0 {
		tx.Rollback()
		r.logger.Error(logging.Postgres, logging.Update, service_errors.RecordNotFound, nil)
		metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Delete", "Failed").Inc()
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}
	tx.Commit()
	metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "Delete", "Success").Inc()
	return nil
}

func (r BaseRepository[TEntity]) GetById(ctx context.Context, id int) (TEntity, error) {
	model := new(TEntity)
	db := database.Preload(r.database, r.preloads)
	err := db.
		Where(softDeleteExp, id).
		First(model).
		Error
	if err != nil {
		metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "GetById", "Failed").Inc()
		return *model, err
	}
	metrics.DbCall.WithLabelValues(reflect.TypeOf(*model).String(), "GetById", "Success").Inc()
	return *model, nil
}

func (r BaseRepository[TEntity]) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error) {
	model := new(TEntity)
	var items *[]TEntity

	db := database.Preload(r.database, r.preloads)
	query := database.GenerateDynamicQuery[TEntity](&req.DynamicFilter)
	sort := database.GenerateDynamicSort[TEntity](&req.DynamicFilter)
	var totalRows int64 = 0

	db.
		Model(model).
		Where(query).
		Count(&totalRows)

	err := db.
		Where(query).
		Offset(req.GetOffset()).
		Limit(req.GetPageSize()).
		Order(sort).
		Find(&items).
		Error

	if err != nil {
		return 0, &[]TEntity{}, err
	}
	return totalRows, items, err

}
