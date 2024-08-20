package models

import (
	"context"

	"github.com/rootspyro/50BEERS/config/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m DrinkModel) GetAllDrinks() ([]Drink, error) {
	filters := bson.D{}

	// make query
	cursor, err := m.Collection.Find(context.TODO(), filters)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer cursor.Close(context.TODO())

	// iterate result
	var drinks []Drink

	for cursor.Next(context.TODO()) {
		var drink Drink

		if err := cursor.Decode(&drink); err != nil {
			return nil, err
		}

		drinks = append(drinks, drink)
	}

	return drinks, nil
}

type Drink struct {
	ID           primitive.ObjectID   `bson:"_id"`
	Name         string               `bson:"name"`
	Type         string               `bson:"type"`
	ABV          float64              `bson:"abv"`
	CountryID    primitive.ObjectID   `bson:"country_id"`
	Date         string               `bson:"date"`
	ChallengeNum float64              `bson:"challeng_number"`
	Stars        float64              `bson:"stars"`
	PictureURL   string               `bson:"picture_url"`
	Location     DrinkLocation        `bson:"location"`
	TagIds       []primitive.ObjectID `bson:"tag_ids"`
	CreatedAt    string               `bson:"created_at"`
	UpdatedAt    string               `bson:"updated_at"`
	Status       string               `bson:"status"`
}

type DrinkLocation struct {
	Name string `bson:"name"`
	URL  string `bson:"url"`
}
