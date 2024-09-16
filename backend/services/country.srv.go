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

func (s *CountrySrv) GetAllCountries(lang string) ([]Country, error) {
	data, err := s.repo.GetAllCountries()
	
	if err != nil {
		return nil, err
	}

	// parse data
	var countries []Country

	for _, country := range data {
		countries = append(countries, parseCountry(country, lang))
	}

	return countries, nil
}

func parseCountry(data models.Country, lang string) Country {
	var name string = data.EN.Name

	if lang == "es" {
		name = data.ES.Name
	}

	return Country{
		ID: ParsePublicId(data.EN.Name),
		Name: cases.Title(language.Und).String(name),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
} 

type Country struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
