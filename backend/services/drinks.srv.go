package services

import (
	"fmt"

	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DrinkSrv struct {
	model *models.DrinkModel
}

func NewDrinkSrv(model *models.DrinkModel) *DrinkSrv {
	return &DrinkSrv{
		model: model,
	}
}

func (s *DrinkSrv) GetAllDrinks(filters DrinkSearchFilters) ([]Drink, error) {

	nameRegex := fmt.Sprintf(".*%s.*", filters.Name)

	response, err := s.model.GetAllDrinks(
		bson.D{
				{"$or",
						bson.A{
								bson.D{{"name", primitive.Regex{Pattern: nameRegex}}},
								bson.D{{"type", primitive.Regex{Pattern: nameRegex}}},
						},
				},
		},
	)

	var drinks []Drink

	for _, drink := range response {
		drinks = append(drinks, parseDrink(drink))
	}

	return drinks, err
}

func parseDrink(data models.Drink) Drink {
	newDrink := Drink{
		ID: data.ID.Hex(),
		Name: data.Name,
		Type: data.Type,
		ABV: data.ABV,
		CountryID: data.CountryID.Hex(),
		Date: data.Date,
		ChallengeNum: data.ChallengeNum,
		Stars: data.Stars,
		PictureURL: data.PictureURL,
		Location: DrinkLocation(data.Location),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Status: data.Status,
	}

	for _, tagId := range data.TagIds {
		newDrink.TagIds = append(newDrink.TagIds, tagId.Hex())	
	}

	return newDrink
}

type Drink struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	ABV          float64       `json:"abv"`
	CountryID    string        `json:"country_id"`
	Date         string        `json:"date"`
	ChallengeNum float64       `json:"challeng_number"`
	Stars        float64       `json:"stars"`
	PictureURL   string        `json:"picture_url"`
	Location     DrinkLocation `json:"location"`
	TagIds       []string      `json:"tag_ids"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
	Status       string        `json:"status"`
}

type DrinkLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type DrinkSearchFilters struct {
	Name string
}
