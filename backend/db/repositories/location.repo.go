package repositories

import (
	"context"
	"time"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationRepo struct {
	Collection *mongo.Collection
}

func NewLocationRepo(collection *mongo.Collection) *LocationRepo{
	return &LocationRepo{
		Collection: collection,
	}
}

func(r *LocationRepo) GetAllLocations() ([]models.Location, error) {

	// make query
	cursor, err := r.Collection.Find(context.TODO(), bson.D{{}})
	
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	// iterate result
	var locations []models.Location

	for cursor.Next(context.TODO()) {
		var location models.Location

		if err := cursor.Decode(&location); err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	return locations, nil
}

func(r *LocationRepo) FindByName(name string) (models.Location, error) {

	filter := bson.D{{"name", name}}

	var location models.Location

	err := r.Collection.FindOne(context.TODO(), filter).Decode(&location)

	if err != nil {
		return models.Location{}, err
	}

	return location, nil
} 

func(r *LocationRepo) InsertMany(data []models.NewLocation) (int, error) {

	var documents []interface{}

	// convert locations into documents
	for _, location := range data {
		location.CreatedAt = time.Now().Local().String()
		location.UpdatedAt = time.Now().Local().String()
		documents = append(documents, location)
	}

	result, err := r.Collection.InsertMany(context.TODO(), documents)

	if err != nil {
		return 0, err
	}

	return len(result.InsertedIDs), nil
}
