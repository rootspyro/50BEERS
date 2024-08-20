package drinks

import "github.com/rootspyro/50BEERS/services"

type DrinksResponse struct {
	ItemsFound int         `json:"itemsFound"`
	Items      []services.Drink `json:"items"`
}
