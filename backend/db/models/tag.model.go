package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tag struct {
	ID        primitive.ObjectID `bson:"_id"`
	EN        TagLangContent     `bson:"en"`
	ES        TagLangContent     `bson:"es"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

type NewTag struct {
	EN        TagLangContent     `bson:"en"`
	ES        TagLangContent     `bson:"es"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}

type TagLangContent struct {
	Name string `bson:"name"`
}
