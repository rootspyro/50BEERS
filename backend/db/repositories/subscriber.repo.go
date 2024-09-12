package repositories

import (
	"context"
	"time"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriberRepo struct {
	Collection *mongo.Collection
}

func NewSubscriberRepo(collection *mongo.Collection) *SubscriberRepo {
	return &SubscriberRepo{
		Collection: collection,	
	}
}

func(r *SubscriberRepo) FindByEmail(email string) (models.Subscriber, error) {
	filter := bson.D{{"email", email}}

	var result models.Subscriber

	err := r.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return models.Subscriber{}, err
	}

	return result, nil
}

func(r *SubscriberRepo) NewSubsciber(email string) (models.Subscriber, error) {
	now := time.Now().Local()
	
	data := models.NewSubscriber {
		Email: email,
		CreatedAt: now.String(),
		UpdatedAt: now.String(),
	}
	
	result, err := r.Collection.InsertOne(context.TODO(), data)
	if err != nil {
		return models.Subscriber{}, err
	}

	newSubscriber := models.Subscriber {
		ID: result.InsertedID.(primitive.ObjectID),
		Email: data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return newSubscriber, nil
} 
