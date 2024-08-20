package repositories

import (
	"context"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CountriesRepo struct {
	Collection *mongo.Collection
}

func NewCountriesRepo(collection *mongo.Collection) *CountriesRepo {
	return &CountriesRepo{
		Collection: collection,
	}
}

func(r *CountriesRepo) FindByName(name string) (models.Country, error) {
	
	filter := bson.D{{"name", name}}
	
	var result models.Country 

	err := r.Collection.FindOne(context.TODO(),filter).Decode(&result)
	if err != nil {
		return models.Country{}, err
	}

	return result, nil
}

func(r *CountriesRepo) FindById(id primitive.ObjectID) (models.Country, error) {
	filter := bson.D{{"_id", id}}

	var result models.Country 

	err := r.Collection.FindOne(context.TODO(),filter).Decode(&result)
	if err != nil {
		return models.Country{}, err
	}

	return result, nil
}

func parseCountry(data models.Country) models.Country {
	return models.Country{
		ID: data.ID,
		Name: data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

