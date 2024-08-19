package drinks

import (
	srv "github.com/rootspyro/50BEERS/services/drinks"
)

type Drink struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type DrinksResponse struct {
	ItemsFound int         `json:"itemsFound"`
	Items      []srv.Drink `json:"items"`
}
