package drinks

import (
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/config/parser"
	srv "github.com/rootspyro/50BEERS/services/drinks"
)

func ListDrinksForBlog(w http.ResponseWriter, r *http.Request) {

	data, err := srv.GetAllDrinks()

	if err != nil {
		log.Error(err.Error())

		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusInternalServerError,
			Error: parser.Error{
				Code: parser.Errors.INTERNAL_SERVER_ERROR.Code,
				Message: parser.Errors.INTERNAL_SERVER_ERROR.Message,
				Details: "Error getting the data of the drinks",
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
		},
	})	
}
