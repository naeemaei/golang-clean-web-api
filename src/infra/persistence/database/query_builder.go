package database

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/naeemaei/golang-clean-web-api/common"
	"gorm.io/gorm"

	filter "github.com/naeemaei/golang-clean-web-api/domain/filter"
)

type PreloadEntity struct {
	Entity string
}

// GenerateDynamicQuery
func GenerateDynamicQuery[T any](filter *filter.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			if fld, ok := typeT.FieldByName(name); ok {
				query = append(query, GenerateDynamicFilter(fld, filter))
			}
		}
	}
	return strings.Join(query, " AND ")
}

func GenerateDynamicFilter(fld reflect.StructField, filter filter.Filter) string {
	conditionQuery := ""
	fld.Name = common.ToSnakeCase(fld.Name)
	switch filter.Type {
	case "contains":
		conditionQuery = fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From)
	case "notContains":
		conditionQuery = fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From)
	case "startsWith":
		conditionQuery = fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From)
	case "endsWith":
		conditionQuery = fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From)
	case "equals":
		conditionQuery = fmt.Sprintf("%s = '%s'", fld.Name, filter.From)
	case "notEqual":
		conditionQuery = fmt.Sprintf("%s != '%s'", fld.Name, filter.From)
	case "lessThan":
		conditionQuery = fmt.Sprintf("%s < %s", fld.Name, filter.From)
	case "lessThanOrEqual":
		conditionQuery = fmt.Sprintf("%s <= %s", fld.Name, filter.From)
	case "greaterThan":
		conditionQuery = fmt.Sprintf("%s > %s", fld.Name, filter.From)
	case "greaterThanOrEqual":
		conditionQuery = fmt.Sprintf("%s >= %s", fld.Name, filter.From)
	case "inRange":
		if fld.Type.Kind() == reflect.String {
			conditionQuery = fmt.Sprintf("%s >= '%s%%' AND ", fld.Name, filter.From)
			conditionQuery += fmt.Sprintf("%s <= '%%%s'", fld.Name, filter.To)
		} else {
			conditionQuery = fmt.Sprintf("%s >= %s AND ", fld.Name, filter.From)
			conditionQuery += fmt.Sprintf("%s <= %s", fld.Name, filter.To)
		}
	}
	return conditionQuery
}

// generateDynamicSort
func GenerateDynamicSort[T any](filter *filter.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range *filter.Sort {
			fld, ok := typeT.FieldByName(tp.ColId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}

// Preload
func Preload(db *gorm.DB, preloads []PreloadEntity) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.Entity)
	}
	return db
}
