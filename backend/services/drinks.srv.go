package services

import (
	"fmt"
	"math"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func (s *DrinkSrv) CountDrinksForBlog() (DrinksCount, error) {

	var result DrinksCount

	all, err := s.repo.CountDrinks(bson.D{})
	if err != nil {
		return result, err 
	}

	whiskys, err := s.repo.CountDrinks(bson.D{{"tags", "whisky"}})
	if err != nil {
		return result, err
	}

	rums, err := s.repo.CountDrinks(bson.D{{"tags", "rum"}})
	if err != nil {
		return result, err
	}

	beers, err := s.repo.CountDrinks(bson.D{{"tags", "beer"}})
	if err != nil {
		return result, err
	}

	wines, err := s.repo.CountDrinks(bson.D{{"tags", "wines"}})
	if err != nil {
		return result, err
	}

	others, err := s.repo.CountDrinks(
		bson.D{{
			"$and",
			bson.A{
				bson.M{"tags": bson.M{"$ne": "whisky"}},
				bson.M{"tags": bson.M{"$ne": "rum"}},
				bson.M{"tags": bson.M{"$ne": "beer"}},
				bson.M{"tags": bson.M{"$ne": "wine"}},
			},
		}},
	)

	if err != nil {
		return result, err
	}

	result.All = all
	result.Whiskys = whiskys
	result.Rums = rums
	result.Beers = beers
	result.Wines = wines
	result.Others = others

	return result, nil
}

func (s *DrinkSrv) CalculatePages(limit int) (int, error) {
	count, err := s.repo.CountDrinks(bson.D{})
	if err != nil {
		return 0, err
	}

	var pagesCalc float64 = float64(count) / float64(limit)

	pages := math.Ceil(pagesCalc)

	return int(pages), nil
}

func (s *DrinkSrv) GetAllDrinks(filters DrinkSearchFilters, lang string) ([]DrinkResume, error) {
	nameRegex := fmt.Sprintf(".*%s.*", filters.Name)

	// validate country exists
	_, err := s.countryRepo.FindByName(filters.Country)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			filters.Country = ""
		} else {
			return nil, err
		}
	}

	// validate if location exists
	_, err = s.locationRepo.FindByName(filters.Location)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			filters.Location = ""
		} else {
			return nil, err
		}
	}

	// build search filter
	searchFilter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{{"tags", bson.D{{"$regex", fmt.Sprintf(".*%s.*", filters.Category)}}}},
				bson.D{{"en.country", bson.D{{"$regex", fmt.Sprintf(".*%s.*", filters.Country)}}}},
				bson.D{{
					"$or",
					bson.A{
						bson.D{{"name", bson.D{{"$regex", nameRegex}}}},
						bson.D{{"type", bson.D{{"$regex", nameRegex}}}},
					},
				}},
				bson.D{{"en.location", bson.D{{"$regex", fmt.Sprintf(".*%s.*", filters.Location)}}}},
				bson.D{{"status", "public"}},
			},
		},
	}

	// build sort filter
	var direction int = -1

	if filters.Direction == "asc" {
		direction = 1
	} else if filters.Direction == "desc" {
		direction = -1
	} else if filters.Direction != "" {
		log.Warning(fmt.Sprintf("invalid sort direction: %s", filters.Direction))
	}

	skip := (filters.Page - 1) * filters.Limit

	sortFilter := options.Find().SetSort(bson.D{
		{Key: filters.SortBy, Value: direction},
	}).SetSkip(int64(skip)).SetLimit(int64(filters.Limit))

	response, err := s.repo.GetAllDrinks(searchFilter, sortFilter)

	var drinks []DrinkResume

	for _, drink := range response {
		drinks = append(drinks, parseResumeDrink(drink, lang))
	}

	return drinks, err
}

func parseDrink(data models.Drink, lang string) Drink {

	var country string = data.EN.Country	
	var location string = data.EN.Location

	if lang == "es" {
		country = data.ES.Country
		location = data.ES.Location
	}

	newDrink := Drink{
		ID:           data.ID.Hex(),
		Name:         data.Name,
		Type:         data.Type,
		ABV:          data.ABV,
		Country:      country,
		Date:         data.Date,
		ChallengeNum: data.ChallengeNum,
		Stars:        data.Stars,
		PictureURL:   data.PictureURL,
		Location:     location,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		Status:       data.Status,
	}

	for _, tag := range data.Tags {
		newDrink.Tags = append(newDrink.Tags, tag)
	}

	return newDrink
}

func parseResumeDrink(data models.Drink, lang string) DrinkResume {

	var country string = data.EN.Country	
	var location string = data.EN.Location

	if lang == "es" {
		country = data.ES.Country
		location = data.ES.Location
	}

	newDrink := DrinkResume{
		ID:           ParsePublicId(data.Name),
		Name:         cases.Title(language.Und).String(data.Name),
		Type:         cases.Title(language.Und).String(data.Type),
		ABV:          data.ABV,
		Country:      cases.Title(language.Und).String(country),
		Date:         data.Date,
		ChallengeNum: data.ChallengeNum,
		Location:     cases.Title(language.Und).String(location),
		Stars:        data.Stars,
		PictureURL:   data.PictureURL,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}

	return newDrink
}

type Drink struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	ABV          float64  `json:"abv"`
	Country      string   `json:"country"`
	Date         string   `json:"date"`
	ChallengeNum float64  `json:"challengeNumber"`
	Stars        float64  `json:"stars"`
	PictureURL   string   `json:"pictureUrl"`
	Location     string   `json:"location"`
	Tags         []string `json:"tags"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
	Status       string   `json:"status"`
}

type DrinkResume struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	ABV          float64 `json:"abv"`
	Country      string  `json:"country"`
	Date         string  `json:"date"`
	ChallengeNum float64 `json:"challengeNumber"`
	Stars        float64 `json:"stars"`
	PictureURL   string  `json:"pictureUrl"`
	Location     string  `json:"location"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
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

type DrinksCount struct {
	All     int64 `json:"all"`
	Whiskys int64 `json:"whiskys"`
	Rums    int64 `json:"rums"`
	Beers   int64 `json:"beers"`
	Wines   int64 `json:"wines"`
	Others  int64 `json:"others"`
}
