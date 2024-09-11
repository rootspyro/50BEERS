package repositories

import (
	"context"
	"time"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CountriesRepo struct {
	Collection *mongo.Collection
}

func NewCountriesRepo(collection *mongo.Collection) *CountriesRepo {
	return &CountriesRepo{
		Collection: collection,
	}
}

func(r *CountriesRepo) GetAllCountries() ([]models.Country, error) {

	// make query
	cursor, err := r.Collection.Find(context.TODO(), bson.D{{}})
	
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	//iterate result

	var countries []models.Country

	for cursor.Next(context.TODO()) {
		var country models.Country

		if err := cursor.Decode(&country); err != nil {
			return nil, err
		}

		countries = append(countries, country)
	}

	return countries, nil
}

func(r *CountriesRepo) FindByName(name string) (models.Country, error) {
	
	filter := bson.D{{"en.name", name}}
	
	var result models.Country 

	err := r.Collection.FindOne(context.TODO(),filter).Decode(&result)
	if err != nil {
		return models.Country{}, err
	}

	return result, nil
}

func(r *CountriesRepo) FindById(id primitive.ObjectID) (models.Country, error) {
	filter := bson.D{{"_id", id}}

	var result models.Country 

	err := r.Collection.FindOne(context.TODO(),filter).Decode(&result)
	if err != nil {
		return models.Country{}, err
	}

	return result, nil
}

func(r *CountriesRepo) InsertMany(data []models.NewCountry) (int, error) {

	var documents []interface{}

	//convert data model into interface
	for _, doc := range data {
		doc.CreatedAt = time.Now().Local().String()
		doc.UpdatedAt = time.Now().Local().String()
		documents = append(documents, doc)
	}

	insertResults, err := r.Collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return 0, err
	}

	return len(insertResults.InsertedIDs), nil 
}
