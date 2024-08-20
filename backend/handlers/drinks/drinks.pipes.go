package drinks

import "github.com/rootspyro/50BEERS/services"

type DrinksResponse struct {
	ItemsFound     int              `json:"itemsFound"`
	Items          []services.Drink `json:"items"`
	FiltersAllowed []string         `json:"filtersAllowed"`
	FiltersApplied Filters          `json:"filtersApplied"`
}

type Filters struct {
	Name    string `json:"name,omitempty"`
	Country string `json:"country,omitempty"`
}
