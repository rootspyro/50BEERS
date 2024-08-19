package routes

import (
	"net/http"

	drinksHandler "github.com/rootspyro/50BEERS/handlers/drinks"
)


var DrinksRouter Router = Router{
	Basepath: "/drinks",
	Routes: []Route{
		{
			Method: http.MethodGet,
			Path: "/blog",
			Handler: drinksHandler.ListDrinksForBlog,
		},
	},
}
