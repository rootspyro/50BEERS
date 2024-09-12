package repositories

import (
	"context"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriberRepo struct {
	Collection *mongo.Collection
}

func NewSubscriberRepo(collection *mongo.Collection) *SubscriberRepo {
	return &SubscriberRepo{
		Collection: collection,	
	}
}

func(r *SubscriberRepo) FindByEmail(email string) (models.Subscriber, error) {
	filter := bson.D{{"email", email}}

	var result models.Subscriber

	err := r.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return models.Subscriber{}, err
	}

	return result, nil
}
