package repositories

import (
	"context"
	"time"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepo struct {
	Collection *mongo.Collection
}

func NewTagRepo(collection *mongo.Collection) *TagRepo{
	return &TagRepo{
		Collection: collection,
	}
}

func(r *TagRepo) GetAllTags() ([]models.Tag, error) {
	cursor, err := r.Collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var tags []models.Tag
	for cursor.Next(context.TODO()) {

		var tag models.Tag

		err := cursor.Decode(&tag)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func(r *TagRepo) InsertMany(data []models.NewTag) (int, error) {

	var documents []interface{}

	for _, tag := range data {
		tag.CreatedAt = time.Now().Local().String()
		tag.UpdatedAt = time.Now().Local().String()
		documents = append(documents, tag)
	}

	results, err := r.Collection.InsertMany(context.TODO(), documents)

	if err != nil {
		return 0, err
	}
	
	return len(results.InsertedIDs), nil
}
