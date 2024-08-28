package location

import "github.com/rootspyro/50BEERS/services"

type LocationsResponse struct {
	ItemsFound int                 `json:"itemsFound"`
	Items      []services.Location `json:"items"`
}
