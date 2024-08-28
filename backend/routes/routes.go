package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/handlers/country"
	"github.com/rootspyro/50BEERS/handlers/drinks"
	"github.com/rootspyro/50BEERS/handlers/health"
	"github.com/rootspyro/50BEERS/handlers/location"
	"github.com/rootspyro/50BEERS/handlers/tag"
	"github.com/rootspyro/50BEERS/middlewares"
)

func New(
	healthHandler *health.HealthHandler,
	tagHandler *tag.TagHandler,
	countryHandler *country.CountryHandler,
	locationHandler *location.LocationHandler,
	drinkHandler *drinks.DrinkHandler,
) *http.ServeMux{

	router := http.ServeMux{}

	// API V1

	// Health
	router.HandleFunc("GET /api/v1/health", middlewares.Logger(healthHandler.ServerStatus))

	// Tag
	router.HandleFunc("GET /api/v1/tag/blog", middlewares.Logger(tagHandler.ListCategoriesForBlog))

	//  Country
	router.HandleFunc("GET /api/v1/country/blog", middlewares.Logger(countryHandler.ListCountriesForBlog))

	// Location
	router.HandleFunc("GET /api/v1/location/blog", middlewares.Logger(locationHandler.ListLocationsForBlog))

	// Drinks
	router.HandleFunc("GET /api/v1/drinks/blog", middlewares.Logger(middlewares.ValidateDrinksBlogFilters(drinkHandler.ListDrinksForBlog)))

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

