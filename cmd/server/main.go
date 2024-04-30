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
	// Initialize logger
	logger.Init()
	lg := logger.General.Component("main")

	// Initialize validation
	validate.Init()

	// Load configuration
	config.Load()
	cfg := config.Get()

	// Initialize database repository
	dbRepo, err := mongo.NewRepository(cfg.DbConnectionString(), cfg.Database.Name, config.DbTimeout)
	if err != nil {
		lg.WithError(err).Fatal("failed to connect to database")
	}
	defer dbRepo.Close()

	// Initialize authentication service
	authService := auth.NewAuth(cfg.Jwt.Secret)

	// Initialize services
	userService := services.NewUserService(
		services.WithUserRepository(dbRepo),
		services.WithAuth(authService),
	)
	filmService := services.NewFilmService(
		services.WithFilmRepository(dbRepo),
	)

	// Initialize HTTP handlers
	userHandler := http.NewUserHandler(userService)
	filmHandler := http.NewFilmHandler(filmService)

	// Initialize HTTP router
	router, err := http.NewRouter(cfg, authService, *userHandler, *filmHandler)
	if err != nil {
		lg.WithError(err).Fatal("failed to initialize router")
	}

	// Start HTTP server
	listenAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	lg.WithField("listen_address", listenAddr).Info("Starting HTTP server")
	if err := router.Serve(listenAddr); err != nil {
		lg.WithError(err).Fatal("failed to start HTTP server")
	}
}
