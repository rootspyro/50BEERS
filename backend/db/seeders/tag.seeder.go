package seeders

import (
	"fmt"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
)

type TagSeeder struct {
	repo *repositories.TagRepo
}

func NewTagSeeder(repo *repositories.TagRepo) *TagSeeder{
	return &TagSeeder{
		repo: repo,
	}
}

func(s *TagSeeder) Seed() error {

	log.Info("Running tag seeder")

	data := []models.NewTag{
		{
			EN: models.TagLangContent{
				Name: "beer",
			},
			ES: models.TagLangContent{
				Name: "cerveza",
			},
		},
		{
			EN: models.TagLangContent{
				Name: "challenge",
			},
			ES: models.TagLangContent{
				Name: "reto",
			},
		},
		{
			EN: models.TagLangContent{
				Name: "licours",
			},
			ES: models.TagLangContent{
				Name: "licores",
			},
		},
		{
			EN: models.TagLangContent{
				Name: "others",
			},
			ES: models.TagLangContent{
				Name: "otros",
			},
		},
		{
			EN: models.TagLangContent{
				Name: "rum",
			},
			ES: models.TagLangContent{
				Name: "ron",
			},
		},
		{
			EN: models.TagLangContent{
				Name: "whisky",
			},
			ES: models.TagLangContent{
				Name: "whisky",
			},
		},
		{
			EN: models.TagLangContent{
				Name: "wine",
			},
			ES: models.TagLangContent{
				Name: "vino",
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
