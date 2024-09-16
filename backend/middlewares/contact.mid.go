package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

func PipeContactBody(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body services.ContactDTO

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			parser.MISSING_BODY(w, r.RequestURI)
			return
		}

		errResponse := parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusBadRequest,
			Error: parser.Error{
				Code: parser.Errors.BAD_REQUEST_BODY.Code,
				Message: parser.Errors.BAD_REQUEST_BODY.Message,
				Path: r.RequestURI,
			},
		}

		if body.Name == "" {
			errResponse.Error.Details = "name cannot be empty"
			errResponse.Error.Timestamp = parser.Timestamp()

			parser.JSON(w, errResponse)
			return
		}

		if !EvalEmail(body.Email) {
			errResponse.Error.Details = "invalid email format"
			errResponse.Error.Timestamp = parser.Timestamp()

			parser.JSON(w, errResponse)
			return
		}

		if body.Message == "" {
			errResponse.Error.Details = "message is required"
			errResponse.Error.Timestamp = parser.Timestamp()

			parser.JSON(w, errResponse)
			return
		}

		if len(body.Message) < 5 {
			errResponse.Error.Details = "the message must be at least 5 characters long"	
			errResponse.Error.Timestamp = parser.Timestamp()

			parser.JSON(w, errResponse)
			return
		}

		if len(body.Message) > 300 {
			errResponse.Error.Details = "the message cannot exceed 300 characters"	
			errResponse.Error.Timestamp = parser.Timestamp()

			parser.JSON(w, errResponse)
			return
		}

		ctx := context.WithValue(r.Context(), "body", body)

		next.ServeHTTP(w,r.WithContext(ctx))
	}
}
