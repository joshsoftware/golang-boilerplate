// +build ignore

package db

import (
	"context"
	"joshsoftware/golang-boilerplate/config"
	"time"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoStore struct {
	db *mongo.Database
}

// Singleton instance method accessible from other packages
func Init() (s Storer, err error) {
	uri := config.ReadEnvString("DB_URI")
	name := config.ReadEnvString("DB_NAME")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot initialize database")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot initialize database context")
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot connect to database")
		return
	}

	logger.WithFields(logger.Fields{
		"uri":  uri,
		"name": name,
	}).Info("Connected to mongo database")

	db := client.Database(name)
	return &mongoStore{db}, nil
}
