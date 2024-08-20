package drinks

import "github.com/rootspyro/50BEERS/services"

type Drink struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type DrinksResponse struct {
	ItemsFound int         `json:"itemsFound"`
	Items      []services.Drink `json:"items"`
}
