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

	if sort.Title != nil {
		u.Sort("title", *sort.Title)
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
	switch field {
	case "user":
		return true, []bson.M{
			{
				"$lookup": bson.M{
					"from":         "users",
					"localField":   "creatorId",
					"foreignField": "_id",
					"as":           "user",
				},
			},
			{
				"$addFields": bson.M{
					"user": bson.M{
						"$arrayElemAt": []interface{}{"$user", 0},
					},
				},
			},
		}
	default:
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
	if filter.Genre != nil {
		bsonFilter["genre"] = stringFilterBson(*filter.Genre)
	}

	if filter.ReleaseDate != nil {
		bsonFilter["releaseDate"] = timeRangeBson(*filter.ReleaseDate)
	}

	if len(bsonFilter) > 0 {
		u.filters = append(u.filters, bsonFilter)
	}

	return u
}
