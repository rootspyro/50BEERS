package routes

import (
	"net/http"

	"github.com/rootspyro/50BEERS/handlers/drinks"
)

var drinksHandler drinks.DrinksHandler = *drinks.New()

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
