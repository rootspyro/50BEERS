package seeders

import (
	"fmt"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db/models"
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

func(s *CountrySeeder) Seed() error {

	log.Info("Running country seeder...")

	data := []models.NewCountry{
		{
			Name: "belgium",
		},
		{
			Name: "germany",
		},
		{
			Name: "ireland",
		},
		{
			Name: "mexico",
		},
		{
			Name: "scotland",
		},
		{
			Name: "spain",
		},
		{
			Name: "united states",
		},
		{
			Name: "venezuela",
		},
	}

	docsCreated, err := s.repo.InsertMany(data)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("%d records successfully created", docsCreated))

	return nil
}
