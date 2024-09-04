package repositories

import (
	"context"
	"time"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *BlogUserRepo) FindByUsername(username string) (models.BlogUser, error) {
	filter := bson.D{{"username", username}}	

	var user models.BlogUser

	err := r.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err 
	}

	return user, nil
}

func (r *BlogUserRepo) FindByEmail(email string) (models.BlogUser, error) {
	filter := bson.D{{"email", email}}	

	var user models.BlogUser

	err := r.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err 
	}

	return user, nil
}

func(r *BlogUserRepo) CreateUser(data models.NewBlogUser) (models.BlogUser, error) {

	now := time.Now().Local()

	data.CreatedAt = now.String()
	data.UpdatedAt = now.String()


	result, err := r.Collection.InsertOne(context.TODO(), data)
	if err != nil {

		return models.BlogUser{}, err
	}

	var newUser = models.BlogUser{
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
		Origin: data.Origin,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	newUser.ID = result.InsertedID.(primitive.ObjectID)

	return newUser, nil
}
