package subscriber

import (
	"net/http"

	"github.com/rootspyro/50BEERS/config/log"
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
			log.Error(err.Error())
			parser.SERVER_ERROR(w, "error trying to get subscriber", r.RequestURI)
			return
		}
	}

	// insert new subscriber
	data, err := h.srv.NewSubsciber(body.Email)

	if err != nil {
		log.Error(err.Error())
		parser.SERVER_ERROR(w, "error adding new subscriber", r.RequestURI)
	}

	parser.JSON(w, parser.SuccessResponse {
		Status: parser.Status.Success,
		StatusCode: http.StatusCreated,
		Data: data,
	})
}

func(h *SubscriberHandler) RemoveSubscriber(w http.ResponseWriter, r *http.Request) {

	body := r.Context().Value("body").(services.SubscriberDTO)

	// validate if email is subscribed to newsletter
	_, err := h.srv.FindByEmail(body.Email)	
	if err != nil {
		if err == mongo.ErrNoDocuments {
			parser.JSON(w, parser.ErrorResponse{
				Status: parser.Status.Error,
				StatusCode: http.StatusNotFound,
				Error: parser.Error{
					Code: parser.Errors.NOT_FOUND.Code,
					Message: parser.Errors.NOT_FOUND.Message,
					Details: "emails as not subscribed to newsletter",
					Suggestion: "verify the email",
					Path: r.RequestURI,
					Timestamp: parser.Timestamp(),
				},
			})
			return
		} 
		
		parser.SERVER_ERROR(w, "error finding subscriber", r.RequestURI)
		return
	}

	if err := h.srv.RemoveSubscriber(body.Email); err != nil {
		log.Error(err.Error())
		parser.SERVER_ERROR(w, "error trying to remove subscriber", r.RequestURI)
		return
	}

	parser.JSON(w, parser.SuccessResponse{
		Status: parser.Status.Success,
		StatusCode: http.StatusOK,
		Data: "subscription successfully cancelled",
	})
}
