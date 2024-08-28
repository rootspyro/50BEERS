package country

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
)

type CountryHandler struct {
}

func NewCountryHandler() *CountryHandler{
	return &CountryHandler{}
}

func(h *CountryHandler) ListCountriesForBlog(w http.ResponseWriter, r *http.Request) {

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "Hello world",
	})
}
