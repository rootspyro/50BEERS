package drinks

import (
	"net/http"
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

func(h *DrinkHandler) CountDrinks(w http.ResponseWriter, r *http.Request) {

	counts, err := h.srv.CountDrinksForBlog()

	if err != nil {
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusInternalServerError,
			Error: parser.Error{
				Code: parser.Errors.INTERNAL_SERVER_ERROR.Code,
				Message: parser.Errors.INTERNAL_SERVER_ERROR.Message,
				Details: "error counting drinks for blog",
				Suggestion: parser.Errors.INTERNAL_SERVER_ERROR.Suggestion,
				Path: r.RequestURI,
				Timestamp: time.Now().Local(),
			},
		})
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: counts,
	})
}

func(h *DrinkHandler) ListDrinksForBlog(w http.ResponseWriter, r *http.Request) {

	// Get Filters
	lang := r.Context().Value("lang").(string)
	var filters services.DrinkSearchFilters = r.Context().Value("filters").(services.DrinkSearchFilters)

	data, err := h.srv.GetAllDrinks(filters, lang)

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

	// build pagination response 
	pages, err := h.srv.CalculatePages(filters.Limit)

	if err != nil {
		log.Error(err.Error())

		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusInternalServerError,
			Error: parser.Error{
				Code: parser.Errors.INTERNAL_SERVER_ERROR.Code,
				Message: parser.Errors.INTERNAL_SERVER_ERROR.Message,
				Details: "error calculating pagination",
				Timestamp: time.Now().Local(),
				Path: r.RequestURI,
			},
		})

		return
	}

	pagination := Pagination{
		Pages: pages,
		Page: filters.Page,
		PageSize: filters.Limit,
	}

	if data == nil {
		data = []services.DrinkResume{}
	}

	// final response
	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: DrinksResponse{
			ItemsFound: len(data),
			Items: data,
			Pagination: pagination,
			FiltersAllowed: []string{"name", "category", "country", "location", "sortBy", "direction", "page", "limit"},
			FiltersApplied: Filters{
				Name: filters.Name,
				Category: filters.Category,
				Country: filters.Country,
				Location: filters.Location,
				SortBy: filters.SortBy,
				Direction: filters.Direction,
			},
		},
	})	
}
