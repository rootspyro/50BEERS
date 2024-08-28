package tag

import "github.com/rootspyro/50BEERS/services"

type TagsResponse struct {
	ItemsFound int            `json:"itemsFound"`
	Items      []services.Tag `json:"items"`
}
