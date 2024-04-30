package specification

import (
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"github.com/behrouz-rfa/kentech/internal/sort"
)

type UserSpecification interface {
	Set
	By(filter filters.UserBy) UserSpecification
	Filter(filter *filters.UserFilter) UserSpecification
	SortBy(sort *sort.UserSort) UserSpecification
	Paginate(paginate *pagination.Pagination) Set
}
