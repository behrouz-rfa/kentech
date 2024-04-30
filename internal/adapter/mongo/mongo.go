package mongo

import (
	"context"
	"time"

	"github.com/behrouz-rfa/kentech/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	logTag              = "mongo"
	DublicatedErrorCode = 11000
)

type Repository struct {
	client  *mongo.Client
	db      *mongo.Database
	timeout time.Duration
	lg      *logger.Entry
}

func NewRepository(uri string, database string, timeout int) (*Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	defer cancel()

	lg := logger.General.Component(logTag)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		lg.WithError(err).Error("Could not connect to mongo")
		return nil, err
	}

	db := client.Database(database)

	return &Repository{
		db:      db,
		client:  client,
		timeout: time.Duration(timeout) * time.Second,
		lg:      lg,
	}, nil
}

func (m Repository) DB() *mongo.Database {
	return m.db
}

func (m Repository) Close() {
	err := m.client.Disconnect(context.Background())
	if err != nil {
		m.lg.WithError(err).Error("Could not disconnect from mongo")
		return
	}
}
