// +build ignore

// IMPORTANT: If you want to use MongoDB, remove the above build tag.
// Add this build tag to the other database drivers (pg.go).
// Remember to add a newline after adding the build tag !
package db

import (
	"context"
	"joshsoftware/golang-boilerplate/config"
	"time"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		logger.Error("Cannot initialize database", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.Error("Cannot initialize database context", err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		logger.Error("Cannot connect to database", err)
		return
	}

	logger.Info("Connected To MongoDB")
	db := client.Database(name)

	return &mongoStore{db}, nil
}

// TODO: delete
func userIDFromContext(ctx context.Context) (userID primitive.ObjectID) {
	userid := ""
	if ctx.Value("UserID") != nil { // verify it exists
		userid = ctx.Value("UserID").(string)
	}

	if userid == "" {
		logger.Error("User not specified in context")
		return
	}

	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		logger.Errorf("UserID is invalid: %v", err)
		return
	}

	return
}
