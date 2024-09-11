package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID        primitive.ObjectID `bson:"_id"`
	EN        LocationLang       `bson:"en"`
	ES        LocationLang       `bson:"es"`
	URL       string             `bson:"url"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

type NewLocation struct {
	EN        LocationLang `bson:"en"`
	ES        LocationLang `bson:"es"`
	URL       string       `bson:"url"`
	CreatedAt string       `bson:"created_at"`
	UpdatedAt string       `bson:"updated_at"`
}

type LocationLang struct {
	Name     string `bson:"name"`
	Comments string `bson:"comments"`
}
