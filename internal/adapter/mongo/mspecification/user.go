package mspecification

import (
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/utils"
	coreSpec "github.com/behrouz-rfa/kentech/internal/core/specefication"
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/sort"
	bson "go.mongodb.org/mongo-driver/bson"
)

type UserSpecification struct {
	BaseSpecification
}

//nolint:dupl // False positive, the resources are different and can't be generalized.
func (u *UserSpecification) SortBy(sort *sort.UserSort) coreSpec.UserSpecification {
	if sort == nil {
		return u
	}

	if sort.Username != nil {
		u.Sort("username", *sort.Username)
	}

	return u
}

// NewUserSpecification creates a new instance of UserSpecification.
func NewUserSpecification() *UserSpecification {
	spec := UserSpecification{}
	spec.parent = &spec

	return &spec
}

func (u *UserSpecification) NeedsPreload(field string) (bool, []bson.M) {
	// If the field is indexed, it means it has already been preloaded.
	if shouldPreload, _ := u.BaseSpecification.NeedsPreload(field); !shouldPreload {
		return false, nil
	}

	switch field {
	case "films":
		return true, []bson.M{
			{
				"$lookup": bson.M{
					"from":         "films",
					"localField":   "_id",
					"foreignField": "creatorId",
					"as":           "films",
				},
			},
		}
	default:
		return false, nil
	}
}

func (u *UserSpecification) By(filter filters.UserBy) coreSpec.UserSpecification {
	f := utils.ToMap(filter, utils.MethodFilter)
	u.filters = append(u.filters, bson.M(f))

	return u
}

func (u *UserSpecification) Filter(filter *filters.UserFilter) coreSpec.UserSpecification {
	if filter == nil {
		return u
	}

	bsonFilter := bson.M{}

	if filter.Username != nil {
		bsonFilter["username"] = stringFilterBson(*filter.Username)
	}

	if len(bsonFilter) > 0 {
		u.filters = append(u.filters, bsonFilter)
	}

	return u
}
