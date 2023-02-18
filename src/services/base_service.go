package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/naeemaei/golang-clean-web-api/api/dto"
	"github.com/naeemaei/golang-clean-web-api/common"
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/constants"
	"github.com/naeemaei/golang-clean-web-api/data/db"
	"github.com/naeemaei/golang-clean-web-api/pkg/logging"
	"github.com/naeemaei/golang-clean-web-api/pkg/service_errors"
	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger   logging.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
		Logger:   logging.NewLogger(cfg),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {

	model, _ := common.TypeConverter[T](req)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.
		Create(model).
		Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return common.TypeConverter[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {

	updateMap, _ := common.TypeConverter[map[string]interface{}](req)
	(*updateMap)["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	(*updateMap)["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	if err := tx.Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(*updateMap).
		Error; err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return s.GetById(ctx, id)

}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin()

	model := new(T)

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	deleteMap["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	deleteMap["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	if ctx.Value(constants.UserIdKey) == nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	if cnt := tx.
		Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(deleteMap).
		RowsAffected; cnt == 0 {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, service_errors.RecordNotFound, nil)
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	err := s.Database.
		Where("id = ? and deleted_by is null", id).
		First(model).
		Error
	if err != nil {
		return nil, err
	}
	return common.TypeConverter[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[Tr], error) {

	pl, err := Paginate[T, Tr](req, s.Database)

	if err != nil {
		return nil, err
	}

	return pl, nil
}

func NewPagedList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PagedList[T] {
	pl := &dto.PagedList[T]{PageNumber: pageNumber, TotalRows: count, Items: items}
	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasPreviousPage = pl.PageNumber > 1

	return pl
}

func Paginate[T any, Tr any](pagination *dto.PaginationInputWithFilter, db *gorm.DB) (*dto.PagedList[Tr], error) {
	model := new(T)
	var items *[]T
	var mitems *[]Tr 
	var query = getQuery[T](&pagination.DynamicFilter)
	var order = getSort[T](&pagination.DynamicFilter)
	var totalRows int64 = 0
	db.
		Model(model).
		Where(query).
		Count(&totalRows)

	err := db.
		Where(query).
		Offset(pagination.GetOffset()).
		Limit(pagination.GetPageSize()).
		Order(order).
		Find(&items).
		Error
	if err != nil {
		return nil, err
	}
	mitems, err = common.TypeConverter[[]Tr](items)
	//err = mapper.MapperSlice(items, mitems)
	if err != nil {
		return nil, err
	}
	return NewPagedList(mitems, totalRows, pagination.PageNumber, int64(pagination.PageSize)), err

}

func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	for name, filter := range filter.Filter {
		fld, ok := typeT.FieldByName(name)
		if ok {
			switch filter.Type {
			case "contains":
				query = append(query, fmt.Sprintf("%s Ilike '%%%s%%'", fld.Name, filter.From))
			case "notContains":
				query = append(query, fmt.Sprintf("%s not Ilike '%%%s%%'", fld.Name, filter.From))
			case "startsWith":
				query = append(query, fmt.Sprintf("%s Ilike '%s%%'", fld.Name, filter.From))
			case "endsWith":
				query = append(query, fmt.Sprintf("%s Ilike '%%%s'", fld.Name, filter.From))
			case "equals":
				query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
			case "notEqual":
				query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
			case "lessThan":
				query = append(query, fmt.Sprintf("%s < %s", fld.Name, filter.From))
			case "lessThanOrEqual":
				query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.From))
			case "greaterThan":
				query = append(query, fmt.Sprintf("%s > %s", fld.Name, filter.From))
			case "greaterThanOrEqual":
				query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
			case "inRange":
				if fld.Type.Kind() == reflect.String {
					query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
					query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
				} else {
					query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
					query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
				}
			}
		}
	}
	return strings.Join(query, " AND ")
}

func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)

	for _, tp := range *filter.Sort {
		fld, ok := typeT.FieldByName(tp.ColId)
		if ok {
			sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
		}
	}
	return strings.Join(sort, ", ")
}
