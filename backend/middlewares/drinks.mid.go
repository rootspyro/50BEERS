package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

var allowedSortFields = []string{"name", "date", "stars", "abv", "created_at"}

func fieldIsValid(field string) bool {
	var valid bool = false

	for _, allowedField := range allowedSortFields {
		if field == allowedField {
			valid = true
			break
		}
	}

	return valid
}

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

		// replace "_" with space
		name = strings.ReplaceAll(name, "_", " ")
		country = strings.ReplaceAll(country, "_", " ")
		location = strings.ReplaceAll(location, "_", " ")

		// default values
		var defDirection = "desc"
		var defSort = "created_at"
		var defPage int = 1	
		var defLimit int = 10

		// validate direction is valid
		if direction != "" {

			if direction != "asc" && direction != "desc" {

				parser.JSON(w, parser.ErrorResponse{
					Status: parser.Status.Error,
					StatusCode: http.StatusBadRequest,
					Error: parser.Error{
						Code: parser.Errors.BAD_REQUEST_QUERY.Code,
						Message: parser.Errors.BAD_REQUEST_QUERY.Message,
						Details: "path query direction is invalid",
						Suggestion: "use asc or desc",
						Path: r.RequestURI,
						Timestamp: time.Now().Local(),
					},
				})

				return
			}

			defDirection = direction
		}

		// validate that sort field is valid
		if sortBy != "" {

			if !fieldIsValid(sortBy) {

				var validOptionsStr string = ""

				for index, field := range allowedSortFields {
					validOptionsStr += field
					if index < len(allowedSortFields) - 1 {
						validOptionsStr += ", "
					}
				}

				parser.JSON(w, parser.ErrorResponse{
					Status: parser.Status.Error,
					StatusCode: http.StatusBadRequest,
					Error: parser.Error{
						Code: parser.Errors.BAD_REQUEST_QUERY.Code,
						Message: parser.Errors.BAD_REQUEST_QUERY.Message,
						Details: "path query sortBy is invalid",
						Suggestion: fmt.Sprintf("Use some of this options: %s", validOptionsStr),
						Timestamp: time.Now().Local(),
						Path: r.RequestURI,
					},
				})
				return
			}
			
			defSort = sortBy
		}

		// validate page and limit are integers
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
			SortBy: defSort,
			Direction: defDirection,
			Page: defPage,
			Limit: defLimit,
		}

		ctx := context.WithValue(r.Context(), "filters", filters)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
 
