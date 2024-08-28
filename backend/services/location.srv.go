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

func (s *LocationSrv) GetAllLocations() ([]Location, error) {
	data, err := s.repo.GetAllLocations()
	if err != nil {
		return nil, err 
	}

	var locations []Location
	for _, location := range data {

		locations = append(locations, parseLocation(location))
	}

	return locations, nil
}

func parseLocation(data models.Location) Location {
	return Location{
		ID: ParsePublicId(data.Name),
		Name: cases.Title(language.Und).String(data.Name),
		URL: data.URL,
		Comments: data.Comments,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

type Location struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Comments  string `json:"comments"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
