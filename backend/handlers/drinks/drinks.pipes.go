package drinks

import "github.com/rootspyro/50BEERS/services"

type DrinksResponse struct {
	ItemsFound     int                    `json:"itemsFound"`
	Items          []services.DrinkResume `json:"items"`
	Pagination     Pagination             `json:"pagination"`
	FiltersAllowed []string               `json:"filtersAllowed"`
	FiltersApplied Filters                `json:"filtersApplied"`
}

type Filters struct {
	Name      string `json:"name,omitempty"`
	Country   string `json:"country,omitempty"`
	Location  string `json:"location,omitempty"`
	SortBy    string `json:"sortBy,omitempty"`
	Direction string `json:"direction,omitempty"`
	Category  string `json:"Category,omitempty"`
}

type Pagination struct {
	Pages      int `json:"pages,omitempty"`
	Page       int `json:"page,omitempty"`
	PageSize   int `json:"pageSize,omitempty"`
}
