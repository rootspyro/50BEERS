package services

import (
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CountrySrv struct {
	repo *repositories.CountriesRepo
}

func NewCountrySrv(repo *repositories.CountriesRepo) *CountrySrv {
	return &CountrySrv{
		repo: repo,
	}
}

func (s *CountrySrv) GetAllCountries() ([]Country, error) {
	data, err := s.repo.GetAllCountries()
	
	if err != nil {
		return nil, err
	}

	// parse data
	var countries []Country

	for _, country := range data {
		countries = append(countries, parseCountry(country))
	}

	return countries, nil
}

func parseCountry(data models.Country) Country {
	return Country{
		ID: ParsePublicId(data.Name),
		Name: cases.Title(language.Und).String(data.Name),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
} 

type Country struct {
	ID        string `json:"id"`
	Name      string `bson:"name"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}
