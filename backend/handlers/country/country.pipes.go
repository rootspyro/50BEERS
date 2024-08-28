package country

import "github.com/rootspyro/50BEERS/services"

type CountriesResponse struct {
	ItemsFound int                `json:"itemsFound"`
	Items      []services.Country `json:"items"`
}
