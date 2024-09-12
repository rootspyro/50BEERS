package subscriber

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/parser"
	"github.com/rootspyro/50BEERS/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriberHandler struct {
	srv *services.SubscriberSrv
}

func NewSubscriberHandler(srv *services.SubscriberSrv) *SubscriberHandler {
	return &SubscriberHandler{
		srv: srv,
	}
}

func(h *SubscriberHandler) NewSub(w http.ResponseWriter, r *http.Request) {

	body := r.Context().Value("body").(services.SubscriberDTO)

	// validate if subscriber already exists
	_, err := h.srv.FindByEmail(body.Email)
	if err == nil {
		parser.JSON(w, parser.ErrorResponse{
			Status: parser.Status.Error,
			StatusCode: http.StatusConflict,
			Error: parser.Error{
				Code: parser.Errors.CONFLICT.Code,
				Message: parser.Errors.CONFLICT.Message,
				Details: "email is already subscribed to newsletter",
				Path: r.RequestURI,
				Timestamp: parser.Timestamp(),
			},
		})
		return

	} else {
		if err != mongo.ErrNoDocuments {
			parser.SERVER_ERROR(w, "error trying to get subscriber", r.RequestURI)
			return
		}
	}

	parser.JSON(w, parser.SuccessResponse {
		Status: parser.Status.Success,
		StatusCode: http.StatusCreated,
		Data: "Subscription confirmed. Thank you for your support",
	})
}
