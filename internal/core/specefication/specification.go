package specification

import (
	"context"

	"github.com/behrouz-rfa/kentech/internal/pagination"
	"github.com/behrouz-rfa/kentech/internal/sort"
)

type Specification interface {
	Query() interface{}
	// QueryNested Help to use filters for child documents on database
	QueryNested() interface{}
}

type BaseSpecification struct{}

type Set interface {
	Specification
	WithContext(ctx context.Context) Set
	FilterByID(id string) Set
	FilterByIDs(ids []string) Set
	FilterEntry(attributes map[string]interface{}) Set
	Prefetch(preload ...string) Set
	Sort(sortBy string, sortOrder sort.Order) Set
	Limit(limit uint64) Set
	LessThan(field string, value interface{}) Set
	GreaterThan(field string, value interface{}) Set
	Paginate(paginate *pagination.Pagination) Set
}
