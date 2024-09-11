package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Drink struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Type         string             `bson:"type"`
	EN           DrinkLang          `bson:"en"`
	ES           DrinkLang          `bson:"es"`
	ABV          float64            `bson:"abv"`
	Date         string             `bson:"date"`
	ChallengeNum float64            `bson:"challenge_number"`
	Stars        float64            `bson:"stars"`
	PictureURL   string             `bson:"picture_url"`
	Tags         []string           `bson:"tags"`
	CreatedAt    string             `bson:"created_at"`
	UpdatedAt    string             `bson:"updated_at"`
	Status       string             `bson:"status"`
}

type DrinkLang struct {
	Country  string `bson:"country"`
	Location string `bson:"location"`
}
