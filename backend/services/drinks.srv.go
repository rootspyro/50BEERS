package services

import (
	"fmt"

	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DrinkSrv struct {
	countryRepo *repositories.CountriesRepo
	repo        *repositories.DrinksRepo
}

func NewDrinkSrv(countryRepo *repositories.CountriesRepo, repo *repositories.DrinksRepo) *DrinkSrv {
	return &DrinkSrv{
		countryRepo: countryRepo,
		repo:        repo,
	}
}

func (s *DrinkSrv) GetAllDrinks(filters DrinkSearchFilters) ([]Drink, error) {
	nameRegex := fmt.Sprintf(".*%s.*", filters.Name)

	// Get Country Id
	var countryId string

	country, err := s.countryRepo.FindByName(filters.Country)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			countryId = ""
		} else {
			return nil, err
		}	
	} else {
		countryId = country.ID.Hex()
	}

	response, err := s.repo.GetAllDrinks(
		bson.D{
			{
				"$and",
				bson.A{
					bson.D{{"country_id", countryId}},
					bson.D{{
						"$or",
						bson.A{
							bson.D{{"name", bson.D{{"$regex", nameRegex}}}},
							bson.D{{"type", bson.D{{"$regex", nameRegex}}}},
						},
					}},
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
		ID:           data.ID.Hex(),
		Name:         data.Name,
		Type:         data.Type,
		ABV:          data.ABV,
		CountryID:    data.CountryID.Hex(),
		Date:         data.Date,
		ChallengeNum: data.ChallengeNum,
		Stars:        data.Stars,
		PictureURL:   data.PictureURL,
		Location:     DrinkLocation(data.Location),
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		Status:       data.Status,
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
	Name    string
	Country string
}
