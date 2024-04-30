//go:build e2e

package gql

import (
	"context"
	"os"
	"testing"

	"github.com/behrouz-rfa/kentech/internal/adapter/auth"
	"github.com/behrouz-rfa/kentech/internal/adapter/config"
	handler "github.com/behrouz-rfa/kentech/internal/adapter/handlers/http"
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo"
	"github.com/behrouz-rfa/kentech/internal/core/services"
	"github.com/behrouz-rfa/kentech/pkg/logger"
	"github.com/stretchr/testify/suite"
)

func TestGqlTestSuite(t *testing.T) {
	suite.Run(t, new(GqlTestSuite))
}

type GqlTestSuite struct {
	suite.Suite
	dbRepo *mongo.Repository
	api    *handler.Router
}

func (s *GqlTestSuite) SetupSuite() {
	logger.Init()
	config.Load()
	cfg := config.Get()

	db, err := mongo.NewRepository(cfg.DbConnectionString(), cfg.Database.Name, config.DbTimeout)
	if err != nil {
		s.T().Fatal(err)
	}

	// Read the firebase.json file
	auth := auth.NewAuth(cfg.Jwt.Secret)

	userService := services.NewUserService(
		services.WithUserRepository(db),
		services.WithAuth(auth),
	)

	filmService := services.NewFilmService(
		services.WithFilmRepository(db),
	)

	userHandler := handler.NewUserHandler(userService)
	filmHandler := handler.NewFilmHandler(filmService)

	// Init router
	router, err := handler.NewRouter(
		cfg,
		auth,
		*userHandler,
		*filmHandler,
	)
	if err != nil {
		os.Exit(1)
	}

	s.dbRepo = db
	s.api = router
}

func (s *GqlTestSuite) TearDownSuite() {
	s.dbRepo.DB().Drop(context.Background())
	s.dbRepo.Close()
}
