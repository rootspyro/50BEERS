package drinks

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
)

type DrinksHandler struct {

}

func New() *DrinksHandler {
	return &DrinksHandler{}
}

func(h *DrinksHandler) ListDrinksForBlog(w http.ResponseWriter, r *http.Request) {

	var tempData []Drink = []Drink{
		{
			ID: "tempID-01",
			Name: "Estrella Damn clasica",
			Type: "Pilsner",
		},
		{
			ID: "tempID-02",
			Name: "Voll Damn",
			Type: "double malt",
		},
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: DrinksResponse{
			ItemsFound: len(tempData),
			Items: tempData,
		},
	})	
}
