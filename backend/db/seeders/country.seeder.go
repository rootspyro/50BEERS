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
			EN: models.CountryLang{
				Name: "belgium",
			},
			ES: models.CountryLang{
				Name: "bélgica",
			},
		},
		{
			EN: models.CountryLang{
				Name: "germany",
			},
			ES: models.CountryLang{
				Name: "alemania",
			},
		},
		{
			EN: models.CountryLang{
				Name: "ireland",
			},
			ES: models.CountryLang{
				Name: "irlanda",
			},
		},
		{
			EN: models.CountryLang{
				Name: "mexico",
			},
			ES: models.CountryLang{
				Name: "méxico",
			},
		},
		{
			EN: models.CountryLang{
				Name: "scotland",
			},
			ES: models.CountryLang{
				Name: "escocia",
			},
		},
		{
			EN: models.CountryLang{
				Name: "spain",
			},
			ES: models.CountryLang{
				Name: "españa",
			},
		},
		{
			EN: models.CountryLang{
				Name: "united states",
			},
			ES: models.CountryLang{
				Name: "estados unidos",
			},
		},
		{
			EN: models.CountryLang{
				Name: "venezuela",
			},
			ES: models.CountryLang{
				Name: "venezuela",
			},
		},
	}

	docsCreated, err := s.repo.InsertMany(data)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("%d records successfully created", docsCreated))

	return nil
}
