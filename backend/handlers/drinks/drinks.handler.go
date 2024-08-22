package drinks

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

type DrinkHandler struct {
	srv *services.DrinkSrv
}

func NewDrinkHandler(drinkSrv *services.DrinkSrv) *DrinkHandler {
	return &DrinkHandler{
		srv: drinkSrv,
	}
}

func(h *DrinkHandler) ListDrinksForBlog(w http.ResponseWriter, r *http.Request) {

	// Get Filters
	queries := r.URL.Query()
	name := strings.ToLower(queries.Get("name"))
	country := strings.ToLower(queries.Get("country"))
	location := strings.ToLower(queries.Get("location"))
	sortBy := strings.ToLower(queries.Get("sortBy"))
	direction := strings.ToLower(queries.Get("direction"))
	page := queries.Get("page")
	limit := queries.Get("limit")

	parsedPage, _ := strconv.Atoi(page)
	parsedLimit, _ := strconv.Atoi(limit)

	if parsedPage == 0 {
		parsedPage = 1
	}

	if parsedLimit == 0 {
		parsedLimit = 10
	}

	data, err := h.srv.GetAllDrinks(services.DrinkSearchFilters{
		Name: name,
		Country: country,
		Location: location,
		SortBy: sortBy,
		Direction: direction,
		Page: parsedPage,
		Limit: parsedLimit,
	})

	if err != nil {

		log.Error(err.Error())
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusInternalServerError,
			Error: parser.Error{
				Code: parser.Errors.INTERNAL_SERVER_ERROR.Code,
				Message: parser.Errors.INTERNAL_SERVER_ERROR.Message,
				Details: "error getting the drinks data",
				Timestamp: time.Now().Local(),
				Path: r.RequestURI,
			},
		})

		return
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: DrinksResponse{
			ItemsFound: len(data),
			Items: data,
			Pagination: Pagination{
				Page: parsedPage,
				PageSize: parsedLimit,
			},
			FiltersAllowed: []string{"name", "country", "location", "sortBy", "direction"},
			FiltersApplied: Filters{
				Name: name,
				Country: country,
				Location: location,
				SortBy: sortBy,
				Direction: direction,
			},
		},
	})	
}
