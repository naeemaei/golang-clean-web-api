package filter

import (
	"math"

	"github.com/naeemaei/golang-clean-web-api/common"
)

func NewPagedList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *PagedList[T] {
	pl := &PagedList[T]{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		TotalRows:  count,
		Items:      items,
	}
	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasPreviousPage = pl.PageNumber > 1

	return pl
}

// Paginate
func Paginate[TInput any, TOutput any](totalRows int64, items *[]TInput, pageNumber int, pageSize int64) (*PagedList[TOutput], error) {
	var rItems []TOutput

	rItems, err := common.TypeConverter[[]TOutput](items)
	if err != nil {
		return nil, err
	}
	return NewPagedList(&rItems, totalRows, pageNumber, pageSize), err

}

type PagedList[T any] struct {
	PageNumber      int   `json:"pageNumber"`
	PageSize        int64 `json:"pageSize"`
	TotalRows       int64 `json:"totalRows"`
	TotalPages      int   `json:"totalPages"`
	HasPreviousPage bool  `json:"hasPreviousPage"`
	HasNextPage     bool  `json:"hasNextPage"`
	Items           *[]T  `json:"items"`
}

type PaginationInput struct {
	PageSize   int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
}

type PaginationInputWithFilter struct {
	PaginationInput
	DynamicFilter
}

func (p *PaginationInputWithFilter) GetOffset() int {
	// 2 , 10 => 11-20
	return (p.GetPageNumber() - 1) * p.GetPageSize()
}

func (p *PaginationInputWithFilter) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *PaginationInputWithFilter) GetPageNumber() int {
	if p.PageNumber == 0 {
		p.PageNumber = 1
	}
	return p.PageNumber
}
