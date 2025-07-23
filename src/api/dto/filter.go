package dto

type Sort struct {
	ColId string `json:"col_id"`
	Sort  string `json:"sort"`
}

type Filter struct {
	Type string `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`

	FilterType string `json:"filterType"`
}

type DynamicFilter struct {
	Sort   *[]Sort           `json:"sort"`
	Filter map[string]Filter `json:"filter"`
}
type PagedList[T any] struct {
	PageNumber      int  `json:"pageNumber"`
	TotalRows       int64  `json:"totalRows"`
	TotalPage       int  `json:"totalPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	HasNextPage     bool `json:"hasNextPage"`
	Item            *[]T `json:"item"`
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
	if p.PageSize == 0 {
		p.PageSize = 1
	}
	return p.PageNumber
}
