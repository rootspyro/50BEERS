package services

import (
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type LocationSrv struct {
	repo *repositories.LocationRepo
}

func NewLocationSrv(repo *repositories.LocationRepo) *LocationSrv {
	return &LocationSrv{
		repo: repo,
	}
}

func (s *LocationSrv) GetAllLocations(lang string) ([]Location, error) {
	data, err := s.repo.GetAllLocations()
	if err != nil {
		return nil, err
	}

	var locations []Location
	for _, location := range data {
		locations = append(locations, parseLocation(location, lang))
	}

	return locations, nil
}

func parseLocation(data models.Location, lang string) Location {
	var name string = data.EN.Name
	var comments string = data.EN.Comments

	if lang == "es" {
		name = data.ES.Name
		comments = data.ES.Comments
	}

	return Location{
		ID:          ParsePublicId(data.EN.Name),
		Name:        cases.Title(language.Und).String(name),
		Coordinates: data.Coordinates,
		Comments:    comments,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

type Location struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
	Comments    string    `json:"comments"`
	CreatedAt   string    `json:"createdAt"`
	UpdatedAt   string    `json:"updatedAt"`
}
