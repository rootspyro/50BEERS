package db

import (
	"context"
	"os"

	"github.com/rootspyro/50BEERS/config"
	"github.com/rootspyro/50BEERS/config/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI(config.App.Database.URL)

	DBClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	err = DBClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	defer DBClient.Disconnect(context.TODO())

	log.Info("Database connected successfully")
}
