//go:build migration

package migrations

import (
	"context"
	"github.com/stackus/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var UniqueUsername = Migration{
	Name:   "UniqueUsername-3003231424",
	Ticket: "Unique",
	Action: func(r Repo, s Services) error {
		userColl := r.DB().Collection("users")
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{"username", 1}}, // Field to create an index on (1 for ascending, -1 for descending)
			Options: options.Index().SetUnique(true),
		}
		_, err := userColl.Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Fatal(err)
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
