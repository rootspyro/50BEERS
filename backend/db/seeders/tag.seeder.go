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
			Name: "beer",
		},
		{
			Name: "challenge",
		},
		{
			Name: "licours",
		},
		{
			Name: "others",
		},
		{
			Name: "rum",
		},
		{
			Name: "whisky",
		},
		{
			Name: "wine",
		},
	}

	docsCreated, err := s.repo.InsertMany(data)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("%d records successfully created", docsCreated))

	return nil
}
