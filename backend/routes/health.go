package routes

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/middlewares"
)

var HealthRouter Router = Router{
	Basepath: "/health",
	Routes: []Route{
		{
			Path: "",
			Method: http.MethodGet,
			Handler: func(w http.ResponseWriter, r *http.Request) {
				parser.JSON(w, parser.SuccessResponse{
					Status: "success",
					StatusCode: http.StatusOK,
					Data: "Server is up!",
				})
			},
			Middlewares: []Middleware{
				middlewares.Example1,
				middlewares.Example2,
			},
		},
	},
}

