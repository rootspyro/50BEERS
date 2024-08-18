package routes

import (
	"net/http"

	handler "github.com/rootspyro/50BEERS/handlers/health"
	"github.com/rootspyro/50BEERS/middlewares"
)

var HealthRouter Router = Router{
	Basepath: "/health",
	Routes: []Route{
		{
			Path: "",
			Method: http.MethodGet,
			Handler: handler.ServerStatus,
			Middlewares: []Middleware{
				middlewares.Example1,
				middlewares.Example2,
			},
		},
	},
}

