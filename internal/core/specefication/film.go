package specification

import (
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/sort"
)

type FilmSpecification interface {
	Set
	By(filter filters.FilmBy) FilmSpecification
	Filter(filter *filters.FilmFilter) FilmSpecification
	SortBy(sort *sort.FilmSort) FilmSpecification
}
