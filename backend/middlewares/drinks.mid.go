package middlewares

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

func ValidateDrinksBlogFilters(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get query parameters
		queries := r.URL.Query()
		name := strings.ToLower(queries.Get("name"))
		category := strings.ToLower(queries.Get("category"))
		country := strings.ToLower(queries.Get("country"))
		location := strings.ToLower(queries.Get("location"))
		sortBy := strings.ToLower(queries.Get("sortBy"))
		direction := strings.ToLower(queries.Get("direction"))
		page := queries.Get("page")
		limit := queries.Get("limit")

		// validate page and limit are integers
		var defPage int = 1	
		var defLimit int = 10

		if page != "" {

			intPage, err := strconv.Atoi(page)
			if err != nil {
				parser.JSON(w, parser.ErrorResponse{
					Status: parser.Status.Error,	
					StatusCode: http.StatusBadRequest,
					Error: parser.Error{
						Code: parser.Errors.BAD_REQUEST_QUERY.Code,
						Message: parser.Errors.BAD_REQUEST_QUERY.Message,
						Details: "page path query must be an integer greater than 0",
						Path: r.RequestURI,
						Timestamp: time.Now().Local(),
					},
				})
				return
			}

			if intPage <= 0 {
				parser.JSON(w, parser.ErrorResponse{
					Status: parser.Status.Error,	
					StatusCode: http.StatusBadRequest,
					Error: parser.Error{
						Code: parser.Errors.BAD_REQUEST_QUERY.Code,
						Message: parser.Errors.BAD_REQUEST_QUERY.Message,
						Details: "page path query must be an integer greater than 0",
						Path: r.RequestURI,
						Timestamp: time.Now().Local(),
					},
				})
				return
			}

			// new value to page filter
			defPage = intPage
		}	

		if limit != "" {
			intLimit, err := strconv.Atoi(limit)

			if err != nil {
				parser.JSON(w, parser.ErrorResponse{
					Status: parser.Status.Error,	
					StatusCode: http.StatusBadRequest,
					Error: parser.Error{
						Code: parser.Errors.BAD_REQUEST_QUERY.Code,
						Message: parser.Errors.BAD_REQUEST_QUERY.Message,
						Details: "limit path query must be an integer greater than 0",
						Path: r.RequestURI,
						Timestamp: time.Now().Local(),
					},
				})
				return
			}

			if intLimit <= 0 {
				parser.JSON(w, parser.ErrorResponse{
					Status: parser.Status.Error,	
					StatusCode: http.StatusBadRequest,
					Error: parser.Error{
						Code: parser.Errors.BAD_REQUEST_QUERY.Code,
						Message: parser.Errors.BAD_REQUEST_QUERY.Message,
						Details: "limit path query must be an integer greater than 0",
						Path: r.RequestURI,
						Timestamp: time.Now().Local(),
					},
				})
				return
			}

			defLimit = intLimit
		}

		var filters = services.DrinkSearchFilters{
			Name: name,
			Category: category,
			Country: country,
			Location: location,
			SortBy: sortBy,
			Direction: direction,
			Page: defPage,
			Limit: defLimit,
		}

		ctx := context.WithValue(r.Context(), "filters", filters)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
