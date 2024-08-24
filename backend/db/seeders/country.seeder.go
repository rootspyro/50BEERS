package seeders

import (
	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db/repositories"
)

type CountrySeeder struct {
	repo *repositories.CountriesRepo
}

func NewCountrySeeder(repo *repositories.CountriesRepo) *CountrySeeder{
	return &CountrySeeder{
		repo: repo,
	}
}

func Seed() error {

	log.Info("Running country seeder...")
	return nil
}
