package repositories

import (
	"context"
	"time"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogUserRepo struct {
	Collection *mongo.Collection
}

func NewBlogUserRepo(collection *mongo.Collection) *BlogUserRepo {
	return &BlogUserRepo{
		Collection: collection,
	}
}

func(r *BlogUserRepo) CreateUser(data models.NewBlogUser) (models.BlogUser, error) {

	now := time.Now().Local()

	var newUser = models.BlogUser{
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
		Origin: data.Origin,
		CreatedAt: now.String(),
		UpdatedAt: now.String(),
	}

	result, err := r.Collection.InsertOne(context.TODO(), data)
	if err != nil {
		return newUser, err
	}

	newUser.ID = result.InsertedID.(primitive.ObjectID)

	return newUser, nil
}
