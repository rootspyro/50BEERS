package services

import (
	"fmt"
	"math"
	"strings"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	defaultSortDirection string = "desc"
	allowedDirections           = []string{"asc", "desc"}
)

var (
	defaultSortField  string = "created_at"
	allowedSortFields        = []string{"name", "date", "stars", "abv"}
)

type DrinkSrv struct {
	countryRepo  *repositories.CountriesRepo
	locationRepo *repositories.LocationRepo
	repo         *repositories.DrinksRepo
}

func NewDrinkSrv(
	countryRepo *repositories.CountriesRepo,
	locationRepo *repositories.LocationRepo,
	repo *repositories.DrinksRepo,
) *DrinkSrv {
	return &DrinkSrv{
		countryRepo:  countryRepo,
		locationRepo: locationRepo,
		repo:         repo,
	}
}

func(s *DrinkSrv) CalculatePages(limit int) (int, error) {

	count, err := s.repo.CountAllDrinks()
	if err != nil {
		return 0, err
	}

	var pagesCalc float64 = float64(count) / float64(limit) 

	pages := math.Ceil(pagesCalc)

	return int(pages), nil
} 

func (s *DrinkSrv) GetAllDrinks(filters DrinkSearchFilters) ([]DrinkResume, error) {
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

	// Get location Id
	var locationId string

	location, err := s.locationRepo.FindByName(filters.Location)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			locationId = ""
		} else {
			return nil, err
		}
	} else {
		locationId = location.ID.Hex()
	}

	// build search filter 
	searchFilter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{{"tags", bson.D{{"$regex", fmt.Sprintf(".*%s.*", filters.Category)}}}},
				bson.D{{"country_id", bson.D{{"$regex", fmt.Sprintf(".*%s.*", countryId)}}}},
				bson.D{{
					"$or",
					bson.A{
						bson.D{{"name", bson.D{{"$regex", nameRegex}}}},
						bson.D{{"type", bson.D{{"$regex", nameRegex}}}},
					},
				}},
				bson.D{{"location_id", bson.D{{"$regex", fmt.Sprintf(".*%s.*", locationId)}}}},
				bson.D{{"status", "public"}},
			},
		},
	}

	// build sort filter
	var direction int = -1
	var field string = defaultSortField

	if fieldIsValid(filters.SortBy) {

		field = filters.SortBy

	} else if filters.SortBy != ""{

		log.Warning(fmt.Sprintf("invalid sort field: %s", filters.SortBy))
	}

	if filters.Direction == "asc" {

		direction = 1

	} else if filters.Direction == "desc" {

		direction = -1

	} else if filters.Direction != ""{

		log.Warning(fmt.Sprintf("invalid sort direction: %s", filters.Direction))
	}

	skip := (filters.Page - 1) * filters.Limit

	sortFilter := options.Find().SetSort(bson.D{
		{Key: field, Value: direction}, 
	}).SetSkip(int64(skip)).SetLimit(int64(filters.Limit))

	response, err := s.repo.GetAllDrinks(searchFilter, sortFilter)

	var drinks []DrinkResume

	for _, drink := range response {
		drinks = append(drinks, parseResumeDrink(drink))
	}

	return drinks, err
}

func parseDrink(data models.Drink) Drink {
	newDrink := Drink{
		ID:           data.ID.Hex(),
		Name:         data.Name,
		Type:         data.Type,
		ABV:          data.ABV,
		CountryID:    data.CountryID,
		Date:         data.Date,
		ChallengeNum: data.ChallengeNum,
		Stars:        data.Stars,
		PictureURL:   data.PictureURL,
		LocationId:   data.LocationId,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		Status:       data.Status,
	}

	for _, tag := range data.Tags {
		newDrink.Tags = append(newDrink.Tags, tag)
	}

	return newDrink
}

func parseResumeDrink(data models.Drink) DrinkResume {
	newDrink := DrinkResume{
		ID:           parseDrinkId(data.Name),
		Name:         data.Name,
		Type:         data.Type,
		ABV:          data.ABV,
		Date:         data.Date,
		ChallengeNum: data.ChallengeNum,
		Stars:        data.Stars,
		PictureURL:   data.PictureURL,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}

	return newDrink
}

func fieldIsValid(field string) bool {
	var valid bool = false

	for _, allowedField := range allowedSortFields {
		if field == allowedField {
			valid = true
			break
		}
	}

	return valid
}

func parseDrinkId(name string) string {
	return strings.ReplaceAll(name, " ", "_")
}

type Drink struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	ABV          float64  `json:"abv"`
	CountryID    string   `json:"country_id"`
	Date         string   `json:"date"`
	ChallengeNum float64  `json:"challeng_number"`
	Stars        float64  `json:"stars"`
	PictureURL   string   `json:"picture_url"`
	LocationId   string   `json:"location_id"`
	Tags         []string `json:"tags"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	Status       string   `json:"status"`
}

type DrinkResume struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	ABV          float64 `json:"abv"`
	Date         string  `json:"date"`
	ChallengeNum float64 `json:"challeng_number"`
	Stars        float64 `json:"stars"`
	PictureURL   string  `json:"picture_url"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type DrinkLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type DrinkSearchFilters struct {
	Name      string
	Category  string
	Country   string
	Location  string
	SortBy    string
	Direction string
	Page      int
	Limit     int
}
