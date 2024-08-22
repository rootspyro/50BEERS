package repositories

import (
	"context"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationRepo struct {
	Collection *mongo.Collection
}

func NewLocationRepo(collection *mongo.Collection) *LocationRepo{
	return &LocationRepo{
		Collection: collection,
	}
}

func(r *LocationRepo) FindByName(name string) (models.Location, error) {

	filter := bson.D{{"name", name}}

	var location models.Location

	err := r.Collection.FindOne(context.TODO(), filter).Decode(&location)

	if err != nil {
		return models.Location{}, err
	}

	return location, nil
} 
