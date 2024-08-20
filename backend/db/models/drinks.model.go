package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Drink struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Type         string             `bson:"type"`
	ABV          float64            `bson:"abv"`
	CountryID    string             `bson:"country_id"`
	Date         string             `bson:"date"`
	ChallengeNum float64            `bson:"challeng_number"`
	Stars        float64            `bson:"stars"`
	PictureURL   string             `bson:"picture_url"`
	LocationId   string             `bson:"location_id"`
	Tags         []string           `bson:"tags"`
	CreatedAt    string             `bson:"created_at"`
	UpdatedAt    string             `bson:"updated_at"`
	Status       string             `bson:"status"`
}
