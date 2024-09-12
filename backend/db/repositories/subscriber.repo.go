package repositories

import "go.mongodb.org/mongo-driver/mongo"

type SubscriberRepo struct {
	Collection *mongo.Collection
}

func NewSubscriberRepo(collection *mongo.Collection) *SubscriberRepo {
	return &SubscriberRepo{
		Collection: collection,	
	}
}
