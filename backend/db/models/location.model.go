package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	URL       string             `bson:"url"`
	Comments  string             `bson:"comments"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
}

type NewLocation struct {
	Name      string `bson:"name"`
	URL       string `bson:"url"`
	Comments  string `bson:"comments"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}
