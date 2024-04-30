//go:build migration

package main

import (
	"github.com/behrouz-rfa/kentech/internal/adapter/config"
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo"
	"github.com/behrouz-rfa/kentech/internal/core/services"
	"github.com/behrouz-rfa/kentech/migrations"
	"github.com/behrouz-rfa/kentech/pkg/logger"
	"os"
)

func main() {
	// Init logger
	logger.Init()

	// Load the config from environment variables
	config.Load()
	env := config.Get()

	repo, err := mongo.NewRepository(env.DbConnectionString(), env.Database.Name, config.DbTimeout)

	if err != nil {
		panic(err)
	}

	userService := services.NewUserService(&services.UserServiceParams{
		UserRepo: repo,
	})

	err = migrations.Migrate(repo, migrations.Services{
		User: userService,
	})

	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
