package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
)

func PipeSubscriberBody(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// parse body
		var body services.SubscriberDTO 

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "body of the request is missing",
					Suggestion: "add the body on json format",
					Path: r.RequestURI,
					Timestamp: parser.Timestamp(),
				},
			})
			return
		}

		if !EvalEmail(body.Email) {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusBadRequest,
				Error: parser.Error{
					Code: parser.Errors.BAD_REQUEST_BODY.Code,
					Message: parser.Errors.BAD_REQUEST_BODY.Message,
					Details: "invalid email format",
					Suggestion: "use a valid email",
					Path: r.RequestURI,
					Timestamp: parser.Timestamp(),
				},
			})
			return
		}

		ctx := context.WithValue(r.Context(), "body", body)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
