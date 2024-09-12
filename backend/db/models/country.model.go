package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	ID        primitive.ObjectID `bson:"_id"`
	EN        CountryLang        `bson:"en"`
	ES        CountryLang        `bson:"es"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

type NewCountry struct {
	EN        CountryLang        `bson:"en"`
	ES        CountryLang        `bson:"es"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

type CountryLang struct {
	Name string
}
