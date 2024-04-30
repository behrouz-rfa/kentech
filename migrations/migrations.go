//go:build migration

package migrations

import (
	"context"
	"github.com/behrouz-rfa/kentech/internal/core/ports"
	"github.com/stackus/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	repo "github.com/behrouz-rfa/kentech/internal/adapter/mongo"
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/utils"
)

type Repo interface {
	DB() *mongo.Database
}

// Services is a collection of services that are used by migrations.
type Services struct {
	User ports.UserService
}

type MigrationOperation func(r Repo, s Services) error

// Migration is a base struct for all migrations
// It is a mongo document that get saved in the database.
type Migration struct {
	repo.DocumentBase `bson:",inline"`
	Name              string             `bson:"name"`
	Ticket            string             `bson:"ticket"`
	Action            MigrationOperation `bson:"-"`
	RollBack          MigrationOperation `bson:"-"`
}

// Migrations is a collection of all migrations.
var Migrations = []Migration{
	InitialMigration,
	UniqueUsername,
}

// Migrate runs all migrations that are not performed yet.
func Migrate(r Repo, s Services) error {
	performedMigrations := make(map[string]string)

	db := r.DB()

	ctx := context.Background()

	// Get all migrations that are already performed
	cursor, err := db.Collection("migrations").Find(ctx, bson.M{})
	if err != nil {
		return errors.Wrap(err, "error finding migrations")
	}

	defer cursor.Close(ctx)

	// Iterate over all migrations and add them to the performedMigrations map
	for cursor.Next(ctx) {
		m := new(Migration)

		err = cursor.Decode(m)
		if err != nil {
			return err
		}

		performedMigrations[m.Name] = m.Ticket
	}

	// Iterate over all migrations and run the ones that are not performed yet
	for _, migration := range Migrations {
		// Check if the migration is already performed
		if _, ok := performedMigrations[migration.Name]; ok {
			continue
		}

		log.Println("Migrating: " + migration.Name)
		// Run the migration
		err = migration.Action(r, s)
		if err != nil {
			log.Println("Rolling back migration: " + migration.Name)
			rErr := migration.RollBack(r, s)

			if rErr != nil {
				log.Println("Error rolling back migration:", rErr)
			}

			return err
		}

		// Convert the migration to a map
		migrationData := utils.ToMap(migration)

		// Save the migration in the database
		_, err = db.Collection("migrations").InsertOne(ctx, migrationData)
		if err != nil {
			return err
		}
	}

	return nil
}
