package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(connUrl string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return client, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return client, err
	}

	return client, nil
}
