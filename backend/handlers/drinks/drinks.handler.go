package drinks

import (
	"net/http"

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
