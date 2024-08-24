package seeders

import (
	"errors"
	"fmt"

	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
)

func SeedCollection(
	collectionName string,
	countryRepo *repositories.CountriesRepo,
	locationRepo *repositories.LocationRepo,
	drinksRepo *repositories.DrinksRepo,
) error {

	if !validCollection(collectionName) {
		return errors.New("Collection was not found")
	}

	switch collectionName {
	case "country":

		countrySeeder := NewCountrySeeder(countryRepo)
		if err := countrySeeder.Seed(); err != nil {
			return err
		}

		break;

	default:
		return errors.New(fmt.Sprintf("%s collection currently has no seeder.", collectionName))
	}
	
	return nil
}

func validCollection(collectionName string) bool {

	for _, collection := range models.Collections {
		if collection == collectionName {
			return true
		}
	}

	return false 
}
