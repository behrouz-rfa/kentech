//go:build integration

package mongo

import (
	"context"
	"testing"

	"github.com/behrouz-rfa/kentech/internal/adapter/config"
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo"
	"github.com/behrouz-rfa/kentech/pkg/logger"
	"github.com/stretchr/testify/suite"
)

func TestMongoTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTestSuite))
}

type MongoTestSuite struct {
	suite.Suite
	db *mongo.Repository
}

func (s *MongoTestSuite) SetupSuite() {
	logger.Init()
	config.Load()
	conf := config.Get()
	db, err := mongo.NewRepository(conf.DbConnectionString(), conf.Database.Name, config.DbTimeout)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
}

func (s *MongoTestSuite) TearDownSuite() {
	s.db.DB().Drop(context.Background()) //nolint:errcheck
	s.db.Close()
}
