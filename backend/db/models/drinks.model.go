package models

import (
	"context"

	"github.com/rootspyro/50BEERS/config/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DrinkModel struct {
	Collection *mongo.Collection
}

func NewDrinkModel(collection *mongo.Collection) DrinkModel {
	return DrinkModel{
		Collection: collection,
	}
}

func(m DrinkModel) GetAllDrinks() {
	filters := bson.D{}

	cursor, err := m.Collection.Find(context.TODO(), filters)

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer cursor.Close(context.TODO())

	return
}
