package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscriber struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"created_at"`
}

type NewSubscriber struct {
	Email     string             `bson:"email"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"created_at"`
}
