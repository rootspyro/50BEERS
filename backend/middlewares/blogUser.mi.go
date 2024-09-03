package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)


func PipeNewDrinkBody(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body services.BlogUserDTO
		
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_QUERY.Code,
					Message: parser.Errors.BAD_REQUEST_QUERY.Message,
					Details: "body of the request is missing",
					Suggestion: "add the body on json format",
					Path: r.RequestURI,
					Timestamp: time.Now().Local(),
				},
			})
			return
		}

		next.ServeHTTP(w,r)
	}
}
