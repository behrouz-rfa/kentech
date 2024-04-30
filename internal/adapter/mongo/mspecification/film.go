package mspecification

import (
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/utils"
	coreSpec "github.com/behrouz-rfa/kentech/internal/core/specefication"
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/sort"
	bson "go.mongodb.org/mongo-driver/bson"
)

type FilmSpecification struct {
	BaseSpecification
}

//nolint:dupl // False positive, the resources are different and can't be generalized.
func (u *FilmSpecification) SortBy(sort *sort.FilmSort) coreSpec.FilmSpecification {
	if sort == nil {
		return u
	}

	if sort.Name != nil {
		u.Sort("name", *sort.Name)
	}

	return u
}

// NewFilmSpecification creates a new instance of FilmSpecification.
func NewFilmSpecification() *FilmSpecification {
	spec := FilmSpecification{}
	spec.parent = &spec

	return &spec
}

func (u *FilmSpecification) NeedsPreload(field string) (bool, []bson.M) {
	// If the field is indexed, it means it has already been preloaded.
	if shouldPreload, _ := u.BaseSpecification.NeedsPreload(field); !shouldPreload {
		return false, nil
	}

	return false, nil
}

func (u *FilmSpecification) By(filter filters.FilmBy) coreSpec.FilmSpecification {
	f := utils.ToMap(filter, utils.MethodFilter)
	u.filters = append(u.filters, bson.M(f))

	return u
}

func (u *FilmSpecification) Filter(filter *filters.FilmFilter) coreSpec.FilmSpecification {
	if filter == nil {
		return u
	}

	bsonFilter := bson.M{}

	if filter.Title != nil {
		bsonFilter["title"] = stringFilterBson(*filter.Title)
	}

	if len(bsonFilter) > 0 {
		u.filters = append(u.filters, bsonFilter)
	}

	return u
}
