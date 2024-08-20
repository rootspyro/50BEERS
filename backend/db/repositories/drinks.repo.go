package repositories

import (
	"context"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DrinksRepo struct {
	Collection *mongo.Collection
}

func NewDrinksRepo(collection *mongo.Collection) *DrinksRepo {
	return &DrinksRepo{
		Collection: collection,
	}
}

func (m DrinksRepo) GetAllDrinks(filters bson.D) ([]models.Drink, error) {

	// make query
	cursor, err := m.Collection.Find(context.TODO(), filters)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer cursor.Close(context.TODO())

	// iterate result
	var drinks []models.Drink

	for cursor.Next(context.TODO()) {
		var drink models.Drink 

		if err := cursor.Decode(&drink); err != nil {
			return nil, err
		}

		drinks = append(drinks, drink)
	}

	return drinks, nil
}

