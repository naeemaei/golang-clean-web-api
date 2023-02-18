package dto

type Sort struct {
	ColId string `json:"colId"`
	Sort  string `json:"sort"`
}

type Filter struct {
	// contains notContains equals notEqual startsWith lessThan lessThanOrEqual greaterThan greaterThanOrEqual inRange endsWith
	Type string `json:"type"`
	From string `json:"filter"`
	To   string `json:"filterTo"`

	// text number
	FilterType string `json:"filterType"`
}
type DynamicFilter struct {
	Sort   *[]Sort           `json:"sort"`
	Filter map[string]Filter `json:"filter"`
}

type PagedList[T any] struct {
	PageNumber      int   `json:"pageNumber"`
	TotalRows       int64 `json:"total"`
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
