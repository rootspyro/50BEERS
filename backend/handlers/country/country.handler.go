package country

import (
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

type CountryHandler struct {
	srv *services.CountrySrv	
}

func NewCountryHandler(service *services.CountrySrv) *CountryHandler{
	return &CountryHandler{
		srv: service,
	}
}

func(h *CountryHandler) ListCountriesForBlog(w http.ResponseWriter, r *http.Request) {

	lang := r.Context().Value("lang").(string)

	countries, err := h.srv.GetAllCountries(lang)

	if err != nil {
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode:  http.StatusInternalServerError,
			Error: parser.Error{
				Code: parser.Errors.INTERNAL_SERVER_ERROR.Code,
				Message: parser.Errors.INTERNAL_SERVER_ERROR.Message,	
				Details: "error getting the countries data",
				Suggestion: parser.Errors.INTERNAL_SERVER_ERROR.Suggestion,
				Path: r.RequestURI,
				Timestamp: time.Now().Local(),
			},
		})
		return
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: CountriesResponse{
			ItemsFound: len(countries),
			Items: countries,
		},
	})
}
