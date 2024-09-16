package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	bloguser "github.com/rootspyro/50BEERS/handlers/blogUser"
	"github.com/rootspyro/50BEERS/handlers/contact"
	"github.com/rootspyro/50BEERS/handlers/country"
	"github.com/rootspyro/50BEERS/handlers/drinks"
	"github.com/rootspyro/50BEERS/handlers/health"
	"github.com/rootspyro/50BEERS/handlers/location"
	"github.com/rootspyro/50BEERS/handlers/subscriber"
	"github.com/rootspyro/50BEERS/handlers/tag"
	mid "github.com/rootspyro/50BEERS/middlewares"
)

func New(
	healthHandler *health.HealthHandler,
	tagHandler *tag.TagHandler,
	countryHandler *country.CountryHandler,
	locationHandler *location.LocationHandler,
	drinkHandler *drinks.DrinkHandler,
	blogUser *bloguser.BlogUserHandler,
	subsHandler *subscriber.SubscriberHandler,
	contactHandler *contact.ContactHadler,
) *http.ServeMux{

	router := http.ServeMux{}

	// API V1

	// Health
	router.HandleFunc("GET /api/v1/health", healthHandler.ServerStatus)

	// Tag
	router.HandleFunc("GET /api/v1/tag/blog", mid.LangHeader(tagHandler.ListCategoriesForBlog))

	//  Country
	router.HandleFunc("GET /api/v1/country/blog", mid.LangHeader(countryHandler.ListCountriesForBlog))

	// Location
	router.HandleFunc("GET /api/v1/location/blog", mid.LangHeader(locationHandler.ListLocationsForBlog))

	// Drinks
	router.HandleFunc("GET /api/v1/drinks/blog", mid.LangHeader(mid.ValidateDrinksBlogFilters(drinkHandler.ListDrinksForBlog)))
	router.HandleFunc("GET /api/v1/drinks/blog/count", drinkHandler.CountDrinks)

	// Authentication
	router.HandleFunc("POST /api/v1/auth/blog/signup", mid.PipeNewBlogUserBody(blogUser.SignUp))
	router.HandleFunc("POST /api/v1/auth/blog/login", mid.PipeLoginBody(blogUser.Login))
	router.HandleFunc("GET /api/v1/auth/blog/profile", blogUser.ValidateToken)
	router.HandleFunc("POST /api/v1/auth/blog/logout", blogUser.Logout)

	// Subscribers
	router.HandleFunc("POST /api/v1/newsletter/subscriber", mid.PipeSubscriberBody(subsHandler.NewSub))
	router.HandleFunc("DELETE /api/v1/newsletter/subscriber", mid.PipeSubscriberBody(subsHandler.RemoveSubscriber))

	// Contact
	router.HandleFunc("POST /api/v1/contact/blog", mid.PipeContactBody(contactHandler.EmailFromBlog))

	// 404 - PATH NOT FOUND
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {

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

	}) 


	return &router 
}

