package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/handlers/drinks"
	"github.com/rootspyro/50BEERS/handlers/health"
	"github.com/rootspyro/50BEERS/middlewares"
)

func New(
	healthHandler *health.HealthHandler,
	drinkHandler *drinks.DrinkHandler,
) *http.ServeMux{

	router := http.ServeMux{}

	// API V1

	// Health
	router.HandleFunc("GET /api/v1/health", middlewares.Logger(healthHandler.ServerStatus))

	// Drinks
	router.HandleFunc("GET /api/v1/drinks/blog", middlewares.Logger(drinkHandler.ListDrinksForBlog))

	// 404 - PATH NOT FOUND
	router.HandleFunc("/", middlewares.Logger(func(w http.ResponseWriter, r *http.Request) {

		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusNotFound,
			Error: parser.Error{
				Code: parser.Errors.PATH_NOT_FOUND.Code,
				Message: parser.Errors.PATH_NOT_FOUND.Message,
				Details: fmt.Sprintf("the resource %s was not found", r.RequestURI),
				Timestamp: time.Now().Local(),	
				Path: r.RequestURI,
			},
		})

	}) )

	return &router
}

