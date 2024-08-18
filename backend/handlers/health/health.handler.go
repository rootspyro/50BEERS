package health

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
)

func ServerStatus(w http.ResponseWriter, r *http.Request) {
	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "Server is up!",
	})
}
