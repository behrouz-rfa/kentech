package mongo

import (
	"time"

	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/utils"
	gUtils "github.com/behrouz-rfa/kentech/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// diskAggregationOption tells mongo to store aggregation results on disk
// instead of in memory.
// use this option when the aggregation result is large.
// see: https://godoc.org/go.mongodb.org/mongo-driver/mongo/options#AggregateOptions
var diskAggregationOption = &options.AggregateOptions{AllowDiskUse: gUtils.ToValue(true)}

type Document interface {
	GetID() string
	SetID(id string)
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
	CollectionName() string
}

type DocumentBase struct {
	ID        string    `json:"_id" bson:"_id"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (d *DocumentBase) GetID() string {
	return d.ID
}

func (d *DocumentBase) SetID(id string) {
	d.ID = id
}

func (d *DocumentBase) GenerateID() {
	d.ID = utils.GenerateUUID()
}

func (d *DocumentBase) SetCreatedAt() {
	d.CreatedAt = time.Now()
}

func (d *DocumentBase) SetUpdatedAt() {
	d.UpdatedAt = time.Now()
}

func (d *DocumentBase) CollectionName() string {
	return collectionName(d)
}

// func (d *DocumentBase) FindOneBy(ctx context.Context, spec specification.Set,
//	collection *mongo.Collection) (*T, error) {
//	var results []Q
//
//	spec.WithContext(ctx).Limit(1)
//	cursor, err := collection.Aggregate(ctx, spec.Query())
//
//	if err != nil {
//		return nil, err
//	}
//
//	if err = cursor.All(ctx, &results); err != nil {
//		return nil, err
//	}
//
//	if len(results) < 1 {
//		return nil, nil
//	}
//
//	return results[0].ToModel(), nil
//}
//
//// ToModel converts a document to a model.
// func (d *DocumentBase) ToModel() *T {
//	panic("implement me")
//}
//
// func (d *DocumentBase) fields() []string {
//	fields := structs.New(d).Fields()
//	var result []string
//	for _, field := range fields {
//		result = append(result, field.Name())
//	}
//
//	return result
//}
