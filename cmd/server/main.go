package main

import (
	"fmt"
	_ "github.com/behrouz-rfa/kentech/docs"
	"github.com/behrouz-rfa/kentech/internal/adapter/auth"
	"github.com/behrouz-rfa/kentech/internal/adapter/config"
	"github.com/behrouz-rfa/kentech/internal/adapter/handlers/http"
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo"
	"github.com/behrouz-rfa/kentech/internal/core/services"
	"github.com/behrouz-rfa/kentech/pkg/logger"
	"github.com/behrouz-rfa/kentech/pkg/utils/validate"
	"os"
)

// @title					KenTech
// @version					1.0
// @description				This is a simple RESTful
//
// @contact.name				Behrouz R Faris
// @contact.url				github.com/behrouz-rfa/kentech
// @contact.email			behrouz-rfa@gmail.com
//
// @license.name				MIT
//
// @BasePath /api/v1
// @schemes					http https
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	// initialize logger
	logger.Init()
	lg := logger.General.Component("main")

	// initialize validate
	validate.Init()

	config.Load()
	cfg := config.Get()

	dbRepo, err := mongo.NewRepository(cfg.DbConnectionString(), cfg.Database.Name, config.DbTimeout)
	if err != nil {
		lg.WithError(err).Fatal("failed to connect to database")
	}

	auth := auth.NewAuth(cfg.Jwt.Secret)

	userService := services.NewUserService(
		services.WithUserRepository(dbRepo),
		services.WithAuth(auth),
	)

	filmService := services.NewFilmService(
		services.WithFilmRepository(dbRepo),
	)

	userHandler := http.NewUserHandler(userService)
	filmHandler := http.NewFilmHandler(filmService)

	// Init router
	router, err := http.NewRouter(
		cfg,
		auth,
		*userHandler,
		*filmHandler,
	)
	if err != nil {
		lg.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	lg.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		lg.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
