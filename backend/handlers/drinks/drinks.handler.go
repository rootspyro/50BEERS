package drinks

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
)

func ListDrinksForBlog(w http.ResponseWriter, r *http.Request) {

	data := []string{}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: DrinksResponse{
			ItemsFound: len(data),
			Items: data,
		},
	})	
}
