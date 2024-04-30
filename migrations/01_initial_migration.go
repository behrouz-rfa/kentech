//go:build migration

package migrations

import (
	"context"
	"github.com/stackus/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var InitialMigration = Migration{
	Name:   "initial-migration-3003231424",
	Ticket: "OUK-287",
	Action: func(r Repo, s Services) error {
		col := r.DB().Collection("migrations")

		_, err := col.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		})

		if err != nil {
			return errors.Wrap(err, "error creating index on InitialMigration")
		}

		return nil
	},
	RollBack: func(r Repo, s Services) error {
		err := r.DB().Collection("migrations").Drop(context.Background())

		if err != nil {
			return errors.Wrap(err, "error dropping collection on InitialMigration rollback")
		}

		return nil
	},
}
